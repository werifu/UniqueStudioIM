package gotest

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func BenchmarkSplit(b *testing.B) {
	str := "access_token=e72e16c7e42f292c6912e7710c838347ae178b4a&token_type=bearer"
	for i:=0;i<b.N;i++ {
		result := strings.Split(strings.Split(str, "&")[0], "=")[1]
		if result == "" {
			fmt.Println("result empty")
		}
	}
}

func BenchmarkRegexp(b *testing.B) {
	str := "access_token=e72e16c7e42f292c6912e7710c838347ae178b4a&token_type=bearer"
	for i:=0;i<b.N;i++ {
		regex := regexp.MustCompile(`access_token=([^\s&]+)`)
		result := regex.FindStringSubmatch(str)[1]

		if result == "" {
			fmt.Printf("%q", result)
		}
	}
}

type AccessTokenResponse struct{
	AccessToken		string	`json:"access_token"`
	Scope  			string	`json:"scope"`
	TokenType 		string	`json:"token_type"`
}
func BenchmarkJsonParse(b *testing.B) {
	res := AccessTokenResponse{}
	str := []byte(`{"access_token":"e72e16c7e42f292c6912e7710c838347ae178b4a", "scope":"repo,gist", "token_type":"bearer"}`)
	for i:=0;i<b.N;i++ {
		err := json.Unmarshal(str, &res)
		if err != nil {
			continue
		}
	}
}
