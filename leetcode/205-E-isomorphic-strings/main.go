package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	set := make(map[byte]bool, 0)
	dict := make(map[byte]byte, 0)
	for i := 0; i < len(s); i++ {
		val, ok := dict[s[i]]
		if !ok {
			_, eok := set[t[i]]
			if eok {
				return false
			}
			dict[s[i]] = t[i]
			set[t[i]] = true
			continue
		}

		if val != t[i] {
			return false
		}
	}
	return true
}

func main() {
}
