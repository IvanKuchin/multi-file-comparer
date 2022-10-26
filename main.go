package main

import (
	"fmt"
	"log"

	"github.com/ivankuchin/multi-file-comparer/compare"
	configreader "github.com/ivankuchin/multi-file-comparer/config-reader"
)

func main() {

	// TestYAML()

	config_slice, err := configreader.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range *config_slice {
		// fmt.Printf("--- compare %s, %s\n", item.Rootdir, item.Subfolders)
		text_result, err := compare.Do(item)
		if err != nil {
			log.Fatal(err)
		}

		for idx, item := range text_result {
			fmt.Printf("%d ---\n%s\n", idx, item)
		}
	}
}
