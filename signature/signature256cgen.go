package signature

import "encoding/binary"
import "bytes"

func (in *SignatureSha256) Unpack(data []byte) error {
	r := bytes.NewReader(data)

	// size
	var sizeRaw uint64
	binary.Read(r, binary.BigEndian, &sizeRaw)
	in.size = uint64(sizeRaw)

	// name
	var nameLenRaw uint16
	binary.Read(r, binary.BigEndian, &nameLenRaw)
	nameRaw := make([]byte, nameLenRaw)
	binary.Read(r, binary.BigEndian, nameRaw)
	in.name = string(nameRaw)

	// signature
	var signatureLenRaw uint16
	binary.Read(r, binary.BigEndian, &signatureLenRaw)
	signatureRaw := make([]byte, signatureLenRaw)
	binary.Read(r, binary.BigEndian, signatureRaw)
	in.signature = signatureRaw
	return nil
}
