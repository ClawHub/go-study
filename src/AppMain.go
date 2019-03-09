package main

import (
	"fmt"
	"go-study/src/document"
	"os"
	"test-log/src/log"
)

func main() {
	fmt.Println("init main")

	fmt.Println("------log---------")
	log.MainLogger.Debug("hello main Debug")
	log.MainLogger.Info("hello main Info")
	log.GatewayLogger.Debug("Hi Gateway Im Debug")
	log.GatewayLogger.Info("Hi Gateway  Im Info")

	fmt.Println("------properties---------")
	//打开文件
	file, err := os.Open("./config.properties")
	if err != nil {
		panic(err)
	}

	/* defer代码块会在函数调用链表中增加一个函数调用。
	 * 这个函数调用不是普通的函数调用，而是会在函数正常返回，也就是return之后添加一个函数调用。
	 * 因此，defer通常用来释放函数内部变量。
	 */
	defer file.Close()

	//加载配置文件
	config, err := document.Load(file)
	if nil != err {
		fmt.Println(err)
		return
	}
	//读取配置
	val := config.String("key")
	fmt.Println("val:" + val)
}
