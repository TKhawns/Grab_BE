package model

type Model struct {
	Model_id    string `json:"Model_id" db:"model_id,omitempty"`
	User_id     string `json:"UserId,omitempty" db:"user_id"`
	Name        string `json:"Name,omitempty" db:"name"`
	Date        string `json:"Date,omitempty" db:"model_date"`
	Description string `json:"Description,omitempty" db:"description"`
	Status      string `json:"Status,omitempty" db:"status"`
	Content     string `json:"Content,omitempty" db:"content"`
	Format      string `json:"Format,omitempty" db:"format"`
}
