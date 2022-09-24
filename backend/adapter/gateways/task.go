package gateways

import (
	"context"
	"database/sql"
	"log"
	"myapp/domain/model"
	"myapp/usecases/ports"
	"time"
)

type TaskGateway struct {
	clientFactory *sql.DB
}

func NewTaskRepository() ports.TaskRepository {
	db, _ := sql.Open("postgres", "user=postgres password=password host=localhost port=5432 dbname=practice sslmode=disable")
	return &TaskGateway{
		clientFactory: db,
	}
}

func (t *TaskGateway) AddTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	newTask, err := Insert(t.clientFactory, task)
	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func (t *TaskGateway) GetTasks(ctx context.Context) ([]*model.Task, error) {
	tasks, err := GetAll(t.clientFactory)

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
