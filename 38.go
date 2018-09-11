func countAndSay(n int) string {
    
	var buf []byte = []byte{'1'}

	for n > 1 {
		buf = say(buf)
		n--
	}

	return string(buf)
}

func say(buf []byte) []byte {

	res := make([]byte, 0, len(buf)*2)

	i, j := 0, 1
	for i < len(buf) {
		
		for j < len(buf) && buf[j] == buf[i] {
			j++
		}

		res = append(res, byte(j-i+'0'), buf[i])

		i = j
	}

	return res
}
