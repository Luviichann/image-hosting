package models

type License struct {
	Id        int
	SecretKey string
}

func (License) TableName() string {
	return "license"
}
