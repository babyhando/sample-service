package fp

func Map[T any, U any](s []T, f func(T) U) []U {
	final := make([]U, len(s))

	for i := range s {
		final[i] = f(s[i])
	}

	return final
}
