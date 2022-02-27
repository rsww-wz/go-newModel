package base

import "constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

type Map[T constraints.Ordered] map[T]any

type Chan[T any] chan T
