package main

import (
	"C"

	"github.com/zpatrick/go-config"
	"github.com/qnib/qwatch-ng/types"
	fTypes "github.com/qnib/qframe/types"
	"github.com/qnib/qframe/utils"
	"fmt"
	"strings"
)

// Run fetches everything from the Data channel and flushes it to stdout
func Run(qChan fTypes.QChan, cfg config.Config) {
	bg := qChan.Data.Join()
	inStr, err := cfg.String("handler.log.inputs")
	if err != nil {
		inStr = ""
	}
	inputs := strings.Split(inStr, ",")
	for {
		val := bg.Recv()
		switch val.(type) {
		case qtypes.Message:
			qm := val.(qtypes.Message)
			if len(inputs) != 0 && ! qutils.IsInput(inputs, qm.Source) {
				continue
			}
			fmt.Printf("%s %-7s sType:%-6s sName:[%d]%-10s %s:%s\n", qm.TimeString(), qm.LogString(), qm.Type, qm.SourceID, qm.Source, qm.Name, qm.Msg)
		}
	}
}


