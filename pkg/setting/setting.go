package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	JwtSecret string
	TokenLife time.Duration
)


func init(){
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("fail to parse 'config/app.ini': %s", err)
	}
	LoadBase()
	LoadServer()

}

func LoadBase(){
	// 默认分区使用空字符表示
	// MustType为自动类型转换
	RunMode = Cfg.Section("").Key("run_mode").MustString("debug")
}

func LoadServer(){
	sec, err := Cfg.GetSection("server")
	if err != nil{
		log.Fatalf("fail to get section 'server':%s", err)
	}

	HTTPPort = sec.Key("http_port").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("read_timeout").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("write_timeout").MustInt(60)) * time.Second
}

