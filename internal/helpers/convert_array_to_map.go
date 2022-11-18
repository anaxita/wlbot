package helpers

func SliceToMap[T comparable](arr []T) map[T]struct{} {
	m := make(map[T]struct{}, len(arr))
	for _, id := range arr {
		m[id] = struct{}{}
	}

	return m
}
