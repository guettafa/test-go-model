package lib

func Help() string {
	return "Help me"
}

func DoSomething() string {
	return "I did something"
}

func AddHundredPtr(num *int) {
	*num += 100
}

func AddHundred(num int) {
	num += 100
}