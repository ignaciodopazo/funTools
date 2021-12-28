package funTools

// Filters the input slice with the given predicate, returning
// a new slice with the elements that satisfy it.
func Filter[T comparable](p func(T) bool, container []T) []T {
	var res []T
	for _, v := range container {
		if p(v) {
			res = append(res, v)
		}
	}
	return res
}

// Returns a new slice where its elements are the results of applying
// the given function to every element of the input slice.
func Fmap[S comparable, T comparable](f func(S) T, container []S) []T {
	// I know the result's size before computation, use that for efficiency?
	// one go routine for each result's elem in the resulting container?
	var res []T
	for _, s := range container {
		res = append(res, f(s))
	}
	return res
}

// Takes the init element and the first element of the slice and applies
// the binary function, where its result will be argument of the binary
// function along with the second element of the slice, and so on.
// Returns the result of the last application of the binary function, or
// the given init when the container is empty.
func Fold[T comparable](f func(T, T) T, init T, container []T) T {
	res := init
	for _, v := range container {
		res = f(res, v)
	}
	return res
}

func Zip[T comparable, U comparable](t []T, u []U) []Pair[T, U] {
	if len(t) == len(u) {
		res := []Pair[T, U] {}
		for i, v := range t {
			res = append(res, Pair[T, U] {v, u[i]})
		}
		return res
	} else if len(t) < len(u) {
		return Zip(t, u[:len(t)])
	} else {
		return Zip(t[:len(u)], u)
	}
}

func Partition[T comparable](p func (T) bool, container []T) ([]T, []T) {
	yes, no := []T {}, []T {}
	for _, v := range container {
		if p(v) {
			yes = append(yes, v)
		} else {
			no = append(no, v)
		}
	}
	return yes, no
}
