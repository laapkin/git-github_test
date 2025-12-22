package generators

type Element = any

type General interface {
	Next() bool
	Val() Element
}
