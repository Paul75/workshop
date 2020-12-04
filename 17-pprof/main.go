package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/pkg/profile"
)

func main() {
	// add this source
	p := profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	defer p.Stop()
	// 1 - build the source with go build
	// run the binary in your terminal
	// then run the following command:
	// go tool pprof -http=:8080 cpu.pprof

	// Simple concat of string.
	rand.Seed(time.Now().UnixNano())
	res := ""
	for i := 0; i < 10e4; i++ {
		res += randomString()
	}
	fmt.Println(res)

	// Usage of a buffer.
	// var buffer bytes.Buffer
	// for i := 0; i < 10e4; i++ {
	// 	buffer.WriteString(randomString())
	// }
	// fmt.Println(buffer.String())
}

const all = "abcdefghijklmnopqrstuvwxyz"

func randomString() string {
	length := rand.Intn(15) + 2
	buf := make([]byte, length)

	for i := 0; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return strings.Title(string(buf))
}
