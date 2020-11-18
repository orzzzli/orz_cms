package model

type steps struct {
	table string
}

var stepsModel *steps

func GetStepsModel() *steps {
	if stepsModel == nil {
		stepsModel = &steps{
			table: "steps",
		}
	}
	return stepsModel
}

func (d *steps) GetTable() string {
	return d.table
}
