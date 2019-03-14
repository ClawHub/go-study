package document

import (
	"fmt"
	"github.com/magiconair/properties"
	"time"
)

func DemoPropertiesV2() {
	fmt.Println("------properties v2---------")
	p := properties.MustLoadFile("G:\\GO\\src\\go-study\\src\\config.properties", properties.UTF8)
	// get values through getters
	//host := p.MustGetString("host")
	host := p.GetString("host", "localhost")
	port := p.GetInt("port", 8080)

	fmt.Println(host)
	fmt.Println(port)
	// or through Decode,注意结构体变量首字母必须大写
	type Config struct {
		Host    string        `properties:"host"`
		Port    int           `properties:"port,default=9000"`
		Accept  []string      `properties:"accept,default=image/png;image;gif"`
		Timeout time.Duration `properties:"timeout,default=5s"`
	}
	var cfg Config
	if err := p.Decode(&cfg); err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(cfg)
}
