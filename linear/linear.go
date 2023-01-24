package linear

var value = 0

func Value() int {
	defer func() {
		value += 1
	}()
	return value
}
