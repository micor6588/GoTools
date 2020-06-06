package main

import (
	parase "GoTools/ReadFile/fileParase"
	"fmt"
)

func main() {
	c, err := parase.NewFileConf("./demo/config.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.GetString("mysql.db1.Host"))
	fmt.Println(c.GetString("mysql.db1.Name"))
	fmt.Println(c.GetString("mysql.db1.User"))
	fmt.Println(c.GetString("mysql.db1.Pwd"))

	// 取得配置时指定类型
	port, err := c.GetInt("mysql.db1.Port")
	if err != nil {
		panic(err)
	}
	fmt.Println(port) // output: 3306(type int)
}
