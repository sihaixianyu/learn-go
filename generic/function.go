package generic

type Student struct {
	Age int
}

func (s *Student) CompareTo(other *Student) int {
	return s.Age - other.Age
}

func Max[T Comparable[T]](o1, o2 T) T {
	//if o1.CompareTo(o2) > 0 {
	//	return o1
	//} else {
	//	return o2
	//}
	return o1
}
