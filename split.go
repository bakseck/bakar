package bakar

func split(s string, sep string) []string {
	a := len(s)
	b := len(sep)
	for i:=0; i<a-b; i++ {
		if s[i:i+b] == sep {
			s = s[:i] + " " + s[b+i:]
			a = a - b
		}
		return SplitWhiteSpaces(s)
	}
}