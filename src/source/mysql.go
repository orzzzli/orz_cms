package source

import (
	"errors"
	"time"

	"github.com/orzzzli/orz_cms/src/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Mysql struct {
	title     string //名称
	scheme    string //连接信息
	connected bool   //是否已经连接过
	db        *sqlx.DB
	pingGap   int //保活间隔，单位s
}

func NewMysql(title string, gap int) *Mysql {
	return &Mysql{
		title:     title,
		scheme:    "",
		connected: false,
		pingGap:   gap,
	}
}

/*
	Sqlx惰性连接，open并不会实际建立连接。
	因此此方法err不能用来验证连接是否成功。
*/
func (m *Mysql) Connect(scheme string) (err error) {
	if m.connected == true {
		return nil
	}
	m.db, err = sqlx.Open("mysql", scheme)
	m.scheme = scheme
	m.connected = true

	go m.keepLive()

	return err
}

/*
	重连，更新结构体下的db，使用一个新的sqlx结构体替换原db，并自动ping
*/
func (m *Mysql) ReConnect() (err error) {
	m.db, err = sqlx.Open("mysql", m.scheme)
	if err != nil {
		return err
	}
	m.connected = true
	err = m.Ping()
	return err
}

/*
	使用mysql_ping进行连接状态的探测。
*/
func (m *Mysql) Ping() (err error) {
	if !m.connected {
		return errors.New("not connected, try Connect first")
	}
	err = m.db.Ping()
	if err != nil {
		m.connected = false
	}
	return err
}

/*
	暂时只支持传统意义的select，结果为[]map[string]string
	todo：验证其他类型语句的正确性。
*/
func (m *Mysql) Get(query string) (interface{}, error) {
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}
	//格式化成[]map[string]string
	columns, err := rows.Columns() //列名list
	if err != nil {
		return nil, err
	}
	columnNumbers := len(columns)
	output := make([]map[string]string, 0) //最终输出map list
	values := make([]interface{}, columnNumbers)
	temp := make([]interface{}, columnNumbers)
	for rows.Next() {
		//scan的值存入values list中
		for i := 0; i < columnNumbers; i++ {
			temp[i] = &values[i]
		}
		err = rows.Scan(temp...)
		if err != nil {
			return nil, err
		}
		oneRow := make(map[string]string)
		for i, key := range columns {
			val := values[i]
			temp, ok := val.([]byte)
			if !ok {
				return nil, errors.New("select value convert error. sql is " + query)
			}
			oneRow[key] = string(temp)
		}
		output = append(output, oneRow)
	}
	return output, nil
}

/*
	返回[]int，第一个值为lastId，第二值为effectedRow
*/
func (m *Mysql) Set(query string) (interface{}, error) {
	res, err := m.db.Exec(query)
	if err != nil {
		return nil, err
	}
	lastInsertId, _ := res.LastInsertId()
	effected, _ := res.RowsAffected()
	temp := make([]int, 0)
	temp = append(temp, int(lastInsertId))
	temp = append(temp, int(effected))
	return temp, nil
}

func (m *Mysql) keepLive() {
	ticker := time.NewTicker(time.Duration(m.pingGap) * time.Second)
	defer ticker.Stop()

	logger.Info(m.title + " heartbeat start.")

	for range ticker.C {
		logger.Info(m.title + " ticker triggered.")
	RETRY:
		err := m.Ping()
		if err != nil {
			logger.Info(m.title + " ticker reconnected.")
			err = m.ReConnect()
			if err != nil {
				logger.Error(m.title + " reconnect fail.")
				return
			}
			goto RETRY
		}
	}
}
