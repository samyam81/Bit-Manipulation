package main

const HEX = "0123456789abcdef"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	var ans string
	count := 0
	for num != 0 && count < 8 {
		ans = string(HEX[num&0xf]) + ans
		num >>= 4 // this is basically like dividing by 16 shifting 4 bits right
		count++
	}
	return ans
}
