package stringzeug

// ContainsString checks if the given slice contains the supplied needle string.
func ContainsString(slice []string, needle string) bool {
	for _, s := range slice {
		if s == needle {
			return true
		}
	}

	return false
}

// IndexOf iterates the slice from front to back and if the needle is found then the index position is returned. If the
// needle is not found then -1 is returned.
func IndexOf(slice []string, needle string) int {
	for idx, s := range slice {
		if s == needle {
			return idx
		}
	}

	return -1
}

// LastIndexOf iterates the slice from back to front and if the needle is found then the index position is returned. If
// the needle is not found then -1 is returned.
func LastIndexOf(slice []string, needle string) int {
	for idx := len(slice) - 1; idx >= 0; idx-- {
		if slice[idx] == needle {
			return idx
		}
	}

	return -1
}

func IsEmpty(slice []string) bool {
	return len(slice) == 0
}

func IsNotEmpty(slice []string) bool {
	return len(slice) != 0
}

// RemoveAtIndex removes a single element from the slice at the given index. The function returns the element at the
// index (if it exists) and a bool that indicates if the operation succeeded.
func RemoveAtIndex(slice []string, index int) ([]string, bool) {
	if index < 0 || IsEmpty(slice) || index > len(slice) {
		return slice, false
	}

	return append(slice[:index], slice[index+1:]...), true
}

// RemoveIf removes multiple elements from the slice if they match the predicate.
func RemoveIf(slice []string, p func(s string) bool) []string {
	if IsEmpty(slice) {
		return slice
	}

	result := make([]string, 0)
	for _, s := range slice {
		if !p(s) {
			result = append(result, s)
		}
	}

	return result
}
