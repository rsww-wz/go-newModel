package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// 创建自定义logger
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

// 配置log
func init() {
	log.SetPrefix("RS:")

	// 设置日志文件
	f, err := os.OpenFile("./rs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	// 日志输出到文件中
	log.SetOutput(f)

	// 设置flag格式 == 日期 时间 秒数 文件名和行号
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	// 自定义logger
	Trace = log.New(ioutil.Discard,
		"TRACE",
		log.Ldate|log.Ltime|log.Llongfile)

	Info = log.New(os.Stdout,
		"INFO",
		log.Ldate|log.Ltime|log.Llongfile)

	Warning = log.New(os.Stdout,
		"WARNING",
		log.Ldate|log.Ltime|log.Llongfile)

	Error = log.New(io.MultiWriter(f,os.Stderr),
		"ERROR",
		log.Ldate|log.Ltime|log.Llongfile)
}

func main() {
	log.Println("1234")
	////log.Panicln("1234")
	//log.Fatalln("1234")

}
