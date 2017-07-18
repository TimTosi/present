// START OMIT

package log

// Logs for apis, contains the api name, duration and httpCode
func ApiInfof(traceId, Api string, durationInMs, httpCode int64, v ...interface{}) {
	...
}

// Logs for apis, contains the api name, duration and httpCode
func ApiErrf(traceId, Api string, httpCode int64, v ...interface{}) {
	...
}

// END OMIT
