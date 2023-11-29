package ch191

func HammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		if num%2 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}
