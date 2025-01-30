package game

// FullGameDetail is a more detailed metadata that is needed for TeamUp.
type FullGameDetail struct {
	Name            string   `json:"name"`
	Id              int      `json:"id"`
	BackgroundImage string   `json:"background_image"`
	Metacritic      int      `json:"metacritic"`
	Released        string   `json:"released"`
	Platforms       []string `json:"platforms"`
	Genres          []string `json:"genres"`
	Players         int      `json:"players"`
}

type GamesList []FullGameDetail

func (l GamesList) Len() int           { return len(l) }
func (l GamesList) Less(i, j int) bool { return l[i].Name < l[j].Name }
func (l GamesList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func (l GamesList) Get(name string) (d *FullGameDetail) {
	for _, v := range l {
		if v.Name == name {
			return &v
		}
	}
	return
}
