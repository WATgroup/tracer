package tracer

import (
	"fmt"
)

func Trace(n uint, s string) {
	if n > threshold {
		return
	}
	fmt.Fprintf(x1, k_TRACE, n)
	_putln(x1, s)
}


func Tracef(n uint, s string, x ...interface{}) {
	if n > threshold {
		return
	}
	fmt.Fprintf(x1, k_TRACE, n)
	fmt.Fprintf(x1, s, x...)
	x1.Write([]byte("\n"))
}

func TraceV(n uint, x ...interface{}) {
	if n > threshold {
		return
	}
	fmt.Fprintf(x1, k_TRACE, n)
	fmt.Fprintln(x1, x...)
	x1.Write([]byte("\n"))
}


func Debugf(s string, x ...interface{}) {
	_puts(x2, k_DEBUG)
	fmt.Fprintf(x2, s, x...)
	x2.Write([]byte("\n"))
}

func DebugV(x ...interface{}) {
	_puts(x1, k_DEBUG)
	fmt.Fprintln(x2, x...)
	x1.Write([]byte("\n"))
}
