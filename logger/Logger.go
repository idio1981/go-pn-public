package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"runtime"
)

// 输出日志等级，默认 LevelInfo
const (
	LevelError = iota
	LevelWarning
	LevelSuccess
	LevelInfo
	LevelDebug
)

// Logger : XDG 重载日志类
type Logger struct {
	level    int
	executor *log.Logger
}

func (logger *Logger) funcName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func (logger *Logger) output(level int, prefix string, format string, v ...interface{}) {
	if level > logger.level {
		return
	}

	for i := 0; i < len(v); i++ {
		if _, ok := v[i].(error); ok {
			v[i] = v[i].(error).Error()
			continue
		}

		if reflect.ValueOf(v[i]).Kind() == reflect.Ptr {
			switch reflect.Indirect(reflect.ValueOf(v[i])).Kind() {
			case reflect.Struct, reflect.Map, reflect.Slice:
				{
					v[i], _ = json.Marshal(v[i])
					continue
				}
			}
		}
		switch reflect.ValueOf(v[i]).Kind() {
		case reflect.Struct, reflect.Map, reflect.Slice:
			v[i], _ = json.Marshal(v[i])
		}
	}

	pc, _, _, _ := runtime.Caller(2)
	reg := regexp.MustCompile(`(\w*)$`)
	fn := reg.Find([]byte(runtime.FuncForPC(pc).Name()))

	msg1 := fmt.Sprintf(format, v...)
	msg2 := fmt.Sprintf("[%s] %s %s\033[0m", fn, prefix, msg1)

	logger.executor.Output(3, msg2)
}

// New Logger 构造函数
func New() *Logger {
	logger := new(Logger)
	logger.executor = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	logger.level = LevelInfo

	return logger
}

var (
	stdLogger = New()
)

// Error print log with level Error.
func Error(format string, v ...interface{}) {
	// stdLogger.output(LevelError, "\033[31;1m[E]", GenerateFmtStr(len(v)), v...)
	stdLogger.output(LevelError, "\033[31;1m[E]", format, v...)
}

// Warn print log with level Warn.
func Warn(format string, v ...interface{}) {
	stdLogger.output(LevelWarning, "\033[33;1m[W]", format, v...)
}

// Info print log with level Info.
func Info(format string, v ...interface{}) {
	stdLogger.output(LevelInfo, "\033[0;1m[I]", format, v...)
}

// Success print log with level Success.
func Success(format string, v ...interface{}) {
	stdLogger.output(LevelSuccess, "\033[32;1m[I]", format, v...)
}

// Debug print log with level Debug.
func Debug(format string, v ...interface{}) {
	stdLogger.output(LevelDebug, "\033[0m[D]", format, v...)
}

// Panic print log and exit.
func Panic(format string, v ...interface{}) {
	// format := GenerateFmtStr(len(v))

	msg1 := fmt.Sprintf(format, v...)
	panic(fmt.Sprintf("\033[31;1m[Fatal] %s\033[0m", msg1))
}

// SetLevel 设置日志输出等级
func SetLevel(level int) {
	stdLogger.level = level
}

// // GenerateFmtStr is a helper function to construct formatter string.
// func GenerateFmtStr(n int) string {
// 	return strings.Repeat("%v", n)
// }
