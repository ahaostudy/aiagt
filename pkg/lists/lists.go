package lists

func FlatMap[T, E any](list []T, m func(t T) []E) []E {
	var result []E

	for _, v := range list {
		result = append(result, m(v)...)
	}

	return result
}

func Map[T, E any](list []T, m func(T) E) []E {
	var result []E

	for _, v := range list {
		result = append(result, m(v))
	}

	return result
}

func Filter[T any](t []T, f func(T) bool) []T {
	var result []T

	for _, v := range t {
		if f(v) {
			result = append(result, v)
		}
	}

	return result
}

func AnyList[T any](list []T) []any {
	if list == nil {
		return nil
	}

	result := make([]any, len(list))
	for i, item := range list {
		result[i] = item
	}

	return result
}

func InsertSlice[T any](slice []T, index int, elements ...T) []T {
	newLen := len(slice) + len(elements)

	if newLen > cap(slice) {
		newSlice := make([]T, newLen)
		copy(newSlice, slice[:index])
		copy(newSlice[index:], elements)
		copy(newSlice[index+len(elements):], slice[index:])
		return newSlice
	}

	newSlice := slice[:newLen]
	copy(newSlice[index+len(elements):], slice[index:])
	copy(newSlice[index:], elements)
	return newSlice
}
