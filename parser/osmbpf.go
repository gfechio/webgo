//package osmbpf
package main

import (
	"encoding/json"
	//"fmt"
	"github.com/go-redis/redis"
	"github.com/qedus/osmpbf"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	//	"time"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0, // use default DB
	})

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d := osmpbf.NewDecoder(f)
	err = d.Start(runtime.GOMAXPROCS(-1)) // use several goroutines for faster decoding
	if err != nil {
		log.Fatal(err)
	}

	var nc, wc uint64
	for {
		if v, err := d.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			switch v := v.(type) {
			case *osmpbf.Node:
				ns := strconv.FormatUint(uint64(nc), 10)
				sj := onNode(v)
				err := client.Set("Node:"+ns, sj, 0).Err()
				if err != nil {
					panic(err)
				}
				nc++
			case *osmpbf.Way:
				ws := strconv.FormatUint(uint64(wc), 10)
				wj := onWay(v)
				err := client.Set("Way"+ws, wj, 0).Err()
				if err != nil {
					panic(err)
				}
				wc++
			default:
				log.Fatalf("unknown type %T\n", v)
			}
		}
	}

	// fmt.Printf("Nodes: %d, Ways: %d, Relations: %d\n", nc, wc, rc)
}

type JsonNode struct {
	Type string            `json:"type"`
	ID   int64             `json:"id"`
	Lat  float64           `json:"lat"`
	Lon  float64           `json:"lon"`
	Tags map[string]string `json:"tags"`
}

func onNode(node *osmpbf.Node) string {
	marshall := JsonNode{"node", node.ID, node.Lat, node.Lon, node.Tags}
	json, _ := json.Marshal(marshall)
	return string(json)
}

type JsonWay struct {
	Type    string            `json:"type"`
	ID      int64             `json:"id"`
	Tags    map[string]string `json:"tags"`
	NodeIDs []int64           `json:"refs"`
}

func onWay(way *osmpbf.Way) string {
	marshall := JsonWay{"way", way.ID, way.Tags, way.NodeIDs}
	json, _ := json.Marshal(marshall)
	return string(json)
}
