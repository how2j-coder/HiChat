package models

type File struct {
	CommonBaseModel `gorm:"embedded"`
	FileName string `json:"file_name"`
	FileSize int64 `json:"file_size"`
	FileType string `json:"file_type" gorm:"default:'Unknown'"`
	FilePath string `json:"file_path"`
	User User `gorm:"foreignkey:UserID" json:"-"`
	UserID string `gorm:"type:varchar(36)" json:"user_id"`
}

func (table *File) TableName() string {
	return "file"
}