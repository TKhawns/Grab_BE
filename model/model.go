package model

type Model struct {
	Model_id    string `json:"-" db:"model_id,omitempty"`
	Name        string `json:"Name,omitempty" db:"name"`
	Date        string `json:"Date,omitempty" db:"model_date"`
	Description string `json:"Description,omitempty" db:"description"`
	Status      string `json:"Status,omitempty" db:"status"`
	Content     []byte `json:"Content,omitempty" db:"content"`
}
