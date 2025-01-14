package game

import (
	_ "embed"
	"github.com/dimuska139/rawg-sdk-go/v3"
	"net/http"
	"realy.lol/context"
	"time"
	"strings"
	"encoding/json"
	"os"
)

var mp = []st{
	"2-players",
	"4-player-local",
	"asynchronous-multiplayer",
	"battle-royale",
	"co-op",
	"competitive",
	"coop",
	"cooperative",
	"e-sports",
	"esports",
	"family-sharing",
	"four-way",
	"group-play",
	"hot-seat",
	"lan-party",
	"lan-party-2",
	"leaderboard",
	"local-co-op",
	"mmo",
	"multiplayer",
	"online-co-op",
	"online-multiplayer",
	"online-pvp",
	"party-game",
	"pvp",
	"real-time-strategy",
	"rts",
	"team-based",
	"two-player-alternating",
	"valve-anti-cheat-enabled",
}

const (
	GamesJsonFile = "games.json"
)

func GetMpGames(jsonFile st) (o st) {
	if jsonFile == "" {
		jsonFile = GamesJsonFile
	}
	var err er
	var fi os.FileInfo
	if fi, err = os.Stat(jsonFile); err == nil {
		modTime := fi.ModTime()
		update := time.Now().Add(time.Hour * 24 * 90)
		if !modTime.After(update) {
			// no need to update more than once a season
			var b by
			if b, err = os.ReadFile(jsonFile); !chk.E(err) {
				o = st(b)
				return
			}
		}
	}
	key := os.Getenv("RAWG_KEY")
	config := rawg.Config{
		ApiKey:   key, // Your personal API key (see https://rawg.io/apidocs)
		Language: "en",
		Rps:      1,
	}
	client := rawg.NewClient(http.DefaultClient, &config)
	var i, count no
	var gameList []*FullGameDetail
	o = "["
	var first bool
	var data []*rawg.Game
	var total no
done:
	for {
		log.I.Ln("query page", i)
		i++
		filter := rawg.NewGamesFilter().SetPageSize(100).SetPage(i)
		ctx, cancel := context.Timeout(context.Bg(), time.Second*20)
		if data, total, err = client.GetGames(ctx, filter); chk.E(err) {
			cancel()
			break done
		}
		time.Sleep(time.Second)
		_ = total
	cont:
		for _, gd := range data {
			// nothing older than 2018
			if gd.Released.Format("2006") < "2018" {
				// log.I.F("%s too old %v", game.Name, game.Released)
				continue cont
			}
			// we won't note games that nobody is playing
			if gd.AddedByStatus.Playing < 10 {
				continue cont
			}
			for _, tag := range gd.Tags {
				if tag.Language == "eng" {
					for _, t := range mp {
						// we are only interested in the mp tagged games
						if strings.Contains(tag.Slug, t) {
							var b by
							// we are only interested in computer games, not consoles
							var relevantPlatform bo
						platforms:
							for _, v := range gd.ParentPlatforms {
								switch v.Platform.Name {
								case "PC", "Apple Macintosh", "Linux":
									relevantPlatform = true
									break platforms
								}
							}
							if !relevantPlatform {
								// log.I.F("%s not relevant platforms %v", game.ParentPlatforms)
								continue cont
							}
							count++
							g := &FullGameDetail{
								Id:              gd.ID,
								Name:            gd.Name,
								BackgroundImage: gd.ImageBackground,
								Metacritic:      gd.Metacritic,
								Released:        gd.Released.Format("2006-01-02"),
								Players:         gd.AddedByStatus.Playing,
							}
							for _, p := range gd.ParentPlatforms {
								// switch p.Platform.Name {
								// case "PC", "Apple Macintosh", "Linux":
								g.Platforms = append(g.Platforms, p.Platform.Name)
								// }
							}
							for _, v := range gd.Genres {
								g.Genres = append(g.Genres, v.Name)
							}
							gameList = append(gameList, g)
							b, err = json.Marshal(g)
							if !first {
								first = true
							} else {
								o += ","
							}
							o += st(b)
							continue cont
						}
					}
				}
			}
		}
	}
	o += "]"
	// cache the current version so we can avoid making it again any time too soon
	chk.E(os.WriteFile(jsonFile, by(o), 0660))
	return
}
