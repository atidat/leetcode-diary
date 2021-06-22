package main

func main() {

}

/*
1 <= binary.length <= 105
binary 仅包含 '0' 和 '1'
*/
func maximumBinaryString(binary string) string {
	// 00 -> 10
	// 10 -> 01
	binbytes := []byte(binary)

	for i := 0; i < len(binary); {
		
	}

	return string(binbytes)
}

func ope(in1, in2 byte) (byte, byte) {
	if in1 == '0' && in2 == '0' {
		return '1', '0'
	}
	if in1 == '1' && in2 == '0' {
		return '0', '1'
	}
	return in1, in2
}