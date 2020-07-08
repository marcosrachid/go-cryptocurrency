package utils

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func Compress(s []byte) []byte {
	zipbuf := bytes.Buffer{}
	zipped := gzip.NewWriter(&zipbuf)
	zipped.Write(s)
	zipped.Close()
	return zipbuf.Bytes()
}

func Decompress(s []byte) ([]byte, error) {
	rdr, _ := gzip.NewReader(bytes.NewReader(s))
	data, err := ioutil.ReadAll(rdr)
	if err != nil {
		return nil, err
	}
	rdr.Close()
	return data, nil
}
