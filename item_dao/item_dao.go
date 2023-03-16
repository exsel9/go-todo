package item_dao

import (
	"database/sql"
	"github.com/ichtrojan/go-todo/models"
	log "github.com/sirupsen/logrus"
	"time"
)

type DAO struct {
	db *sql.DB
}

func New(db *sql.DB) *DAO {
	return &DAO{db: db}
}

func (dao *DAO) Add(item string) {
	_, err := dao.db.Exec(`INSERT INTO todos (item) VALUE (?)`, item)

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
	statement, err := dao.db.Query(`SELECT * FROM todos`)

	if err != nil {
		log.Error(err)
	}

	return resultToObject(statement)
}

func (dao *DAO) Focus() []*models.Todo {
	statement, err := dao.db.Query(`SELECT * FROM todos WHERE focused = 1 AND completed = 0`)

	if err != nil {
		log.Error(err)
	}

	return resultToObject(statement)
}

func (dao *DAO) NotCompleted() []*models.Todo {
	statement, err := dao.db.Query(`SELECT * FROM todos WHERE completed = 0`)
	if err != nil {
		log.Error(err)
	}

	return resultToObject(statement)
}

func (dao *DAO) NotDeferred() []*models.Todo {
	statement, err := dao.db.Query(`SELECT * FROM todos WHERE postponed_until_date <= CURRENT_DATE AND completed = 0`)
	if err != nil {
		log.Error(err)
	}

	return resultToObject(statement)
}

func resultToObject(statement *sql.Rows) []*models.Todo {
	var (
		id            int64
		item          string
		completed     int
		focused       int
		repeated      int
		postponedDate time.Time
	)

	var todos []*models.Todo

	for statement.Next() {
		err := statement.Scan(&id, &item, &completed, &focused, &repeated, &postponedDate)

		if err != nil {
			log.Error(err)
		}

		todo := &models.Todo{
			Id:           id,
			Item:         item,
			Completed:    completed == 1,
			Focused:      focused == 1,
			Repeated:     repeated == 1,
			PostponeDate: postponedDate,
		}

		todos = append(todos, todo)
	}

	return todos
}
