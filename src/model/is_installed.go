package model

type isInstalled struct {
	table string
}

var isInstalledModel *isInstalled

func GetIsInstalledModel() *isInstalled {
	if isInstalledModel == nil {
		isInstalledModel = &isInstalled{
			table: "is_installed",
		}
	}
	return isInstalledModel
}

func (d *isInstalled) GetTable() string {
	return d.table
}
