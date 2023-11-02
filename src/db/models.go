package db

import "gorm.io/gorm"

type Lesson struct {
	gorm.Model
	Id    uint64 `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"unique"`
	Files []File `gorm:"foreignkey:LessonID"` 
}

type File struct {
	gorm.Model
	Id   uint64 `gorm:"primaryKey;autoIncrement"`
	LessonID uint
	Name string
	Type string
}
