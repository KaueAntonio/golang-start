package Models

import db "golang-start/App/App/Database"

func Create(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO dbo.todos (title, description, done)
			VALUES ($1, $2, $3)
			RETURNING ID`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
