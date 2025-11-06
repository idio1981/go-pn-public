package test

import (
	"os"
	"regexp"
	"runtime"

	"github.com/idio1981/go-pn-public/logger"
)

func perpare() {
	os.Chdir("../")
	logger.SetLevel(logger.LevelDebug)

	pc, _, _, _ := runtime.Caller(1)
	reg := regexp.MustCompile(`(\w*)$`)
	fn := reg.Find([]byte(runtime.FuncForPC(pc).Name()))

	logger.Info("%s", string(fn))
}
