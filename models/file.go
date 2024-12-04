package models

type File struct {
	commonBaseModel `gorm:"embedded"`
	FileName string `json:"file_name"`
	FileSize int64 `json:"file_size"`
	FileType string `json:"file_type" gorm:"default:'Unknown'"`
	FilePath string `json:"file_path"`
	User User `gorm:"ForeignKey:UserId"`
	UserId int `json:"user_id"`
}

func (table *File) TableName() string {
	return "file"
}