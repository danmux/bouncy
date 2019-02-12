package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/cld2"
	"github.com/slyrz/warc"

	// to trigger the init func
	_ "github.com/danmux/bouncy/foundationdb"
)

func main() {

	in := flag.String("in", "", "the input warc file")
	flag.Parse()

	f, err := os.Open(*in)
	if err != nil {
		log.Fatal(err)
	}
	reader, err := warc.NewReader(f)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	c := 0

	// open a new index
	mapping := bleve.NewIndexMapping()

	// you will need to `rm -rf example.bleve`
	index, err := bleve.NewUsing("example.bleve", mapping, "upside_down", "foundationdb", nil)
	if err != nil {
		log.Fatalf("new failed <%v>", err)
	}

	start := time.Now()
	for c < 200 {
		// fmt.Println()
		record, err := reader.ReadRecord()
		if err != nil {
			break
		}
		if c%100 == 0 {
			fmt.Println(c, "in", time.Since(start))
		}

		url, ok := record.Header["warc-target-uri"]
		if !ok {
			fmt.Println("no url")
			continue
		}
		// fmt.Println(url)
		b, err := ioutil.ReadAll(record.Content)
		if err != nil {
			log.Println(err)
			continue
		}

		l := cld2.Detect(string(b))
		if l == "en" || l == "fr" || l == "de" {
			c++
			err = index.Index(url, b)
			if err != nil {
				log.Println(err)
			}
		}

		// for key, value := range record.Header {
		// 	fmt.Printf("%v = %v\n", key, value)
		// }
	}
	fmt.Println(c, "in", time.Since(start))

	query := bleve.NewMatchQuery("widget")
	search := bleve.NewSearchRequest(query)
	searchResult, err := index.Search(search)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(searchResult.String())
}

// func printMapSorted(m map[string]int) {
// 	keys := make([]string, 0, len(m))
// 	for k := range m {
// 		keys = append(keys, k)
// 	}
// 	sort.Strings(keys)
// 	for _, k := range keys {
// 		fmt.Println(k, m[k])
// 	}
// }
