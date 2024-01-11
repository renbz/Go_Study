package main

func minLength(s string) int {
	var stack []byte
	for i := range s {
		stack = append(stack, s[i])
		m := len(stack)
		if m >= 2 && (string(stack[m-2:]) == "AB" || string(stack[m-2:]) == "CD") {
			stack = stack[:m-2]
		}
	}
	return len(stack)
}
