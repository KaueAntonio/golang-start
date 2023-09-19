package Models

import (
	db "golang-start/App/App/Database"
)

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := `UPDATE 
			dbo.TODOS
			SET TODOS.title       = $1,
				TODOS.description = $2,
				TODOS.done		= $3
			WHERE TODOS.id = $4`

	res, err := conn.Exec(sql, todo.Title, todo.Description, todo.Done, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
