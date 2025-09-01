package main

import (
	"github.com/fibonacid/holeworm/common"
	"github.com/tmaxmax/go-sse"
	"log"
	"net/http"
	"time"
)

var t common.Transfer

func main() {
	s := &sse.Server{}

	go func() {
		m := &sse.Message{}
		m.AppendData("Hello world")

		for range time.Tick(time.Second) {
			_ = s.Publish(m)
		}
	}()

	if err := http.ListenAndServe(":8000", s); err != nil {
		log.Fatalln(err)
	}
}
