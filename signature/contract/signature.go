package contract

import "time"

type Signature interface {
	Date() time.Time
	Size() string
	// Name name of file
	Name() string
	SignatureByte() []byte
	Equals(s Signature) bool
}
