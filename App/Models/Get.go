package Models

import db "golang-start/App/App/Database"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	sql := `SELECT *
			FROM dbo.TODOS
			WHERE TODOS.ID = $1`

	rows := conn.QueryRow(sql, id)

	err = rows.Scan(&todo.ID, &todo.Title, &todo.Done)

	return
}
