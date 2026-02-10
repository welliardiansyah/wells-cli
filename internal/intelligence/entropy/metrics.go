package entropy

type Metrics struct {
	FanOut          int
	LayerViolations int
	LinesOfCode     int
}

func (m Metrics) Score() float64 {
	score := float64(m.FanOut)*0.3 +
		float64(m.LayerViolations)*0.5 +
		float64(m.LinesOfCode)*0.2

	if score > 1 {
		return 1
	}
	return score
}
