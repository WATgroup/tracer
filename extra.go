package tracer

import (
	"fmt"
)


func FatalV(z ...interface{}) {
	// FATAL: ...
	_puts(x2, k_FATAL)
	fmt.Fprintln(x2, z...)
}

func Fatalf(f string,z ...interface{}) {
	// FATAL: ...
	_puts(x2, k_FATAL)
	fmt.Fprintf(x2, f, z...)
}

func ErrorV(z ...interface{}) {
	_puts(x2, k_ERROR)
	fmt.Fprintln(x2, z...)
}

func Errorf(f string, z ...interface{}) {
	_puts(x2, k_ERROR)
	fmt.Fprintf(x2, f, z...)
}


func WarnV(z ...interface{}) {
	_puts(x2, k_WARN)
	fmt.Fprintln(x2, z...)
}

func Warnf(f string, z ...interface{}) {
	_puts(x2, k_WARN)
	fmt.Fprintf(x2, f, z...)
}


func InfoV(z ...interface{}) {
	_puts(x2, k_INFO)
	fmt.Fprintln(x2, z...)
}

func Infof(f string, z ...interface{}) {
	_puts(x2, k_INFO)
	fmt.Fprintf(x2, f, z...)
}


func NoticeV(z ...interface{}) {
	_puts(x2, k_NOTICE)
	fmt.Fprintln(x2, z...)
}

func Noticef(f string, z ...interface{}) {
	_puts(x2, k_NOTICE)
	fmt.Fprintf(x2, f, z...)
}


func Putf(x string,z ...interface{}) {
	fmt.Fprintf(x1, x, z...)
}
