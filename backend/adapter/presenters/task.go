package presenters

import (
	"encoding/json"
	"myapp/domain/model"
	"myapp/usecases/ports"
	"net/http"
)

type TaskPresenter struct {
	w http.ResponseWriter
}

func NewUserOutputPort(w http.ResponseWriter) ports.TaskOutputPort {
	return &TaskPresenter{
		w: w,
	}
}

func (t *TaskPresenter) OutputTasks(tasks []*model.Task) {
	writeJson(t.w, 200, tasks)
}

func (t *TaskPresenter) OutputError(err error) {
	writeJson(t.w, 400, err)
}

func writeJson(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	w.Write(out)

	return nil
}
