package main

import (
	"compress/gzip"
	"github.com/Hubmakerlabs/gilder-data/locale"
	"os"
	"github.com/Hubmakerlabs/gilder-data/game"
)

const (
	localeJson      = "countriestimezoneslanguages.json"
	localeJsonGz    = localeJson + ".gz"
	gamesJsonFile   = "games.json"
	gamesJsonFileGz = gamesJsonFile + ".gz"
)

func main() {
	nations := locale.GetNations(localeJson)
	fh, err := os.OpenFile(localeJsonGz, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	var w *gzip.Writer
	w, err = gzip.NewWriterLevel(fh, gzip.BestCompression)
	_, err = w.Write(by(nations))
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
	_, err = w.Write(by(games))
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
