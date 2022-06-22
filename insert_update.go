package main

import (
	"fmt"
	"log"
	"time"
)

func submit_repair(dno string, pno int, rsubmit string, rseason string) {
	_, err := Db.Exec("insert into repair (dno,pno,rsubmit,rsolve,rseason)values (?,?,?,'',?)",dno,pno,rsubmit,rseason)
	if err != nil {
		fmt.Println("提交错误")
		log.Fatal(err)
		return
	}
	fmt.Println("successfully")
}
func submit_leave(sno string,ltime time.Time,lreturn string)  {
	_, err := Db.Exec("insert into leave (sno,ltime,lreturn)values (?,?,?)",sno,ltime,lreturn)
	if err != nil {
		fmt.Println("提交错误")
		log.Fatal(err)
		return
	}
	fmt.Println("successfully")
}
func submit_pre(mess prevent) {
	_, err := Db.Exec("insert into prevention (sno,vtemp,vtime1,vtime2,vresult,vin)values (?,?,?,?,?,?)",mess.sno,mess.vtemp,mess.vtime1,mess.vtime2,mess.vresult,mess.vin)
	if err != nil {
		fmt.Println("提交错误")
		log.Fatal(err)
		return
	}
	fmt.Println("successfully")
}
func change_pwd(id string,pwd string)  {
	rows,err:= Db.Query("select pwd from dbuser where ID=?",id)
	if err != nil {
		fmt.Println(err)
	}
	var temp string
	for rows.Next() {
		err:=rows.Scan(&temp)
		if err != nil {
			fmt.Println(err)
		}
	}
	if temp!=pwd {
		fmt.Println("旧密码错误")
		return
	}
	fmt.Println("输入新密码")
	fmt.Scanln(&temp)
	_, err = Db.Exec("update dbuser set pwd=? where ID=? and pwd=?", temp, id, pwd)
	if err != nil {
		fmt.Println(err)
	}
}
func change_repair(dno string,pno int,rsubmit string)  {
	fmt.Print("解决时间:")
	var rsolve string
	fmt.Scanln(&rsolve)
	_, err := Db.Exec("update repair set rsolve=? where dno=? and pno=? and rsubmit=?", rsolve, dno, pno,rsubmit)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("successfully")
}
func submit_late(sno string,btime time.Time,breason string)  {
	_, err := Db.Exec("insert into late (sno,btime,breason)values (?,?,?)",sno,btime,breason)
	if err != nil {
		fmt.Println("提交错误")
		log.Fatal(err)
		return
	}
	fmt.Println("successfully")
}
