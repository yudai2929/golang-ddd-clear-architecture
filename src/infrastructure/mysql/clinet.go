package mysql

import (
	"database/sql"
	"golang-ddd-clear-architecture/day4/task3/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	env := config.NewEnv()
	db, err := sql.Open(env.Driver, env.Username+":"+env.Password+"@"+env.Protocol+"("+env.Host+":"+env.Port+")/"+env.Name+"?parseTime="+env.Parse)
	if err != nil {
		panic(err)
	}
	return db
}
