package main

import (
	"C"
	"log"
	"fmt"
	"strings"
	"time"

	"github.com/zpatrick/go-config"
	"github.com/qnib/qwatch-ng/types"
	fTypes "github.com/qnib/qframe/types"
	"github.com/qnib/qframe/utils"

	"github.com/OwnLocal/goes"
	"net/url"
)

func createIndex(conn *goes.Connection, idx string) error {
	// Create an index
	mapping := map[string]interface{}{
		"settings": map[string]interface{}{
			"index.number_of_shards":   1,
			"index.number_of_replicas": 0,
		},
		"mappings": map[string]interface{}{
			"_default_": map[string]interface{}{
				"_source": map[string]interface{}{
					"enabled": true,
				},
				"_all": map[string]interface{}{
					"enabled": false,
				},
			},
		},
	}

	resp, err := conn.CreateIndex(idx, mapping)
	_ = resp
	log.Printf("%v\n", resp)
	return err
}

func indexLog(conn *goes.Connection, idx string, log qtypes.Message) error {
	d := goes.Document{
		Index: idx,
		Type:  "log",
		Fields: map[string]interface{}{
			"Timestamp": log.Time.Format("2006-01-02T15:04:05.999999-07:00"),
			"msg":       log.Msg,
			"source":    log.Source,
			"type":      log.Type,
			"host":      log.Host,
		},
	}
	extraArgs := make(url.Values, 1)
	//extraArgs.Set("ttl", "86400000")
	response, err := conn.Index(d, extraArgs)

	//_ = response
	fmt.Printf("%s | %s\n", d, response.Error)
	return err
}

// Run fetches everything from the Data channel and flushes it to stdout
func Run(qChan fTypes.QChan, cfg config.Config) {
	bg := qChan.Data.Join()
	inStr, err := cfg.String("handler.elasticsearch.inputs")
	if err != nil {
		inStr = ""
	}
	inputs := strings.Split(inStr, ",")
	host, _ := cfg.StringOr("handler.elasticsearch.host", "kibana_backend")
	port, _ := cfg.StringOr("handler.elasticsearch.port", "9200")
	idxForm, _ := cfg.StringOr("handler.elasticsearch.index-format", "logstash-2017-10-24")
	idx := time.Now().Format(idxForm)
	conn := goes.NewConnection(host, port)
	createIndex(conn, idx)
	for {
		val := bg.Recv()
		switch val.(type) {
		case qtypes.Message:
			qm := val.(qtypes.Message)
			if len(inputs) != 0 && ! qutils.IsInput(inputs, qm.Source) {
				continue
			}
			indexLog(conn, idx, qm)
			fmt.Printf("%s %-7s sType:%-6s sName:[%d]%-10s %s:%s\n", qm.TimeString(), qm.LogString(), qm.Type, qm.SourceID, qm.Source, qm.Name, qm.Msg)
		}
	}
}
