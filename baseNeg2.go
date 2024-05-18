package main

func baseNeg2(n int) string {
	// 2。     010
	// 4 - 2   110
	// 5
	// 4 + 1  101
	// 6      110
	// 4 + 2。  1010
	h := make([]uint8, 0)
	for n > 0 {
		if (n & 1) > 0 {
			h = append(h, 1)
		} else {
			h = append(h, 0)
		}
		n = n >> 1
	}
	s := uint8(0)
	for i := 0; i < len(h); i += 1 {
		t := h[i] + s
		if t >= 2 {
			h[i] = t % 2
			s = 1
		}
		if i%2 == 1 {
			s += 1
		}
	}
	if s == 1 {
		h = append(h, 1)
	}
	b := []byte{}
	for i := len(h) - 1; i >= 0; i -= 1 {
		b = append(b, h[i])
	}
	return string(b)

}
