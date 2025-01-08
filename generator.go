package main

import (
	"compress/gzip"
	"github.com/Hubmakerlabs/gilder-data/locale"
	"os"
)

const (
	localeJson   = "countriestimezoneslanguages.json"
	localeJsonGz = localeJson + ".gz"
)

func main() {
	data := locale.GetData("countriestimezoneslanguages.json")
	fh, err := os.OpenFile(localeJsonGz, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	var w *gzip.Writer
	w, err = gzip.NewWriterLevel(fh, gzip.BestCompression)
	_, err = w.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	if err = w.Flush(); err != nil {
		panic(err)
	}
	if err = w.Close(); err != nil {
		panic(err)
	}
}
