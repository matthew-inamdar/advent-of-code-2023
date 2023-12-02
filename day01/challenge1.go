package day01

func Challenge1(input string) uint {
	const (
		unset        = 10
		zeroCharCode = 0x30
	)

	var total uint

	var left, right uint = unset, unset
	for _, c := range input {
		if c == '\n' {
			total += (left * 10) + right
			left, right = unset, unset
			continue
		}
		if '0' > c || c > '9' {
			continue
		}
		right = uint(c - zeroCharCode)
		if left == unset {
			left = right
		}
	}
	total += (left * 10) + right

	return total
}
