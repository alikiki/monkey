package object

import "fmt"

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

func (e *Environment) SeeStore() {
	for k, v := range e.store {
		fmt.Println(k, ": ", v)
	}
	if e.outer != nil {
		fmt.Println("---")
		e.outer.SeeStore()
	}
}

func (e *Environment) OuterExists() bool {
	return e.outer != nil
}
