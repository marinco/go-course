package second

func Gcd(b int, c int) []Step {
	var steps = make([]Step, 0)
	for c > 0 {
		var step = Step{A: b / c, B: b, R: b % c}
		steps = append(steps, step)
		b, c = c, b%c
	}
	return steps
}
