package util

// IfElse evaluates a condition, if true returns the first parameter otherwise the second
func IfElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
