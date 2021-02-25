package gotest

import (
	"crypto/sha256"
	"io"
	"reflect"
	"testing"
	"thchat/pkg/util"
)


func TestSha256(t *testing.T) {
	str := "His money is twice tainted: 'taint yours and 'taint mine."
	h := sha256.New()
	io.WriteString(h, str)
	a := h.Sum(nil)

	h2 := sha256.New()
	io.WriteString(h2, str)
	b := h.Sum(nil)
	if !reflect.DeepEqual(a, b){
		t.Error()
	}
}


func BenchmarkStringSha256(b *testing.B) {
	str := "1145141919810"
	for i:=0;i<b.N;i++ {
		util.StringSha256(str)
	}
}
