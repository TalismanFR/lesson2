package main

import (
	"flag"
	"fmt"
	"lesson2/mycrypt"
	"log"
)

func main() {
	var fileSource, hashFile, outFile string

	flag.StringVar(&fileSource, "source-file", "", "File source")
	flag.StringVar(&hashFile, "hash-file", "", "File hash")
	flag.StringVar(&outFile, "out-file", "sign.txt", "File output")

	flag.Parse()
	action := flag.Args()[0]

	switch action {
	case "enc":
		encoder, err := mycrypt.NewEncoder(fileSource, hashFile)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(encoder)

		err = encoder.EncryptSha256()
		if err != nil {
			panic(err)
		}

		err = encoder.SaveToFile(outFile)
		if err != nil {
			panic(err)
		}
	case "dec":

	default:
		log.Fatal("Use enc or dec param")
	}
}
