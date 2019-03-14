package document

import (
	"fmt"
	"os"
)

//初始化
func DemoPropertiesV1() {
	fmt.Println("------properties v1---------")
	//打开文件
	file, err := os.Open("G:\\GO\\src\\go-study\\src\\config.properties")
	if err != nil {
		panic(err)
	}

	/* defer代码块会在函数调用链表中增加一个函数调用。
	* 这个函数调用不是普通的函数调用，而是会在函数正常返回，也就是return之后添加一个函数调用。
	* 因此，defer通常用来释放函数内部变量。
	 */
	defer file.Close()

	//加载配置文件
	config, err := Load(file)
	if nil != err {
		fmt.Println(err)
		return
	}
	//读取配置
	val := config.String("key")
	fmt.Println("val:" + val)
}
