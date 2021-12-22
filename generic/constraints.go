package generic

type Comparable[T any] interface {
	CompareTo(other T) int
}
