package prepared

import (
	"database/sql"
	"errors"
)

type query string
type name string

var NotFoundError = errors.New("query not found")
var NilStatementError = errors.New("query statement is nil")

var queries = map[name]query{
	"SetVerdict":      `UPDATE Users SET msg = ? WHERE student = ?;`,
	"UpdateUserScore": `UPDATE Users SET task_id = task_id + 1, msg = '' WHERE student = ?`,
	"UpdateTask":      `INSERT OR REPLACE INTO Tasks (task_id, title) VALUES (?, ?)`,
	"AddPullRequest":  `INSERT INTO PullRequests (link, student) VALUES (?, ?)`,
	"ChangePassword":  `UPDATE Users SET password = ? WHERE student = ?`,
	"CheckPassword":   `SELECT count(*) FROM Users WHERE student = ? AND password = ?`,
	//"CheckIsMentor":     `SELECT count(*) FROM Mentors WHERE username = ?`,
	"CreateUser":        `INSERT INTO Users (student, password, task_id, pending) VALUES (?, ?, 1, 0)`,
	"DeletePullRequest": `DELETE FROM PullRequests WHERE student = ?`,
	"DeleteTask":        `DELETE FROM Tasks WHERE task_id = ?`,
	"GetTaskIDAndMsg":   `SELECT task_id, msg, (SELECT count(*) FROM Tasks) as tasks_count FROM Users WHERE student = ?`,
	"GetWorks":          `SELECT student, link FROM PullRequests`,
	"GetTasks":          `SELECT task_id, title FROM Tasks`,
	"GetTitle":          `SELECT title FROM Tasks WHERE task_id = ?`,
	"GetUsers":          `SELECT student, task_id, msg FROM Users`,
	"SetPending":        `UPDATE Users SET pending = ? WHERE student = ?`,
	"CheckIsPending":    `SELECT pending FROM Users WHERE student = ?`,
}

var statements = make(map[name]*sql.Stmt, 10)

func Prepare(DB *sql.DB) error {
	for n, q := range queries {
		prep, err := DB.Prepare(string(q))
		if err != nil {
			return err
		}
		statements[n] = prep
	}
	return nil
}

func GetPreparedStatement(Name string) (*sql.Stmt, error) {
	stmt, ok := statements[name(Name)]
	if !ok {
		return nil, NotFoundError
	}

	if stmt == nil {
		return nil, NilStatementError
	}

	return stmt, nil
}
