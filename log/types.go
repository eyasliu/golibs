package log

import (
	"fmt"
	"strings"
)

func FormatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}

// Debug wrapper Debug logger
func Debug(f interface{}, args ...interface{}) {
	L.Debug(FormatLog(f, args...))
}

// Info wrapper Info logger
func Info(f interface{}, args ...interface{}) {
	L.Info(FormatLog(f, args...))
}

// Warn wrapper Warn logger
func Warn(f interface{}, args ...interface{}) {
	L.Warn(FormatLog(f, args...))
}

func Print(f interface{}, args ...interface{}) {
	L.Print(FormatLog(f, args...))
}

// Printf wrapper Printf logger
func Printf(f interface{}, args ...interface{}) {
	L.Print(FormatLog(f, args...))
}

// Panic wrapper Panic logger
func Panic(f interface{}, args ...interface{}) {
	L.Panic(FormatLog(f, args...))
}

// Fatal wrapper Fatal logger
func Fatal(f interface{}, args ...interface{}) {
	L.Fatal(FormatLog(f, args...))
}

// Error wrapper Error logger
func Error(f interface{}, args ...interface{}) {
	L.Error(FormatLog(f, args...))
}

// Debugln wrapper Debugln logger
func Debugln(v ...interface{}) {
	L.Debug(fmt.Sprintln(v...))
}

// Infoln wrapper Infoln logger
func Infoln(args ...interface{}) {
	L.Info(fmt.Sprintln(args...))
}

// Warnln wrapper Warnln logger
func Warnln(args ...interface{}) {
	L.Warn(fmt.Sprintln(args...))
}

// Printfln wrapper Printfln logger
func Printfln(args ...interface{}) {
	L.Print(fmt.Sprintln(args...))
}

// Panicln wrapper Panicln logger
func Panicln(args ...interface{}) {
	L.Panic(fmt.Sprintln(args...))
}

// Fatalln wrapper Fatalln logger
func Fatalln(args ...interface{}) {
	L.Fatal(fmt.Sprintln(args...))
}

// Errorln wrapper Errorln logger
func Errorln(args ...interface{}) {
	L.Error(fmt.Sprintln(args...))
}
