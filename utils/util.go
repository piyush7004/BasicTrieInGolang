package utils

//IsJustAlphabets checks if the provided string has just alphabets or not
func IsJustAlphabets(s string) bool {
	for _, b := range []byte(s) {
		if b < 65 || b > 122 {
			return false
		} else if b > 90 && b < 97 {
			return false
		} else {
			continue
		}
	}
	return true

	// 	if b >= 97 && b <= 122 {
	// 		continue
	// 	} else if b >= 65 && b <= 90 {
	// 		continue
	// 	} else {
	// 		return false
	// 	}
	// }
	// return true
}

//GetCaseInsensitive returns the lower cased value of provided string
func GetCaseInsensitive(s string) string {
	out := []byte{}
	for _, b := range []byte(s) {
		if b < 97 {
			out = append(out, b+32)
			continue
		}
		out = append(out, b)
	}
	return string(out)
}
