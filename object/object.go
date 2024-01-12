package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

type ObjectType string
type BuiltinFunction func(args ...Object) Object

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

type String struct {
	Value string
}

type Boolean struct {
	Value bool
}

type ReturnValue struct {
	Value Object
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

type Array struct {
	Elements []Object
}

type Builtin struct {
	Fn BuiltinFunction
}

type Null struct{}

type Error struct {
	Message string
}

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	ARRAY_OBJ        = "ARRAY"
	BUILTIN_OBJ      = "BUILTIN"
)

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string {
	return s.Value
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string {
	return "null"
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

func (a *Array) Type() ObjectType { return ARRAY_OBJ }
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string {
	return "builtin function"
}
