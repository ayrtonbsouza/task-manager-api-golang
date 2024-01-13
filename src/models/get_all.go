package models

import (
	"task-manager-api/src/db"
)

func GetAll() (tasks []Task, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM Task")
	if err != nil {
		return
	}

	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Deadline)
		if err != nil {
			continue
		}

		tasks = append(tasks, task)
	}
	return
}
