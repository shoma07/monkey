package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

// 整数
type Integer struct {
	Value int64
}

func (o *Integer) Type() ObjectType {
	return INTEGER_OBJ
}
func (o *Integer) Inspect() string {
	return fmt.Sprintf("%d", o.Value)
}

// 真偽値
type Boolean struct {
	Value bool
}

func (o *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}
func (o *Boolean) Inspect() string {
	return fmt.Sprintf("%t", o.Value)
}

// null
type Null struct{}

func (o *Null) Type() ObjectType {
	return NULL_OBJ
}
func (o *Null) Inspect() string {
	return "null"
}

// 戻り値 他のオブジェクトのラッパー
type ReturnValue struct {
	Value Object
}

func (o *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}
func (o *ReturnValue) Inspect() string {
	return o.Value.Inspect()
}
