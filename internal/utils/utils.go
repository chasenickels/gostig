package utils

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}
