# GoTools
## 截入配置文件并读取其中内容
 ## @Authr micor
 ## @Date  2019/12/08
 
## 与INI配置文风格一样 根据顺序读取文件和每一行 如果在行首出现了;号，则认为是配置文件的注释
## 当INI不规范时 如[mysql] 的注释被写为[mysql 则会返回错误
## 本包的配置文件是严格区分大小写的 需要禁止区分大小写 将在后期加入或自行加入

## 这是一个用于引入配置文件的包，加载类似于ini风格的配置文件。 与INI配置文件风格一样 根据顺序读取文件和每一行 如果在行首出现了;号，则认为是配置文件的注释。 除此之外，还提供多种可选类型。

## 主要支持以下几种类型转换：String、Int、Int64、Float64

## INI演示文件：./demo/config.ini
## 测试：
# [mysql]
## db1.Name = testMysqlDB
## db1.Host = 127.0.0.1
## db1.Port = 3306
## db1.User = root
## db1.Pwd = test

## ; 测试INI风格的注释
## ; 这两行数据的前前面加入了分号，因此，这些数据被认为是注释，将不会影响配置文件正常情况

## [tcp]
## Port=3309

### 测试代码：
### func main() {
###	c, err := parase.NewFileConf("./demo/config.ini")
###	if err != nil {
###		fmt.Println(err)
###		return
###	}
###	fmt.Println(c.GetString("mysql.db1.Host"))
###	fmt.Println(c.GetString("mysql.db1.Name"))
###	fmt.Println(c.GetString("mysql.db1.User"))
###	fmt.Println(c.GetString("mysql.db1.Pwd"))

###	// 取得配置时指定类型
###	port, err := c.GetInt("mysql.db1.Port")
###	if err != nil {
###		panic(err)
###	}
###	fmt.Println(port) // output: 3306(type int)
### }

