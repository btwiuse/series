package flipper

var value = 0

func Value() int {
	defer func() {
		value = 1 - value
	}()
	return value
}
