package main

import (
	"compress/gzip"
	"os"

	"github.com/Hubmakerlabs/gilder-data/game"
	"github.com/Hubmakerlabs/gilder-data/locale"
)

const (
	localeJson      = "countriestimezoneslanguages.json"
	localeJsonGz    = localeJson + ".gz"
	gamesJsonFile   = "games.json"
	gamesJsonFileGz = gamesJsonFile + ".gz"
	timezonesJson   = "timezones.json"
)

func main() {
	var err error
	_, _, err = locale.GetTimezones(timezonesJson)
	nations := locale.GetNations(localeJson)
	var fh *os.File
	fh, err = os.OpenFile(localeJsonGz, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	var w *gzip.Writer
	w, err = gzip.NewWriterLevel(fh, gzip.BestCompression)
	_, err = w.Write([]byte(nations))
	if err != nil {
		panic(err)
	}
	if err = w.Flush(); err != nil {
		panic(err)
	}
	if err = w.Close(); err != nil {
		panic(err)
	}
	games := game.GetMpGames(gamesJsonFile)
	fh, err = os.OpenFile(gamesJsonFileGz, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	w, err = gzip.NewWriterLevel(fh, gzip.BestCompression)
	_, err = w.Write([]byte(games))
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
