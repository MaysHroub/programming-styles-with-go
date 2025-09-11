package main

import (
	"fmt"
)

type quarantine struct {
	funcs []func(any) any // list of pure functions
}

func NewQuarantine() quarantine {
	return quarantine{
		funcs: make([]func(any) any, 0),
	}
}

func (q *quarantine) bind(f func(any) any) *quarantine {
	q.funcs = append(q.funcs, f)
	return q 
}

func (q *quarantine) execute() {
	guardCallable := func(v any) any {
		if f, ok := v.(func() any); ok {
			return f()
		}
		return v
	}
	var val any = func() any { return nil }
	for _, f := range q.funcs {
		val = f(guardCallable(val))
	}
	fmt.Println(guardCallable(val))
}
