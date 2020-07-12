package main

type Logger struct {
	log map[string]int
}

/** Initialize your data structure here. */
func Constructor() Logger {
	return Logger{log: map[string]int{}}
}

/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
  If this method returns false, the message will not be printed.
  The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if oldTime, found := this.log[message]; !found || timestamp-oldTime >= 10 {
		this.log[message] = timestamp
		return true
	}
	return false
}
