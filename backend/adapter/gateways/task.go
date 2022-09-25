package gateways

import (
	"database/sql"
	"log"
	"myapp/domain/model"
	"myapp/usecases/ports"
	"time"
)

type TaskRepositoryGateway struct {
	databaseHandler *sql.DB
}

type DatabaseHandler struct {
	Conn *sql.DB
}

func NewTaskRepository(dbHandler DatabaseHandler) ports.TaskRepository {
	return &TaskRepositoryGateway{
		databaseHandler: dbHandler.Conn,
	}
}

func (t *TaskRepositoryGateway) AddTask(task *model.Task) (*model.Task, error) {
	newTask, err := Insert(t.databaseHandler, task)
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func (t *TaskRepositoryGateway) GetTasks() ([]*model.Task, error) {
	tasks, err := GetAll(t.databaseHandler)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func Insert(db *sql.DB, task *model.Task) (*model.Task, error) {
	stmt := `insert into tasks (user_id, title, created_at, updated_at)
		values ($1, $2, $3, $4) returning id`

	_, err := db.Exec(stmt,
		task.UserID,
		task.Title,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetALl returns all tasks in db
func GetAll(db *sql.DB) ([]*model.Task, error) {
	query := `select id, user_id, title, created_at, updated_at from tasks`

	var tasks []*model.Task
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.UserID,
			&task.Title,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}
