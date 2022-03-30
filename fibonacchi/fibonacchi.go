package fibonacchi

func Loop(num int) (int, error) {

	if num == 0 {
		return 0, nil
	}

	if num == 1 {
		return 1, nil
	}

	result := 0
	a := 0
	b := 0

	for x := 0; x < num-1; x++ {

		if a == 0 && b == 0 { //первый проход
			a = x
			b = x + 1
		}

		result = a + b
		b = a
		a = result

	}

	return result, nil
}
