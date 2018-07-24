package log

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

const (
	sys_log_root_path = "./log"
	Fatal             = 0
	Err               = 1
	Warn              = 2
	Notice            = 3
	Info              = 4
	Debug             = 5
)

var DEBUG = Err

func Log(level int, format string, args ...interface{}) {
	DEBUG_tmp, err := strconv.Atoi(os.Getenv("DEBUG"))
	//从环境变量中取DEBUG级别
	if err == nil && DEBUG_tmp < 9 {
		DEBUG = DEBUG_tmp
	}
	if level > DEBUG && level != Info {
		return
	}
	date_time_format := time.Now().Format("20060102150405")
	date_str := date_time_format[0:8]
	_, exist := os.Stat(sys_log_root_path)
	if exist != nil {
		os.MkdirAll(sys_log_root_path, os.ModePerm)
	}
	file, _ := os.OpenFile(sys_log_root_path+"/"+date_str+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	var output_log_format = ""
	switch level {
	case Fatal:
		output_log_format += "[FATAL]  " + format + "\n"
	case Err:
		output_log_format += "[ERROR]  " + format + "\n"
	case Warn:
		output_log_format += "[WARN]   " + format + "\n"
	case Notice:
		output_log_format += "[NOTICE] " + format + "\n"
	case Info:
		output_log_format += "[INFO]   " + format + "\n"
	case Debug:
		output_log_format += "[DEBUG]  " + format + "\n"
	}
	var output_log = "[" + date_time_format + "]"
	output_log += Print(output_log_format, args)
	fmt.Fprint(file, output_log)
	file.Close()
}

func Print(format string, args []interface{}) string {
	p := newPrinter()
	p.DoPrintf(format, args)
	printStr := string(p.buf)
	p.free()
	return printStr
}

func File_line() string {
	_, fileName, fileLine, ok := runtime.Caller(1)
	var s string
	if ok {
		s = fmt.Sprintf("%s:%d", fileName, fileLine)
	} else {
		s = ""
	}
	return s
}
