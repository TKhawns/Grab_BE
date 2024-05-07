package request

type RequestUpload struct {
	UserId      string `json:"user_id"`
	Date        string `json:"model_date"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Format      string `json:"format"`
}
