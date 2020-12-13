package main

var (
	c   []byte
	res []string
)

func permutation(s string) []string {
	if len(s) == 0 {
		return []string{s}
	}
	res = []string{}
	c = []byte(s)
	dfsPer(0)
	return res
}

func dfsPer(x int) {
	if x == len(c)-1 {
		res = append(res, string(c))
		return
	}
	m := map[byte]bool{}
	for i := x; i < len(c); i++ {
		if _, ok := m[c[i]]; ok {
			continue
		}
		m[c[i]] = true
		c[x], c[i] = c[i], c[x]
		dfsPer(x + 1)
		c[x], c[i] = c[i], c[x]
	}
}
