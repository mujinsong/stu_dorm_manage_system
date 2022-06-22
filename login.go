package main

import (
	"fmt"
	_ "fmt"
)

func login() int{
	var user dbuser
	fmt.Print("用户名:")
	fmt.Scan(&user.user)
	fmt.Print("密码:")
	fmt.Scan(&user.pwd)

	row,err:= Db.Query("select the_grant from dbuser where ID=? and pwd=?",user.user,user.pwd)
	if err!=nil {
		fmt.Println("登录失败")
		return 0
	}
	for row.Next() {
		row.Scan(&user.grant)
	}
	if user.grant==0 {
		fmt.Println("用户名或密码有误")
		return 0
	}
	fmt.Println("登录成功!")
	return user.grant
}