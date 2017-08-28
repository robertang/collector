package common

import "github.com/robertang/collector/cncf"

type MetricModule struct {
	NewMetric 	func(config cncf.MetricConfig) Metric
}

type OutputModule struct {
	NewOutput 	func(oc cncf.OutputConfig) Output
}

type Output interface {
	Output(text interface{}) error
}

type Metric interface {
	Collect(host string) (interface{}, error)
	Start(output Output)
}

