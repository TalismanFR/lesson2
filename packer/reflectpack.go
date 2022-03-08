package packer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
)

func Unpack(u interface{}, data []byte) (err error) {
	r := bytes.NewReader(data)
	v := reflect.ValueOf(u)
	//v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr()))
	orderByte := binary.BigEndian
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.Uint32, reflect.Uint64:
			var value uint64
			binary.Read(r, orderByte, &value)
			//value = 12
			//v.Field(i).SetUint(value)
		case reflect.String:
			var lenRaw uint16
			binary.Read(r, orderByte, &lenRaw)
			dataRaw := make([]byte, lenRaw)
			binary.Read(r, orderByte, &dataRaw)
			v.Field(i).SetString(string(dataRaw))
		case reflect.Slice:
			var lenRaw uint16
			binary.Read(r, orderByte, &lenRaw)
			s := make([]byte, lenRaw)
			r.Read(s)
			v.Field(i).SetBytes(s)
		default:
			fmt.Printf("undefined type %v", v.Field(i).Type().Kind().String())
		}
	}
	return
}

func Pack(sig interface{}) (buf *bytes.Buffer, err error) {
	v := reflect.ValueOf(sig)

	if v.Kind() != reflect.Struct {
		err = errors.New("Is not struct type")
		return
	}
	orderByte := binary.BigEndian

	buff := &bytes.Buffer{}
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.Uint32, reflect.Uint64:
			s := v.Field(i).Uint()
			_ = s
			err = binary.Write(buff, orderByte, s)
		case reflect.String:
			s := v.Field(i).String()
			b := []byte(s)
			err = binary.Write(buff, orderByte, uint16(len(b)))
			if err != nil {
				return
			}
			//buf.Write(b)
			err = binary.Write(buff, orderByte, b)
		case reflect.Slice:
			//buf.Write(v.Field(i).Bytes())
			err = binary.Write(buff, orderByte, uint16(len(v.Field(i).Bytes())))
			err = binary.Write(buff, orderByte, v.Field(i).Bytes())
		default:
			fmt.Printf("undefined type field %v\n", v.Field(i).Type().Kind())
		}

	}

	buf = buff
	return
}
