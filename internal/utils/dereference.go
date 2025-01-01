package utils

func DerefString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func DerefInt(ptr *int) int {
	if ptr == nil {
		return 0
	}
	return *ptr
}

func DerefBool(ptr *bool) bool {
	if ptr == nil {
		return false
	}
	return *ptr
}

func DerefFloat64(ptr *float64) float64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}
