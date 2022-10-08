package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func main() {

	siteurl := "https://www.baidu.com"
	converter := md.NewConverter("", true, nil)

	resp, _ := http.Get(siteurl)
	respbytes, _ := ioutil.ReadAll(resp.Body)

	html := string(respbytes)

	fmt.Println(html)
	markdown, err := converter.ConvertURL(siteurl)
	// markdown, err := converter.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("md ->", markdown)

}
