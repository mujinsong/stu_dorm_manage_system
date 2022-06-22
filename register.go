package main

import (
	"database/sql"
	"fmt"
)

func register(Db *sql.DB,grant int8) bool {
	var (
		id  string
		pwd string
	)
	for true {
		fmt.Print("用户名:")
		fmt.Scanln(&id)
		fmt.Print("密码:")
		fmt.Scanln(&pwd)
		rows, err := Db.Query("select id from dbuser where id=?",id)
		if err != nil {
			return false
		}
		var t string
		for rows.Next() {
			rows.Scan(&t)
		}
		if t=="" {
			fmt.Println("用户名未被使用,可以注册")
			_, err = Db.Exec("insert into dbuser(id,pwd,the_grant)values (?,?,?)", id, pwd, grant)
			if err != nil {
				fmt.Println("注册失败!", err)
				return false
			}
			return true
		}
		fmt.Println("该用户名已被注册")
	}

	return false
}
