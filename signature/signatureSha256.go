package signature

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"lesson2/signature/contract"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type SignatureSha256 struct {
	date      time.Time
	size      string
	name      string
	signature []byte
}

const separator = "====sign===="

func NewSignatureSha256FromFileSource(file *os.File, hashString string) (sig SignatureSha256, err error) {
	stat, _ := file.Stat()
	sig.size = strconv.FormatInt(stat.Size(), 10)
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

func (s SignatureSha256) encrypt(text string) []byte {
	sha := sha256.New()
	sha.Write([]byte(text))
	return sha.Sum(nil)
}

func NewSignatureSha256(date time.Time, size string, name string, signature []byte) *SignatureSha256 {
	return &SignatureSha256{date: date, size: size, name: name, signature: signature}
}

func (s SignatureSha256) Date() time.Time {
	return s.date
}

func (s SignatureSha256) Size() string {
	return s.size
}

func (s SignatureSha256) Name() string {
	return s.name
}

//yyyy-mm-dd hh:ii:ss
//2006-01-02 15-04-05
func (s SignatureSha256) headString() string {
	return strings.Join([]string{s.Date().Format("2006-01-02 15-04-05"), s.Size(), s.Name()}, ":")
}

func (s SignatureSha256) SignatureByte() []byte {
	result := bytes.NewBufferString(s.headString())
	result.WriteString(separator)
	result.Write(s.signature)
	return result.Bytes()
}

func (s SignatureSha256) Equals(ss contract.Signature) bool {
	return true
}
