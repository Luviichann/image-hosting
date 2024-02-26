package models

type Folder struct {
	Id         int
	FolderName string
}

func (Folder) TableName() string {
	return "folder"
}
