/*
Package log provides provides a minimal interface for structured logging
*/
package log

import (
	"log"
	"os"

	gklog "github.com/go-kit/kit/log"
	gklevels "github.com/go-kit/kit/log/levels"
)

// Logger provides provides a minimal interface for structured logging
// It supplies leveled logging functionw which create a log event from keyvals,
// a variadic sequence of alternating keys and values.
//TODO: replace panic as it would result in cyclic errors
type Logger interface {
	Debug(keyvals ...interface{})
	Info(keyvals ...interface{})
	Error(keyvals ...interface{})
	Warn(keyvals ...interface{})
	Crit(keyvals ...interface{})

	With(keyvals ...interface{}) Logger
}

// New takes the name of the file as an argument, creates the file and returns
// a leveled logger that logs to the file
func New(file string) Logger {
	fw, err := GetFile(file)
	if err != nil {
		log.Println("error opening log file")
		log.Fatal(err)
	}
	l := gklog.NewLogfmtLogger(fw)
	kitlevels := gklevels.New(
		l,

		// Fudge values so that switching between debug/info levels does not
		// mess with the log justification
		gklevels.DebugValue("dbug"),
		gklevels.ErrorValue("errr"),
	)

	kitlevels = kitlevels.With("ts", gklog.DefaultTimestampUTC)

	return levels{kitlevels}
}

//GetFile opens a file in read/write to append data to it
func GetFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

type levels struct {
	kit gklevels.Levels
}

func (l levels) Debug(keyvals ...interface{}) {
	if err := l.kit.Debug(keyvals...); err != nil {
		panic(err)
	}
}

func (l levels) Info(keyvals ...interface{}) {
	if err := l.kit.Info(keyvals...); err != nil {
		panic(err)
	}
}

func (l levels) Error(keyvals ...interface{}) {
	if err := l.kit.Error(keyvals...); err != nil {
		panic(err)
	}
}
func (l levels) Warn(keyvals ...interface{}) {
	if err := l.kit.Warn(keyvals...); err != nil {
		panic(err)
	}
}
func (l levels) Crit(keyvals ...interface{}) {
	if err := l.kit.Crit(keyvals...); err != nil {
		panic(err)
	}
}

func (l levels) With(keyvals ...interface{}) Logger {
	return levels{l.kit.With(keyvals...)}
}
