package dbops

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init()  {
	dbConn, err = sql.Open("mysql",
		"root:12345678o!@#@tcp(localhost:3308)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}