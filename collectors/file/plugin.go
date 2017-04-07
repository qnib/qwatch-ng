package main

import (
	"C"
	"log"

	"github.com/zpatrick/go-config"
	"github.com/qnib/qwatch-ng/types"
	fTypes "github.com/qnib/qframe/types"
	"github.com/hpcloud/tail"
)

func Run(qChan fTypes.QChan, cfg config.Config) {
	log.Println("[II] Start file collector")
	fPath, err := cfg.String("collector.file.path")
	if err != nil {
		log.Println("[EE] No file path for collector.file.path set")
		return
	}
	fileReopen, err := cfg.BoolOr("collector.file.reopen", true)
	t, err := tail.TailFile(fPath, tail.Config{Follow: true, ReOpen: fileReopen})
	if err != nil {
		log.Printf("[WW] File collector failed to open %s: %s", fPath, err)
	}
	for line := range t.Lines {
		qm := qtypes.NewMessage("input", "file", "test")
		qm.Msg = line.Text
		qChan.Data.Send(qm)
	}
}
