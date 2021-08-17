package Printlog

import (
	"fmt"
	"strconv"
	"time"
)



type NewLog struct {
	// IsDebug 默认false 那么Debug信息将不在输出，相反输出Debug 信息
	IsDebug bool

	// OutFile 将日志输入到OutFile文件中去,默认是"mmddHHMMSS.log"
	OutFile string
}

func NewPrintlog () *NewLog {
	t := time.Now().Format("20060102150405")
	t = t[4:]
	return &NewLog{
		IsDebug:false,
		OutFile:t +  ".log",
	}
}


func colorStr(colorInt int, str string) string {
	return fmt.Sprintf("\x1b[%dm%s \x1b[0m\n", colorInt, str)
}

func timeStr(level string, data []interface{}) string {
	now := time.Now()
	str := now.Format("2006-01-02 #15:04:05 ")
	str += fmt.Sprintf("[%s]", level)
	for _, v := range data {
		if value, ok := v.(string); ok {
			str += " " + value
		} else if value, ok := v.(int); ok {
			str += " " + strconv.Itoa(value)
		} else {

		}
	}
	return str
}
