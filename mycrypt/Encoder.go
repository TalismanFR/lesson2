package mycrypt

import (
	"io/ioutil"
	"lesson2/signature"
	"lesson2/signature/contract"
	"os"
)

type Encoder struct {
	hashSign   string
	fileSource string
	signature  contract.Signature
}

func NewEncoder(fileSource string, fileHashSign string) (enc *Encoder, err error) {
	hashString, err := ioutil.ReadFile(fileHashSign)
	if err != nil {
		return
	}

	enc = &Encoder{fileSource: fileSource, hashSign: string(hashString)}
	return
}

func (enc *Encoder) EncryptSha256() (err error) {
	file, err := os.Open(enc.fileSource)
	if err != nil {
		return err
	}
	defer file.Close()

	sign, err := signature.NewSignatureSha256FromFileSource(file, enc.hashSign)
	if err != nil {
		return
	}

	enc.signature = sign
	return
}

func (enc Encoder) SaveToFile(path string) (err error) {
	err = ioutil.WriteFile(path, enc.signature.SignatureByte(), 0644)
	return
}
