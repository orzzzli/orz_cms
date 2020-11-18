package model

import (
	"fmt"
)

type dataSources struct {
	table string
}

var dataSourceModel *dataSources

func GetDataSourcesModel() *dataSources {
	fmt.Println(dataSourceModel)
	if dataSourceModel == nil {
		dataSourceModel = &dataSources{
			table: "data_sources",
		}
	}
	return dataSourceModel
}

func (d *dataSources) GetTable() string {
	return d.table
}
