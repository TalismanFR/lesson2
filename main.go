package main

import (
	"GO_MTS/lesson7/mycrypt"
	"GO_MTS/lesson7/packer"
	"GO_MTS/lesson7/signature"
	"flag"
	"fmt"
	"log"
	"reflect"
)

func main() {

	var fileSource, hashFile, outFile string

	flag.StringVar(&fileSource, "source-file", "", "File source")
	flag.StringVar(&hashFile, "hash-file", "", "File hash")
	flag.StringVar(&outFile, "out-file", "sign.txt", "FIle output")

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

		sign := encoder.Signature()
		PrintReflect(sign)
		buf, err := packer.Pack(sign)
		if err != nil {
			panic(err)
		}
		fmt.Println("bytes", buf.Bytes())

		si := signature.SignatureSha256{}
		packer.Unpack(si, buf.Bytes())
		fmt.Println("unpack struct ", si)

		//var s = struct{ foo int }{654}
		//rf := reflect.ValueOf(&s).Elem().Field(0)
		//rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
		//rf.SetInt(6)
		//fmt.Println(s)

		err = encoder.SaveToFile(outFile)
		if err != nil {
			panic(err)
		}
	case "dec":
		decoder, err := mycrypt.NewDecryptor(hashFile, outFile, fileSource)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(decoder)

		err = decoder.ValidateSign()
		if err != nil {
			panic(err)
		}

	default:
		log.Fatal("Use enc or dec param")
	}
}

func PrintReflect(u interface{}) error {
	val := reflect.ValueOf(u)
	fmt.Printf("%T имеет %d свойств и %d методов\n", u, val.NumField(), val.NumMethod())

	for i := 0; i < val.NumField(); i++ {
		fmt.Printf("\tname=%v, type=%v, valu=%v, tag='%v'\n",
			val.Type().Field(i).Name,
			val.Field(i).Type().Kind(),
			val.Field(i),
			val.Type().Field(i).Tag)
	}
	return nil
}
