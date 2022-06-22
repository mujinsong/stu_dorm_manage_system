package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/golang-sql/civil"
	_ "github.com/mattn/go-adodb"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var Db *sql.DB
var IsOpen bool = false

type dbuser struct {
	user  string
	pwd   string
	grant int
}

type dbUserIn struct {
	dataSource string // 服务器名称
	port       string //数据端口号
	user       string //用户名
	pwd        string //密码
	database   string //数据库
	provider   string //数据类型
}
var sf int8
func Start() {
	var user dbUserIn
	fmt.Println("1.student")
	fmt.Println("2.manager")
	fmt.Print("选择你的身份:")
	fmt.Scanln(&sf)
	if sf == 2 {
		user.user = "zsm"
	} else if sf == 1 {
		user.user = "stulogin"
	} else {
		fmt.Println("错误输入")
		return
	}
	user.pwd = "zsm20020609"
	user.dataSource = "localhost"
	user.port = "61837"
	user.database = "dbtest"
	user.provider = "SQLOLEDB"
	IsOpen = OpenDataSource(user)
	if !IsOpen {
		return
	}

}

func main() {
	var isout bool=true
	for true {
		Start()
		isout=true
		choice := f_meun()
		if choice!=1&&choice!=2 {
			break
		}
		for isout {
			if choice==0 {
				fmt.Println("out")
				break
			}
			switch choice {
			case 1:
				f := register(Db,sf)
				if f {
					fmt.Println("registered successfully")
					choice=0
				}
			case 2:
				grant := login()
				if grant==0 {
					isout=false
					Db.Close()
					break
				}
				var out bool= true
				for out {
					switch grant {
					case 1:
						c1 := stu_meun()
						if c1==0 {
							fmt.Println("out")
							out=false
							break
						}
						switch c1 {
						case 1:
							fmt.Print("输入学号:")
							var sno string
							fmt.Scanln(&sno)
							query_stu(sno)
						case 2:
							fmt.Print("输入学号:")
							var sno string
							fmt.Scanln(&sno)
							query_prevention(sno)
						case 3:
							var (
								dno     string
								pno     int
								rseason string
								rsubmit string
							)
							fmt.Print("输入宿舍号:")
							fmt.Scanln(&dno)
							fmt.Print("输入物品号:")
							fmt.Scanln(&pno)
							rsubmit = time.Now().Format(time.RFC3339)
							fmt.Print("维修理由(选填):")
							fmt.Scanln(&rseason)
							submit_repair(dno, pno, rsubmit, rseason)
						case 4:
							fmt.Print("输入宿舍号:")
							var dno string
							fmt.Scanln(&dno)
							query_repair(dno)
						case 5:
							var (
								sno     string
								ltime   time.Time
								lreturn string
							)
							fmt.Print("学号:")
							fmt.Scanln(&sno)
							fmt.Print("预计返校时间(选填):")
							fmt.Scanln(&lreturn)
							ltime = time.Now()
							submit_leave(sno, ltime, lreturn)
						case 6:
							var sno string
							fmt.Print("输入学号:")
							fmt.Scanln(&sno)
							query_leave(sno)
							fmt.Print("删除哪个?输入:")
							inputReader := bufio.NewReader(os.Stdin)
							input, err := inputReader.ReadString('\n')
							if err != nil {
								fmt.Println("输入错误")
							}
							strs := strings.Fields(input)
							var nums []int
							for x:=range strs {
								num, err := strconv.Atoi(strs[x])
								if err != nil {
									fmt.Println(err)
									return
								}
								nums = append(nums, num)
							}
							sort.Ints(nums)
							delete_leave(sno,nums)
						case 7:
							var  (
								id string
								pwd string
							)
							fmt.Print("登录名:")
							fmt.Scanln(&id)
							fmt.Print("旧密码:")
							fmt.Scanln(&pwd)
							change_pwd(id,pwd)
						}
					case 2:
						c2 := manage_meun()
						if c2==0 {
							fmt.Println("out")
							out=false
							break
						}
						switch c2 {
						case 1:
							fmt.Print("输入学号:")
							var sno string
							fmt.Scanln(&sno)
							query_stu(sno)
						case 2:
							fmt.Print("输入学号:")
							var sno string
							fmt.Scanln(&sno)
							query_prevention(sno)
						case 3:
							var mess prevent
							fmt.Println("输入(学号 体温 测体温时间 核酸时间 核酸检测结果 已接种第几针):")
							fmt.Scanln(&mess.sno,&mess.vtemp,&mess.vtime1,&mess.vtime2,&mess.vresult,&mess.vin)
							submit_pre(mess)
						case 4:
							var (
								dno     string
								pno     int
								rseason string
								rsubmit string
							)
							fmt.Print("输入宿舍号:")
							fmt.Scanln(&dno)
							fmt.Print("输入物品号:")
							fmt.Scanln(&pno)
							rsubmit = time.Now().Format(time.RFC3339)
							fmt.Print("维修理由(选填):")
							fmt.Scanln(&rseason)
							submit_repair(dno, pno, rsubmit, rseason)
						case 5:
							fmt.Print("输入宿舍号:")
							var dno string
							fmt.Scanln(&dno)
							query_repair(dno)
						case 6:
							var (
								dno     string
								pno     int
								rsubmit string
							)
							fmt.Print("输入宿舍号:")
							fmt.Scanln(&dno)
							fmt.Print("输入物品号:")
							fmt.Scanln(&pno)
							fmt.Print("输入报修时间:")
							fmt.Scanln(&rsubmit)
							change_repair(dno, pno,rsubmit)
						case 7:
							var nowtime=time.Now()
							query_inschool(nowtime)
						case 8:
							var(
								sno string
								btime time.Time
								breason string
							)
							fmt.Println("输入学号 晚归原因")
							fmt.Scanln(&sno,&breason)
							btime=time.Now()
							submit_late(sno,btime,breason)
						case 9:
							query_late()
						case 10:
							var  (
								id string
								pwd string
							)
							fmt.Print("登录名:")
							fmt.Scanln(&id)
							fmt.Print("旧密码:")
							fmt.Scanln(&pwd)
							change_pwd(id,pwd)
						}
					}
				}
				if !out {
					choice=0
					break
				}
			}
		}
	}
}
