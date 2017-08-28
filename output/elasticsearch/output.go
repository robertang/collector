package elasticsearch

import (
	"github.com/robertang/collector/cncf"
	//"github.com/tguangch/stats/output"
	"github.com/robertang/collector/common"
)

type _output struct {

}

func NewOutput(oc cncf.OutputConfig) common.Output{

	return nil
}

func NewModule() common.OutputModule {
	return common.OutputModule{NewOutput}
}