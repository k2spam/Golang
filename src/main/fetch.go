package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
	"os"
	"io"
	"strings"
)

func main () {
	//for _, url := range os.Args[1:] {
		url := "http://localhost/connections/api/lists"

		prefix := strings.HasPrefix( url,"http://" )
		if !prefix {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf( "fetch: %v\n", err )
			os.Exit(1)
		}

		b, err := io.Copy( os.Stdout, resp.Body )
		resp.Body.Close()

		if err != nil {
			fmt.Printf("fetch: read %s: %v", url, err )
			os.Exit(1)
		}

		fmt.Printf("%v\n", resp.Status)
		fmt.Printf("%s\n", b )
	//}
}