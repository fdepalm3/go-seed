package pkg

func ToPointer[T interface{}](v T) *T {
	return &v
}
