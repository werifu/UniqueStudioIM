package gotest

import (
	"testing"
	"time"
)

func Test_ticker(t *testing.T){
	ticker := time.NewTicker(time.Second * 1)
	go func(){
		for c := range ticker.C{
			t.Log(c)
		}
	}()
}

func Test_map(t *testing.T){
	data := make(map[string]int)
	data["s"] = 1
	data["b"] = 2
	for a,b := range data{
		t.Log(a, b)
	}
}
