package model

const (
	SourceTypeLink = 1
	SourceTypeMysql
	SourceTypeRedis
	SourceTypeFile
	SourceTypeStep

	StepViewTypeList = 1
	StepViewTypeForm
	StepViewTypeGraph

	StepActionTypeInsert = 1
	StepActionTypeUpdate
	StepActionTypeDelete
	StepActionTypeDownload
)

func CalculateAction(canInsert bool, canUpdate bool, canDelete bool, canDownload bool) int {
	temp := 0
	if canInsert {
		temp = temp | (1 << 0)
	}
	if canUpdate {
		temp = temp | (1 << 1)
	}
	if canDelete {
		temp = temp | (1 << 2)
	}
	if canDownload {
		temp = temp | (1 << 3)
	}
	return temp
}
