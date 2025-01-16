package game

// FullGameDetail is a more detailed metadata that is needed for TeamUp.
type FullGameDetail struct {
	Name            st   `json:"name"`
	Id              no   `json:"id"`
	BackgroundImage st   `json:"background_image"`
	Metacritic      no   `json:"metacritic"`
	Released        st   `json:"released"`
	Platforms       []st `json:"platforms"`
	Genres          []st `json:"genres"`
	Players         no   `json:"players"`
}

type GamesList []FullGameDetail

func (l GamesList) Len() int           { return len(l) }
func (l GamesList) Less(i, j int) bool { return l[i].Name < l[j].Name }
func (l GamesList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func (l GamesList) Get(name st) (d *FullGameDetail) {
	for _, v := range l {
		if v.Name == name {
			return &v
		}
	}
	return
}
