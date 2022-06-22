package main

import "fmt"

func f_meun() int8{
	fmt.Println("1.regiter")
	fmt.Println("2.login")
	fmt.Println("0.out")
	var choice int8
	fmt.Print("You want:")
	fmt.Scanln(&choice)
	if choice==1||choice==2 {
		return choice
	} else {
		return 0
	}
}

func stu_meun() int {
	fmt.Print("1.查询学生信息\n2.查询核酸信息\n3.提交报修信息\n4.查询报修信息\n5.插入离返校信息\n6.删除离返校信息\n7.改密码\n0.out\n")
	fmt.Print("你的操作:")
	var c1 int
	fmt.Scanln(&c1)
	return c1
}

func manage_meun() int {
	fmt.Print("1.查询学生信息\n2.查询核酸信息\n3.上传核酸信息\n4.提交报修信息\n5.查询报修信息\n6.修改报修信息\n7.查询在校人员信息\n8.插入夜归信息\n9.查询夜归人员\n10.改密码\n0.out\n")
	fmt.Print("你的操作:")
	var c1 int
	fmt.Scanln(&c1)
	return c1
}
