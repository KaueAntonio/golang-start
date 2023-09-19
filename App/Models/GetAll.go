package Models

import db "golang-start/App/App/Database"

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT *
			FROM dbo.TODOS`

	rows, err := conn.Query(sql)

	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
