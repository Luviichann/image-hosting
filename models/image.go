package models

type Image struct {
	Id         int
	Filename   string
	FileBelong string
	UserId     int
}

func (Image) TableName() string {
	return "image"
}
