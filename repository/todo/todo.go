package todo

import (
	"gorm.io/gorm"
	"todo/entity"
)

type (
	Todo struct {
		db *gorm.DB
	}
	TodoRepository interface {
		CreateTodo(todo *entity.ToDo) (*entity.ToDo, error)
		GetTodoByID(id int64) (*entity.ToDo, error)
		GetAllTodo(groupId int64) ([]*entity.ToDo, error)
		UpdateTodo(todo *entity.ToDo) (*entity.ToDo, error)
		DeleteTodoByID(id int64) error
	}
)

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &Todo{
		db: db,
	}
}

func (t *Todo) CreateTodo(todo *entity.ToDo) (*entity.ToDo, error) {
	tx := t.db.Create(todo)

	if tx.Error != nil {
		return todo, tx.Error
	}

	return todo, nil
}

func (t *Todo) GetTodoByID(id int64) (*entity.ToDo, error) {
	var todo = &entity.ToDo{}
	tx := t.db.First(&todo, id)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return todo, nil
}

func (t *Todo) GetAllTodo(groupId int64) ([]*entity.ToDo, error) {
	var (
		todos []*entity.ToDo
		tx    *gorm.DB
	)

	if groupId > 0 {
		tx = t.db.Where("activity_group_id = ?", groupId).Find(&todos)
	} else {
		tx = t.db.Find(&todos)
	}

	if tx.Error != nil {
		return nil, tx.Error
	}

	return todos, nil
}

func (t *Todo) UpdateTodo(todo *entity.ToDo) (*entity.ToDo, error) {

	var result = &entity.ToDo{}
	tx := t.db.First(result, todo.ID).Updates(todo)

	if tx.Error != nil {
		return result, tx.Error
	}

	return result, nil
}

func (t *Todo) DeleteTodoByID(id int64) error {
	var todo = &entity.ToDo{}
	tx := t.db.First(todo, id).Delete(todo, id)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
