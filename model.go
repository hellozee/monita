package main

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type Task struct {
	ID         int    `json:"id"`
	Desciption string `json:"description"`
}

type TaskModel struct {
	db           *gorm.DB
	currentIndex int
}

func (tm *TaskModel) AddTask(description string) {
	ts := Task{
		Desciption: description,
	}
	tm.db.Create(&ts)
}

func (tm *TaskModel) DeleteTask(id int) {
	tm.db.Delete(&Task{}, id)
}

func (tm *TaskModel) ToJSON() []byte {
	tasks := []Task{}
	tm.db.Find(&tasks)
	bytes, _ := json.Marshal(tasks)
	return bytes
}

func NewTaskModel() TaskModel {
	viper.SetConfigName("dbconfig")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	viper.SetDefault("user", "root")
	viper.SetDefault("pass", "root")
	viper.SetDefault("address", "localhost:3306")

	user := viper.GetString("user")
	pass := viper.GetString("pass")
	address := viper.GetString("address")

	db, err := gorm.Open("mysql", user+":"+pass+"@tcp("+address+")/monitadb?charset=utf8&parseTime=True")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Task{})

	tasks := []Task{}
	db.Find(&tasks)
	dbEntries.Set(float64(len(tasks)))

	return TaskModel{
		currentIndex: 0,
		db:           db,
	}
}
