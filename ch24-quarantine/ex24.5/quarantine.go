package main

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
	var val any = func() any { return nil }
	for _, f := range q.funcs {
		val = f(guardCallable(val))
	}
	guardCallable(val) // this is for the printSorted function
}

func guardCallable(v any) any {
	if f, ok := v.(func() any); ok {
		return f()
	}
	return v
}
