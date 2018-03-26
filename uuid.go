package goutils


import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"
	"io/ioutil"
)


type UUID [16]byte
var  Nil   UUID
var  rander = rand.Reader
var filename = "/var/.uuid"

func newuuid() UUID {
	return must(newRandom())
}


//GetUUID 调用判断如果没有 uuid 文件,创建一个, 返回 uuid
func GetUUID() string {
	_, err := os.Stat(filename)
	//没有找到文件
	if err != nil {
		uuid := newuuid().String()
		f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
		defer f.Close()
		f.WriteString(uuid)
		return uuid
	}

	//找到文件
	f,_ := os.Open(filename)
	fuuid,_ := ioutil.ReadAll(f)
	return string(fuuid)

}

func newRandom() (UUID, error) {
	var uuid UUID
	_, err := io.ReadFull(rander, uuid[:])
	if err != nil {
		return Nil, err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return uuid, nil
}

func (uuid UUID) String() string {
	var buf [32]byte
	encodeHex(buf[:], uuid)
	return string(buf[:])
}

func encodeHex(dst []byte, uuid UUID) {
	hex.Encode(dst[:], uuid[:])
	//dst[8] = '-'
	//hex.Encode(dst[8:12], uuid[4:6])
	//dst[13] = '-'
	//hex.Encode(dst[14:18], uuid[6:8])
	//dst[18] = '-'
	//hex.Encode(dst[19:23], uuid[8:10])
	//dst[23] = '-'
	//hex.Encode(dst[24:], uuid[10:])
}


// Must returns uuid if err is nil and panics otherwise.
func must(uuid UUID, err error) UUID {
	if err != nil {

	}
	return uuid
}

