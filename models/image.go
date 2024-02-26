package models

type Image struct {
	Id         int
	Filename   string
	FileBelong string
}

func (Image) TableName() string {
	return "image"
}
