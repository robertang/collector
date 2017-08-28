package main

import (
	"flag"
	"github.com/robertang/collector/cncf"
	_ "github.com/robertang/collector/output"
	_ "github.com/robertang/collector/metric"
	"github.com/robertang/collector/collect"
	"fmt"
	"encoding/json"
)

var (
	configFile = flag.String("conf", "./src/github.com/robertang/collector/cmd/proxy/redis.yml", "location of config file")
)

func main(){
	flag.Parse()
	ymls := cncf.Parse(*configFile)

	_j, _ := json.Marshal(ymls)
	fmt.Println(string(_j))

	collect.Collect(ymls)
	<- make (chan bool)
}


