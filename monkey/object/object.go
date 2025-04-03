package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
)

/* Every value in Monkey is represented as an Object
* It is an interface. Every value needs a different
* internal representation and it's easier to define
* two struct types than trying to fit booleans and
* integers into the same struct field
 */
type Object interface {
	// interface methods that return type and it's string representation
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}
