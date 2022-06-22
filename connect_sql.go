package main
import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-adodb"
	"log"
)
func OpenDataSource(user dbUserIn) bool {
	var IsOpen bool
	conStr := fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;password=%s",
		user.dataSource, user.port, user.database, user.user, user.pwd)
	var err error
	Db, err = sql.Open("mssql", conStr)
	if err != nil {
		log.Fatalln("open函数发生错误！")
		IsOpen = false
		return IsOpen
	}
	err = (Db).Ping()
	if err != nil {
		fmt.Println("连接数据库失败！" + conStr + err.Error())
		IsOpen = false
		return IsOpen
	}
	fmt.Println("连接上数据库")
	fmt.Printf("%s opened successfully\n", user.database)
	IsOpen = true
	return IsOpen
}
