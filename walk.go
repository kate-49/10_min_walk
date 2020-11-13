package walk

func Run(directions []string) bool {

	var answer bool
	var n, s, e, w int

	if len(directions) != 10 {
		answer = false
	} else {
		for _, letter := range directions {
			switch letter {
			case "n":
				n += 1
			case "s":
				s += 1
			case "e":
				e += 1
			case "w":
				w += 1
			}

			if n == s && e == w {
				answer = true
			} else {
				answer = false
			}
		}

	}
	return answer
}
