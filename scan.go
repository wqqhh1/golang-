package main

func main() {
	f1()
}
func f1() {
	score := 90
	if score >= 60 && score <= 70 {
		println("C")
	} else if score >= 80 {
		println("A")
	}
}
