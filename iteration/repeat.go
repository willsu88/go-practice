package main

func Repeat(c string, n int) string {

	var res string
	for i := 0; i < n; i++ {
		res += c
	}
	return res
}
