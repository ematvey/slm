package slm

import (
	"sort"
	"strings"
)

func unique(a []string) []string {
	if len(a) <= 1 {
		return a
	}
	b := a
	sort.Strings(b)
	r := make([]string, 0, len(b))
	r = append(r, b[0])
	for i := 1; i < len(b); i++ {
		if b[i] != r[len(r)-1] {
			r = append(r, b[i])
		}
	}
	return r
}

func combinations(pool []string, n int) [][]string {
	var result [][]string
	pool = unique(pool)
	delim := "~*~"
	if n <= 1 {
		result = make([][]string, 0)
		for _, el := range pool {
			result = append(result, []string{el})
		}
	} else {
		x := make([]string, 0)
		for i := range pool {
			for _, tail := range combinations(pool[1:len(pool)], n-1) {
				sub := []string{pool[i]}
				sub = append(sub, tail...)
				sub = unique(sub)
				if len(sub) == n {
					x = append(x, strings.Join(sub, delim))
				}
			}
		}
		x = unique(x)
		for _, el := range x {
			result = append(result, strings.Split(el, delim))
		}
	}
	return result
}

func mask(m []string) string {
	s := make([]string, len(m))
	copy(s, m)
	sort.Strings(s)
	return strings.Join(s, ";")
}
