package slm

type fallbackPlane struct {
	FeatureSets [][]string
	Weights     []float64
}

type FallbackGraph map[string]fallbackPlane

func (f FallbackGraph) Add(
	node_features []string,
	parent_features [][]string,
	weights []float64,
) {
	var ws float64
	for _, w := range weights {
		ws += w
	}
	for i := range weights {
		weights[i] /= ws
	}
	f[mask(node_features)] = fallbackPlane{FeatureSets: parent_features, Weights: weights}
}

func (f FallbackGraph) Find(node_features []string) fallbackPlane {
	v, ok := f[mask(node_features)]
	if ok {
		return v
	} else {
		comb := combinations(node_features, len(node_features)-1)
		ws := make([]float64, len(comb))
		for i := range ws {
			ws[i] = float64(1.0) / float64(len(ws))
		}
		v = fallbackPlane{FeatureSets: comb, Weights: ws}
		return v
	}
}
