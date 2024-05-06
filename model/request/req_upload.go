package request

type RequestUpload struct {
	Date        string `json:"model_date"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Content     []byte `json:"content"`
}
