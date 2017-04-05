package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

var strListenAddr string
var strRootUrl string
var strRootDir string

func main() {
	flag.Parse()

	fmt.Printf("strListenAddr=%s", strListenAddr)
	fmt.Printf("strRootUrl=%s", strRootUrl)
	fmt.Printf("strRootDir=%s", strRootDir)

	var dir = http.Dir(strRootDir)
	var server = http.FileServer(dir)
	var handle = http.StripPrefix(strRootUrl, server)

	var logHandle = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s%s\r\n", r.RemoteAddr, r.Method, r.Host, r.URL)

		handle.ServeHTTP(w, r)
	})

	http.Handle(strRootUrl, logHandle)
	http.ListenAndServe(strListenAddr, nil)
}

func init() {
	flag.StringVar(&strListenAddr, "addr", ":10001", "监听地址")
	flag.StringVar(&strRootUrl, "url", "/", "根url /x/ end with /")
	flag.StringVar(&strRootDir, "dir", "./", "根目录")
}

// func a(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(1)
// }

// func t() {
// 	for {
// 		time.Sleep(time.Duration(5 * time.Second))
// 		strLines := utils.ReadLines()
// 		for k, v := range strLines {
// 			fmt.Println(k, v)
// 		}
// 	}
// }
