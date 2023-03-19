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

func (dao *DAO) Add(item string) int64 {
	result, err := dao.db.Exec(`INSERT INTO todos (item) VALUE (?)`, item)
	if err != nil {
		log.Error(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Error(err)
	}

	return id
}

func (dao *DAO) Delete(id string) {
	_, err := dao.db.Exec(`DELETE FROM todos WHERE id = ?`, id)

	if err != nil {
		log.Error(err)
	}
}

func (dao *DAO) MarkAsComplete(id string) {
	_, err := dao.db.Exec(`UPDATE todos SET completed_date = CURRENT_DATE WHERE id = ?`, id)

	if err != nil {
		log.Error(err)
	}
}

func (dao *DAO) MarkAsUnComplete(id string) {
	_, err := dao.db.Exec(`UPDATE todos SET completed_date = NULL WHERE id = ?`, id)

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
	statement, err := dao.db.Query(`SELECT * FROM todos WHERE focused = 1 AND completed_date IS NULL`)

	if err != nil {
		log.Error(err)
	}

	return resultToObject(statement)
}

func (dao *DAO) NotCompleted() []*models.Todo {
	statement, err := dao.db.Query(`SELECT * FROM todos WHERE completed_date IS NULL`)
	if err != nil {
		log.Error(err)
	}

	return resultToObject(statement)
}

func (dao *DAO) Completed() []*models.Todo {
	statement, err := dao.db.Query(`SELECT * FROM todos WHERE completed_date IS NOT NULL`)
	if err != nil {
		log.Error(err)
	}

	return resultToObject(statement)
}

func (dao *DAO) CompletedToday() []*models.Todo {
	statement, err := dao.db.Query(`SELECT * FROM todos WHERE completed_date = CURRENT_DATE`)
	if err != nil {
		log.Error(err)
	}

	return resultToObject(statement)
}

func (dao *DAO) Today() []*models.Todo {
	statement, err := dao.db.Query(`SELECT * FROM todos WHERE completed_date = CURRENT_DATE 
                       OR (completed_date IS NULL AND postponed_until_date <= CURRENT_DATE)`)
	if err != nil {
		log.Error(err)
	}

	return resultToObject(statement)
}

func (dao *DAO) NotPostponed() []*models.Todo {
	statement, err := dao.db.Query(`SELECT * FROM todos WHERE postponed_until_date <= CURRENT_DATE AND completed_date IS NULL`)
	if err != nil {
		log.Error(err)
	}

	return resultToObject(statement)
}

func (dao *DAO) Postponed() []*models.Todo {
	statement, err := dao.db.Query(`SELECT * FROM todos WHERE postponed_until_date > CURRENT_DATE AND completed_date IS NULL`)
	if err != nil {
		log.Error(err)
	}

	return resultToObject(statement)
}

func resultToObject(statement *sql.Rows) []*models.Todo {
	var (
		id            int64
		item          string
		focused       int
		repeated      int
		postponedDate time.Time
		competedDate  *time.Time
	)

	var todos []*models.Todo

	for statement.Next() {
		err := statement.Scan(&id, &item, &focused, &repeated, &postponedDate, &competedDate)

		if err != nil {
			log.Error(err)
		}

		todo := &models.Todo{
			Id:            id,
			Item:          item,
			Focused:       focused == 1,
			Repeated:      repeated == 1,
			PostponeDate:  postponedDate,
			CompletedDate: competedDate,
		}

		todos = append(todos, todo)
	}

	return todos
}
