package slm

// Collector is a container for each node in Sample tree
type Collector map[string]SampleValue

func MakeCollector(data []DataSlice, fbg FallbackGraph) Collector {
	if len(data) < 2 {
		return nil
	}
	res := make(Collector)

	var feature_names = data[0].Features.Features()

	var root_value SampleValue
	for i := range data {
		if i == 0 {
			root_value = data[i].Value
		} else {
			root_value = root_value.Merge(data[i].Value)
		}
	}
	root_key := data[0].Features.ReduceTo(nil).Key()
	res[root_key] = root_value.SetKey(root_key)

	for i := 1; i <= len(feature_names); i++ {
		for _, comb := range combinations(feature_names, i) {
			groups := make(map[string]DataSlice)

			for _, d := range data {
				reduced_features := d.Features.ReduceTo(comb)
				key := reduced_features.Key()
				if e, ok := groups[key]; !ok {
					groups[key] = DataSlice{
						Features: reduced_features,
						Value:    d.Value,
					}
				} else {
					e.Value = e.Value.Merge(d.Value)
					groups[key] = e
				}
			}

			for key, gr := range groups {
				if key == root_key {
					continue
				}
				v := gr.Value
				v = v.SetKey(key)
				switch {

				// smooth with base rate
				case len(comb) == 1:
					res[key] = v.SetParents(
						[]SampleValue{root_value},
						[]float64{1},
					)

				// smooth with neighbors
				case len(comb) > 1:
					var parent_v []SampleValue
					var parent_w []float64
					plane := fbg.Find(comb)
					for i := range plane.FeatureSets {
						key := gr.Features.ReduceTo(plane.FeatureSets[i]).Key()
						if cv, ok := res[key]; ok {
							parent_v = append(parent_v, cv)
							parent_w = append(parent_w, plane.Weights[i])
						}
					}
					res[key] = v.SetParents(parent_v, parent_w)
				}
			}
		}
	}
	return res
}
