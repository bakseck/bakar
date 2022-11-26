package bakar

func isSafe(a, b int, pos [8]int) bool {
	for i:=0; a<=a; i++ {
		t := pos[i]
		if t == b || t == b - (a-i) || t == b + (a-i) {
			return false
		}
	}
	return true
}

func resolveThePositionOftheQueens()