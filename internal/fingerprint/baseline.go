package fingerprint

type Baseline struct {
	Counts map[string]int
}

func NewBaseLine() *Baseline {
	return &Baseline{
		Counts: make(map[string]int),
	}
}

func (b *Baseline) Add(fingerprint string) {
	b.Counts[fingerprint]++
}

func (b *Baseline) Dominant() string {
	var dominant string
	maxCount := 0

	for fingerprint, count := range b.Counts {
		if count > maxCount {
			dominant = fingerprint

			maxCount = count
		}
	}
	return dominant
}
