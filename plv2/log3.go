// START OMIT

package log

// Logs an alert and panics the goroutine. Values will be printed one after the
// other. (Note this is internally logged at seelog's Critical level).
func Pnc(traceId string, v ...interface{}) {
	...
	panic(msgRaw)
}

// Logs an alert and panics the goroutine. Values will be printed according to
// the format string. (Note this is internally logged at seelog's Critical level).
func Pncf(traceId, format string, args ...interface{}) {
	...
	panic(msgRaw)
}

// END OMIT
