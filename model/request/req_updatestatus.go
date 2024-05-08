package request

type RequestUpdateStatus struct {
	ModelId string `json:"model_id" db:"model_id"`
	Status  string `json:"status" db:"status"`
}
