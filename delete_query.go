package main

import (
	"fmt"
	"log"
	"time"
)

func delete_leave(sno string, a []int) {
	rows, err := Db.Query("select * from leave where sno=?", sno)
	if err != nil {
		fmt.Println(err)
		return
	}
	var mess leave
	var i, j = 1, 0
	for rows.Next() {
		if err := rows.Scan(&mess.sno, &mess.ltime, &mess.lreturn); err != nil {
			log.Fatal(err)
		}
		if i == a[j] {
			_, err := Db.Exec("delete from leave where sno=? and ltime=?", sno, mess.ltime)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println("ok")
			j += 1
		}
		i += 1
	}
}
func query_inschool(now time.Time) {
	rows, err := Db.Query("select * from student where sno in (select sno from leave where lreturn<?) or sno in ((select sno from student) except (select sno from leave))", now)
	if err != nil {
		fmt.Println(err)
		return
	}
	var mess stu
	for rows.Next() {
		if err := rows.Scan(&mess.sno,&mess.sname,&mess.ssex,&mess.sdept,&mess.dno,&mess.scheckin,&mess.sphone,&mess.ephone); err != nil {
			log.Fatal(err)
		}
		fmt.Println(mess)
	}
}
type repair struct {
	dno string
	pno int
	rsubmit string
	rsolve string
	rseason string
}
type stu struct {
	sno      string
	sname    string
	ssex     string
	sdept    string
	dno      string
	scheckin string
	sphone   string
	ephone   string
}
type leave struct {
	sno string
	ltime string
	lreturn string
}
type prevent struct {
	sno string
	vtemp float64
	vtime1 string
	vtime2 string
	vresult string
	vin string
}
func query_stu(chose string) {
	rows, err := Db.Query("select * from dbo.stu_mess where sno=?", chose)
	if err != nil {
		fmt.Println(err)
		return
	}
	var mess stu
	for rows.Next() {
		if err := rows.Scan(&mess.sno,&mess.sname,&mess.ssex,&mess.sdept,&mess.dno,&mess.scheckin,&mess.sphone,&mess.ephone); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(mess)
}
func query_repair(chose string)  {
	rows, err := Db.Query("select * from repair where dno=?", chose)
	if err != nil {
		fmt.Println(err)
		return
	}
	var mess repair
	for rows.Next() {
		if err := rows.Scan(&mess.dno,&mess.pno,&mess.rsubmit,&mess.rsolve,&mess.rseason); err != nil {
			log.Fatal(err)
		}
		fmt.Println(mess)
	}

}
func query_prevention(chose string) {
	rows, err := Db.Query("select * from prevention where sno=?", chose)
	if err != nil {
		fmt.Println(err)
		return
	}
	var mess prevent
	for rows.Next() {
		if err := rows.Scan(&mess.sno,&mess.vtemp,&mess.vtime1,&mess.vtime2,&mess.vresult,&mess.vin); err != nil {
			log.Fatal(err)
		}
		fmt.Println(mess)
	}

}
func query_leave(chose string)  {
	rows, err := Db.Query("select * from leave where sno=?", chose)
	if err != nil {
		fmt.Println(err)
		return
	}
	var mess leave
	var i = 1
	for rows.Next() {
		if err := rows.Scan(&mess.sno,&mess.ltime,&mess.lreturn); err != nil {
			log.Fatal(err)
		}
		fmt.Println(i,mess)
		i+=1
	}
}
func query_late()  {
	var today=time.Now()
	latetime:=time.Date(today.Year(),today.Month(),today.Day(),22,30,0,0,time.Local)
	rows, err := Db.Query("select * from student where sno in (select sno from late where btime>?)",latetime)
	if err != nil {
		fmt.Println(err)
		return
	}
	var mess stu
	for rows.Next() {
		if err := rows.Scan(&mess.sno,&mess.sname,&mess.ssex,&mess.sdept,&mess.dno,&mess.scheckin,&mess.sphone,&mess.ephone); err != nil {
			log.Fatal(err)
		}
		fmt.Println(mess)
	}

}
