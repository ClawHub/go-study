package document

import (
	"github.com/magiconair/properties"
)

//变量
var Properties *properties.Properties

//初始化
func init() {
	//加载配置文件
	Properties = properties.MustLoadFile("G:\\GO\\src\\go-study\\src\\config.properties", properties.UTF8)
	// or multiple files
	//p = properties.MustLoadFiles([]string{
	//	"${HOME}/config.properties",
	//	"${HOME}/config-${USER}.properties",
	//}, properties.UTF8, true)

	// or from a map
	//p = properties.LoadMap(map[string]string{"key": "value", "abc": "def"})

	// or from a string
	//p = properties.MustLoadString("key=value\nabc=def")

	// or from a URL
	//p = properties.MustLoadURL("http://host/path")

	// or from multiple URLs
	//p = properties.MustLoadURL([]string{
	//	"http://host/config",
	//	"http://host/config-${USER}",
	//}, true)

	// or from flags
	//p.MustFlag(flag.CommandLine)
}
