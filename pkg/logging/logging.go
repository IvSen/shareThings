package logging

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func Init(ctx context.Context) {
	logrusLogger := logrus.New()
	logrusLogger.SetLevel(logrus.TraceLevel)
	logrusLogger.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s:%d", filename, f.Line), fmt.Sprintf("%s()", f.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}
	logrusLogger.SetOutput(os.Stdout)
	logrusLogger.SetReportCaller(true)
	logrusLogger.WithContext(ctx)

	e = logrus.NewEntry(logrusLogger)
}
