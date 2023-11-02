package db

import (
	"fmt"

	"gorm.io/gorm"
)


var db  = InitDB()


func CreateLesson(name string) {
	lesson := Lesson{Name: name}


	if err := db.Where("Name = ?", name).First(&lesson).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			db.Create(&lesson)
			fmt.Println("Lesson created successfully")
		} else {
			fmt.Printf("Error while checking lesson existence: %v\n", err)
		}
	} else {
		fmt.Println("Lesson already exists")
	}
}

func ListOfLessons() []Lesson {

	var lessons []Lesson

	db.Find(&lessons)

	return lessons
}


func CreateFileForLesson(name string){
	
}