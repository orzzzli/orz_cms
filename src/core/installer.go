package core

import (
	"strconv"

	"github.com/orzzzli/orz_cms/src/logger"
	"github.com/orzzzli/orz_cms/src/source"
)

var tableComplete = 0

//检查数据表完整性
func CheckTableComplete(source *source.Mysql) bool {
	res, err := source.Get("show tables")
	if err != nil {
		logger.Error("check table error:" + err.Error())
		return false
	}
	res1, ok := res.([]map[string]string)
	if !ok {
		logger.Error("check table res format error")
		return false
	}
	for _, value := range res1 {
		for _, value1 := range value {
			if value1 == "data_sources" {
				tableComplete = tableComplete | (1 << 0)
			}
			if value1 == "is_installed" {
				tableComplete = tableComplete | (1 << 1)
			}
			if value1 == "steps" {
				tableComplete = tableComplete | (1 << 2)
			}
		}
	}
	if tableComplete != 7 {
		logger.Error("table is not complete, complete is " + strconv.Itoa(tableComplete))
		return false
	}
	return true
}

//安装默认表
func InstallTable(source *source.Mysql) bool {
	for key, value := range AllDefaultTables {
		_, err := source.Set(value)
		if err != nil {
			logger.Error("install table " + key + " error:" + err.Error())
			return false
		}
	}
	return true
}

//确认安装，往已安装表中插入数据
func ConfirmInstall(source *source.Mysql) bool {
	_, err := source.Set(InstalledSql)
	if err != nil {
		logger.Error("confirm install error:" + err.Error() + " sql is:" + InstalledSql)
		return false
	}
	return true
}

//是否安装
func IsInstalled(source *source.Mysql) bool {
	res, err := source.Get(IsInstalledSql)
	if err != nil {
		logger.Error("check is installed error:" + err.Error() + " sql is:" + IsInstalledSql)
		return false
	}
	res1, ok := res.([]map[string]string)
	if !ok {
		logger.Error("check is installed res format error")
		return false
	}
	//已安装
	if len(res1) > 0 {
		return true
	}
	return false
}
