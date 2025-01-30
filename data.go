package gilder_data

import (
	_ "embed"
	"encoding/json"
	"sort"

	"github.com/Hubmakerlabs/gilder-data/game"
	"github.com/Hubmakerlabs/gilder-data/locale"
)

//go:embed data/countriestimezoneslanguages.json
var nations []byte

//go:embed data/timezones.json
var timezones []byte

//go:embed data/games.json
var games []byte

var Nations locale.Nations
var TimeZones locale.TimeZoneList
var Games game.GamesList

func init() {
	var err error
	if err = json.Unmarshal(nations, &Nations); chk.E(err) {
		panic(err)
	}
	sort.Sort(Nations)
	if err = json.Unmarshal(timezones, &TimeZones); chk.E(err) {
		panic(err)
	}
	sort.Sort(TimeZones)
	if err = json.Unmarshal(games, &Games); chk.E(err) {
		panic(err)
	}
	sort.Sort(Games)
}
