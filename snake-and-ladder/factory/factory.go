package factory

func CreateSnakes() map[int]int {
	return map[int]int{
		17: 7,
		54: 34,
		62: 19,
		98: 79,
	}
}

func CreateLadders() map[int]int {
	return map[int]int{
		3:  22,
		5:  8,
		20: 29,
		27: 84,
		72: 91,
	}
}
