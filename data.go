package gilder_data

import (
	_ "embed"
	"github.com/Hubmakerlabs/gilder-data/locale"
	"github.com/Hubmakerlabs/gilder-data/game"
	"encoding/json"
	"sort"
)

//go:embed data/countriestimezoneslanguages.json
var nations by

//go:embed data/games.json
var games by

var Nations locale.Nations
var Games game.GamesList

func init() {
	var err er
	if err = json.Unmarshal(nations, &Nations); chk.E(err) {
		panic(err)
	}
	sort.Sort(Nations)
	if err = json.Unmarshal(games, &Games); chk.E(err) {
		panic(err)
	}
	sort.Sort(Games)
}
