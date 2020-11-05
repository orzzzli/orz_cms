package source

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Mysql struct {
	scheme    string //连接信息
	connected bool   //是否已经连接过
	db        *sqlx.DB
}

func NewMysql() *Mysql {
	return &Mysql{
		scheme:    "",
		connected: false,
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
