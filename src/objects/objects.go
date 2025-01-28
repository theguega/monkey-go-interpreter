// objects/objects.go

package objects

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}
