package models

import (
	"bubble/dao"
)

// todo model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*
Todo这个模型的增删改查操作都在这里
*/
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todolist []*Todo, err error) {
	if err = dao.DB.Find(&todolist).Error; err != nil {
		return nil, err
	}
	return
}

func GetAtodo(id string) (todo *Todo, err error) {
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {

		return nil, err

	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeletATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
