package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

func httpsrv() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `hello world!`)
	})
	log.Println("Listening on 18080...")
	log.Fatal(http.ListenAndServe(":18080", nil))
}

func main() {

	go httpsrv()

	getenvironment := func(data []string, getkeyval func(item string) (key, val string)) map[string]string {
		items := make(map[string]string)
		for _, item := range data {
			key, val := getkeyval(item)
			items[key] = val

		}
		return items

	}
	cnt := 0
	for {
		environment := getenvironment(os.Environ(), func(item string) (key, val string) {
			splits := strings.Split(item, "=")
			key = splits[0]
			val = splits[1]
			return

		})
		keys := []string{}
		for k := range environment {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, key := range keys {
			fmt.Println(key, "=", environment[key])
		}
		cnt += 1
		fmt.Println("hello#", cnt)
		/*
			for k, v := range environment {
				fmt.Println(k, "=", v)
			}
		*/
		time.Sleep(time.Second * 3)
	}
}
