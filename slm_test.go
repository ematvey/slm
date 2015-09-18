package slm

import "testing"

func Test1(t *testing.T) {
	var fbg = FallbackGraph{}
	fbg.Add(
		[]string{"a", "b", "c"},
		[][]string{
			{"a", "b"},
		},
		[]float64{1},
	)
	fbg.Add(
		[]string{"a", "b"},
		[][]string{
			{"a"},
			{"b"},
		},
		[]float64{0.5, 0.5},
	)

	var dts = []DataSlice{
		DataSlice{
			Features: FeatureSet{
				"a": "a1",
				"b": "b1",
				"c": "c1",
			},
			Value: BinomialNew(10, 100),
		},
		DataSlice{
			Features: FeatureSet{
				"a": "a1",
				"b": "b1",
				"c": "c2",
			},
			Value: BinomialNew(3, 100),
		},
		DataSlice{
			Features: FeatureSet{
				"a": "a1",
				"b": "b2",
				"c": "c3",
			},
			Value: BinomialNew(1, 100),
		},
		DataSlice{
			Features: FeatureSet{
				"a": "a1",
				"b": "b2",
				"c": "c1",
			},
			Value: BinomialNew(50, 100),
		},
	}

	var c = MakeCollector(dts, fbg)

	fs := FeatureSet{
		"a": "a1",
		"b": "b1",
	}
	k := fs.Key()

	v := c[k]

	if v.Measure() != 0.08076431631028652 {
		t.FailNow()
	}
}
