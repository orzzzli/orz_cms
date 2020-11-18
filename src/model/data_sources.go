package model

type dataSources struct {
	table string
}

var dataSourceModel *dataSources

func GetDataSourcesModel() *dataSources {
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
