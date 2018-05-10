package log

import "fmt"

const (
	None int = iota - 1
	ErrorLogLevel
	WarningLogLevel
	InfoLogLevel
	DebugLogLevel
)

var logLevelReverseTable = map[string]int{
	"Debug":   DebugLogLevel,
	"Info":    InfoLogLevel,
	"Warning": WarningLogLevel,
	"Error":   ErrorLogLevel,
	"None":    None,
}

var logLevelTable = map[int]string{
	None:            "None",
	ErrorLogLevel:   "Error",
	WarningLogLevel: "Warning",
	InfoLogLevel:    "Info",
	DebugLogLevel:   "Debug",
}

var Error func(interface{})
var Warning func(interface{})
var Info func(interface{})
var Debug func(interface{})

var functions []*(func(interface{})) = []*(func(interface{})){&Error, &Warning, &Info, &Debug}

func LogLevelFromString(s string) int {
	return logLevelReverseTable[s]
}

func PrintLog(level string, log interface{}) {
	fmt.Printf("[%s] %v\n", level, log)
}

func SetLogLevel(l int) string {
	for i := 0; i <= l; i++ {
		level := logLevelTable[i]
		*functions[i] = func(s interface{}) {
			PrintLog(level, s)
		}
	}

	for i := l + 1; i < len(functions); i++ {
		*functions[i] = func(interface{}) {}
	}

	return logLevelTable[l]
}
