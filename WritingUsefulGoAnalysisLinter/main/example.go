package main

import "log"

// myLog outputs with specific prefix.
func myLog(format string, args ...interface{}) {
	const prefix = "[my] "
	log.Printf(prefix+format, args...)
}
