package signature

import (
	"GO_MTS/lesson7/signature/contract"
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// File Signature
// cgen: binpack
type SignatureSha256 struct {
	date      time.Time `cgen:"-"`
	size      uint64
	name      string
	signature []byte
}

func (s *SignatureSha256) SetDate(date time.Time) {
	s.date = date
}

func (s *SignatureSha256) SetSize(size uint64) {
	s.size = size
}

func (s *SignatureSha256) SetName(name string) {
	s.name = name
}

func (s *SignatureSha256) SetSignature(signature []byte) {
	s.signature = signature
}

const separator = "====sign===="

func NewSignatureSha256FromFileSource(file *os.File, hashString string) (sig SignatureSha256, err error) {
	sig = SignatureSha256{}
	stat, _ := file.Stat()
	sig.size = uint64(stat.Size())
	sig.name = path.Base(file.Name())
	sig.date = stat.ModTime()

	var fileData = make([]byte, stat.Size())
	_, err = file.Read(fileData)
	if err != nil {
		return
	}
	data := string(fileData) + hashString
	fmt.Println("data ", data)
	sig.signature = sig.encrypt(data)
	fmt.Printf("sign line %x \n", sig.signature)

	return
}

func NewFileSourceFromSignatureSha256(file *os.File) (sig SignatureSha256, err error) {
	sig = SignatureSha256{}
	stat, _ := file.Stat()
	var fileData = make([]byte, stat.Size())
	_, err = file.Read(fileData)
	if err != nil {
		return
	}

	data := strings.Split(string(fileData), separator)
	println(string(fileData))
	sig.signature = []byte(data[1])
	dataWithDateSizeName := strings.Split(data[0], ":")
	sig.date, err = time.Parse("2006-01-02 15-04-05", dataWithDateSizeName[0])
	if err != nil {
		return
	}
	sizeString := dataWithDateSizeName[1]
	sig.size, _ = strconv.ParseUint(sizeString, 10, 64)
	sig.name = dataWithDateSizeName[2]

	return
}

func (s SignatureSha256) encrypt(text string) []byte {
	sha := sha256.New()
	sha.Write([]byte(text))
	return sha.Sum(nil)

}

func (s SignatureSha256) headString() string {
	return strings.Join([]string{s.Date().Format("2006-01-02 15-04-05"), s.Size(), s.Name()}, ":")
}

func NewSignatureSha256(date time.Time, size string, name string, signature []byte) *SignatureSha256 {
	s, _ := strconv.ParseUint(size, 20, 64)
	return &SignatureSha256{date: date, size: s, name: name, signature: signature}
}

func (s SignatureSha256) Date() time.Time {
	return s.date
}

func (s SignatureSha256) Size() string {
	return string(s.size)
}

func (s SignatureSha256) Name() string {
	return s.name
}

func (s SignatureSha256) Signature() []byte {
	return s.signature
}

func (s SignatureSha256) SignatureByte() []byte {
	result := bytes.NewBufferString(s.headString())
	result.WriteString(separator)
	result.Write(s.signature)
	return result.Bytes()
}

func (s SignatureSha256) Equals(ss contract.Signature) bool {
	if s.name != ss.Name() {
		println("The names are not same!\n")
		return false
	}
	if string(s.size) != ss.Size() {
		println("The sizes are not same!\n")
		return false
	}
	//if s.Date() != ss.Date(){
	//	println("The dates are not same!\n")
	//	return false
	//}
	if !s.Date().Equal(ss.Date()) {
		println("The date are not same!\n")
		return false
	}
	if len(s.Signature()) != len(ss.Signature()) {
		println("The signatures are not same!")
		return false
	}
	for i := range s.Signature() {
		if s.Signature()[i] != ss.Signature()[i] {
			println("The byte in signature is no same!")
			return false
		}
	}
	println("The signatures are same!")
	return true
}
