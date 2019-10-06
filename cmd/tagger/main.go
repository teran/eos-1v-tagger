package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	tagger "github.com/teran/eos-1v-tagger"
)

func main() {
	t, err := tagger.NewCSVParser(os.Args[1])
	if err != nil {
		log.Fatalf("error initializing CSV parser: %s", err)
	}

	film, err := t.Parse()
	if err != nil {
		log.Fatalf("error parsing CSV: %s", err)
	}

	j, err := json.Marshal(film)
	if err != nil {
		log.Fatalf("error marshaling JSON: %s", err)
	}

	fmt.Println(string(j))
}
