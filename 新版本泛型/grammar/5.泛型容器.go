package grammar

import "constraints"

// List 泛型切片
type List[T any] []T

// Map 泛型map
type Map[T constraints.Ordered] map[T]any
