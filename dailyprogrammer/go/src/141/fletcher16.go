package main

func Fletcher16(input string) uint16 {
	var sum1 uint16 = 0
	var sum2 uint16 = 0
	
	for _, char := range input {
		sum1 = (sum1 + uint16(char)) % 255
		sum2 = (sum2 + sum1) % 255
	}
	
	return (sum2 << 8) | sum1
}