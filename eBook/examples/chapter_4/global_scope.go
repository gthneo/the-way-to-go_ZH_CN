package main

var A = "G"

func main() {
	n()
	m()
	n()
}
func n() {
	print(A)
}
func m() {
	A = "O"
	print(A)
}
