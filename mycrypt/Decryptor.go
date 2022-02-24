package mycrypt

import "io/ioutil"

type Decryptor struct {
	fileHash   string
	hashString string
	fileSource string
	fileSigned string
}

func NewDecryptor(fileHash string, fileSource string, fileSigned string) (dec *Decryptor, err error) {
	hashString, err := ioutil.ReadFile(fileHash)
	if err != nil {
		return
	}
	dec = &Decryptor{fileHash: fileHash, hashString: string(hashString), fileSource: fileSource, fileSigned: fileSigned}
	return
}

//add validate()

//gen signature
//build signature from signFile.xt

//signature.Equals(signature)
