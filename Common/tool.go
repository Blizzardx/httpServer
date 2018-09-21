package common

import (
	"log"
	"runtime/debug"
)

func SafeCall(f func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(string(debug.Stack()))
		}
	}()
	f()
}
