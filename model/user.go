package model

type User struct {
	UserId   string `json:"-" db:"user_id,omitempty"`
	FullName string `json:"fullName,omitempty" db:"full_name"`
	Phone    string `json:"Phone,omitempty" db:"phone"`
	Password string `json:"Password,omitempty" db:"password"`
	Role     string `json:"Role,omitempty" db:"role"`
	Token    string `json:"Token,omitempty"`
}
