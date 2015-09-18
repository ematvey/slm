package slm

import (
	"fmt"
	"sort"
	"strings"
)

type FeatureSet map[string]string

func (f FeatureSet) Features() []string {
	fs := make([]string, 0, len(f))
	for k, _ := range f {
		fs = append(fs, k)
	}
	sort.Strings(fs)
	return fs
}

func (f FeatureSet) Key() string {
	subkeys := make([]string, 0, len(f))
	for k, v := range f {
		if v == "" {
			continue
		}
		subkeys = append(subkeys, fmt.Sprintf("%+v=%+v", k, v))
	}
	sort.Strings(subkeys)
	key := strings.Join(subkeys, "&")
	return key
}

func (f FeatureSet) ReduceTo(keep []string) FeatureSet {
	var dummy_f bool
	feat := make(FeatureSet, len(f))
	for k, v := range f {
		dummy_f = true
		for _, f := range keep {
			if k == f {
				dummy_f = false
				break
			}
		}
		if !dummy_f {
			feat[k] = v
		}
	}
	return feat
}

type SampleValue interface {

	// Add two values
	Merge(SampleValue) SampleValue

	SetKey(string) SampleValue

	// List of Values to merge with; list of merge Weights
	SetParents([]SampleValue, []float64) SampleValue

	// Calculate final value
	Measure() float64
}

type DataSlice struct {
	Features FeatureSet
	Value    SampleValue
}
