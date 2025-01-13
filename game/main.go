package game

type T struct {
	Name            string   `json:"name"`
	Id              int      `json:"id"`
	BackgroundImage string   `json:"background_image"`
	Genres          []string `json:"genres"`
	Metacritic      int      `json:"metacritic"`
	Platforms       []string `json:"platforms"`
	ReleaseDate     string   `json:"release_date"`
}

type List []T

func (l List) Len() int           { return len(l) }
func (l List) Less(i, j int) bool { return l[i].Name < l[j].Name }
func (l List) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
