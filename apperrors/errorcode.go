package apperrors

type ErrCode string

const (
	Unknown ErrCode = "UOOO"

	InsertDataFaild  ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NAData           ErrCode = "S003"
	NoTargetData     ErrCode = "S004"
	UpdataDataFailed ErrCode = "S005"
)
