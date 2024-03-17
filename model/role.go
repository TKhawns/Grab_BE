package model

type Role int

const (
	MEMBER Role = iota // auto increase STT //0
	ADMIN              //1
)

func (r Role) String() string {
	return []string{"MEMBER", "ADMIN"}[r]
}
