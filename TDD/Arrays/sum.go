package main

func Sum2(numbers [5]int)int{
	sum := 0

	//range returns two values, the inddex and the val
	for _, number := range numbers{
		sum += number
	}
	return sum
}
