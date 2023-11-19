package fn

func MapSlice[I any, O any](slice []I, mapFn func(I) O) []O {
	sliceOut := make([]O, len(slice))
	for i, e := range slice {
		sliceOut[i] = mapFn(e)
	}
	return sliceOut
}

func FilterSlice[T any](slice []T, filterFn func(T) bool) []T {
	sliceOut := []T{}
	for _, e := range slice {
		if filterFn(e) {
			sliceOut = append(sliceOut, e)
		}
	}
	return sliceOut
}
