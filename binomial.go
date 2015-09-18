package slm

import (
	"math"

	"github.com/ematvey/gostat"
)

const BINOMIAL_SMOOTHING_FACTOR = 300

type Binomial struct {
	Parents []Binomial `json:"parents"`
	Key     string     `json:"key"`

	Positive     float64 `json:"positive"`
	Total        float64 `json:"total"`
	SamplingSize float64 `json:"sampling_size"`
}

func BinomialNew(p, a float64) SampleValue {
	return SampleValue(Binomial{Positive: p, Total: a})
}

func (d Binomial) SetKey(k string) SampleValue {
	d.Key = k
	return SampleValue(d)
}

func (d Binomial) Merge(o SampleValue) SampleValue {
	res := o.(Binomial)
	res.Positive += d.Positive
	res.Total += d.Total
	return SampleValue(res)
}

func (d Binomial) SetParents(o []SampleValue, w []float64) SampleValue {
	if len(o) != len(w) {
		panic("bad len")
	}
	v := d.Total
	ts := (v+1)/(1-math.Exp(-(v+1)/BINOMIAL_SMOOTHING_FACTOR)) - v - 1
	for i, ov := range o {
		od := ov.(Binomial)
		od.SamplingSize = w[i] * ts
		d.Parents = append(d.Parents, od)
	}
	d.SamplingSize = d.Total
	return SampleValue(d)
}

func (d Binomial) ab() (float64, float64) {
	p := d.Positive
	t := d.Total
	for _, d := range d.Parents {
		p += d.Measure() * d.SamplingSize
		t += d.SamplingSize
	}
	return p, t - p
}

func (d Binomial) Measure() float64 {
	p, n := d.ab()
	return p / (n + p)
}

func (d Binomial) Quantile(q float64) float64 {
	p, n := d.ab()
	return stat.BetaInv_CDF(p, n)(q)
}
