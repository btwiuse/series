package fibonacci

var value = 0
var valueNext = 1

func Value() int {
	defer func() {
		value, valueNext = next(value, valueNext)
	}()
	return value
}

func next(n int, nn int) (int, int) {
	return nn, n + nn
}
