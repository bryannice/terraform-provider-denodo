package denodo

func TenaryString(condition bool, val1 string, val2 string) string {
	if condition {
		return val1
	}
	return val2
}
