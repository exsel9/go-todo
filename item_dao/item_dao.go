package item_dao

import (
	"database/sql"
	"github.com/ichtrojan/go-todo/models"
	log "github.com/sirupsen/logrus"
)

type DAO struct {
	db *sql.DB
}

func New(db *sql.DB) *DAO {
	return &DAO{db: db}
}

func (dao *DAO) Add(todo models.Todo) {
	_, err := dao.db.Exec(`INSERT INTO todos (item) VALUE (?)`, todo.Item)

	if err != nil {
		log.Error(err)
	}
}

func (dao *DAO) Delete(id string) {
	_, err := dao.db.Exec(`DELETE FROM todos WHERE id = ?`, id)

	if err != nil {
		log.Error(err)
	}
}

func (dao *DAO) MarkAsComplete(id string) {
	_, err := dao.db.Exec(`UPDATE todos SET completed = 1 WHERE id = ?`, id)

	if err != nil {
		log.Error(err)
	}
}

func (dao *DAO) MarkAsUnComplete(id string) {
	_, err := dao.db.Exec(`UPDATE todos SET completed = 0 WHERE id = ?`, id)

	if err != nil {
		log.Error(err)
	}
}

func (dao *DAO) All() []*models.Todo {
	var (
		id        int
		item      string
		completed int
	)

	statement, err := dao.db.Query(`SELECT * FROM todos`)

	if err != nil {
		log.Error(err)
	}

	var todos []*models.Todo

	for statement.Next() {
		err = statement.Scan(&id, &item, &completed)

		if err != nil {
			log.Error(err)
		}

		todo := &models.Todo{
			Id:        id,
			Item:      item,
			Completed: completed,
		}

		todos = append(todos, todo)
	}

	return todos
}
