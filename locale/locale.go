package locale

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"sort"
	"time"
)

type TimeZone struct {
	Name      string `json:"name"`
	Offset    int    `json:"offset"`
	DstName   string `json:"dst_name,omitempty"`
	DstOffset int    `json:"dst_offset,omitempty"`
	DstStart  int64  `json:"dst_start,omitempty"`
	DstEnd    int64  `json:"dst_end,omitempty"`
}

type TimeZones map[string]TimeZone

type TimeZoneList []TimeZone

func (t TimeZoneList) Len() int           { return len(t) }
func (t TimeZoneList) Less(i, j int) bool { return t[i].Name < t[j].Name }
func (t TimeZoneList) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t TimeZones) GetSortedNames() (s []string) {
	for i := range t {
		s = append(s, i)
	}
	sort.Strings(s)
	return
}

func (t TimeZoneList) Get(name string) (d *TimeZone) {
	for _, v := range t {
		if v.Name == name {
			return &v
		}
	}
	return
}

func (t TimeZoneList) GetInForceOffset(name string) (off time.Duration, err error) {
	tz := t.Get(name)
	if tz == nil {
		err = errorf.E("unknown timezone %s", name)
		return
	}
	off = time.Duration(tz.Offset) * time.Second
	now := time.Now().Unix()
	if now > tz.DstStart && now < tz.DstEnd {
		off = time.Duration(tz.DstOffset) * time.Second
	}
	return
}

// Nation is the structure containing the needed information for a web app
// handling information about user locations in Shonen Network.
type Nation struct {
	Id             int      `json:"id"`
	Name           string   `json:"name"`
	Iso3           string   `json:"iso3"`
	Iso2           string   `json:"iso2"`
	NumericCode    string   `json:"numeric_code"`
	PhoneCode      string   `json:"phone_code"`
	Capital        string   `json:"capital"`
	Currency       string   `json:"currency"`
	CurrencyName   string   `json:"currency_name"`
	CurrencySymbol string   `json:"currency_symbol"`
	Tld            string   `json:"tld"`
	Native         string   `json:"native"`
	Emoji          string   `json:"emoji"`
	Languages      []string `json:"languages"`
	TimeZones      `json:"timezones"`
}

type Nations []Nation

func (n Nations) Len() int           { return len(n) }
func (n Nations) Less(i, j int) bool { return n[i].Name < n[j].Name }
func (n Nations) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Nations) Get(name string) (d *Nation) {
	for _, v := range n {
		if v.Name == name {
			return &v
		}
	}
	return
}

const (
	countries          = "https://github.com/dr5hn/countries-states-cities-database/raw/refs/heads/master/json/countries+states+cities.json"
	availableTimeZones = "https://www.timeapi.io/api/timezone/availabletimezones"
	getZoneInfo        = "https://www.timeapi.io/api/timezone/zone?timeZone="
	CountriesJsonFile  = "countriestimezoneslanguages.json"
	TimezonesJsonFile  = "timezones.json"
)

func GetTimezones(jsonFile string) (o string, tzs TimeZones, err error) {
	if jsonFile == "" {
		jsonFile = TimezonesJsonFile
	}
	log.I.F("creating timezone json %s", jsonFile)
	var fi os.FileInfo
	if fi, err = os.Stat(jsonFile); err == nil {
		modTime := fi.ModTime()
		update := time.Now().Add(time.Hour * 24 * 90)
		if !modTime.After(update) {
			// int need to update more than once a season
			var b []byte
			if b, err = os.ReadFile(TimezonesJsonFile); !chk.E(err) {
				o = string(b)
				var tzl TimeZoneList
				if err = json.Unmarshal(b, &tzl); chk.E(err) {
					return
				}
				tzs = make(TimeZones)
				for _, v := range tzl {
					tzs[v.Name] = v
				}
				return
			}
		}
	}
	var res *http.Response
	// first get the available timezones
	if res, err = http.Get(availableTimeZones); chk.E(err) {
		return
	}
	var b []byte
	if b, err = io.ReadAll(res.Body); chk.E(err) {
		return
	}
	tzs = make(TimeZones)
	var zoneNames []string
	if err = json.Unmarshal(b, &zoneNames); chk.E(err) {
		return
	}
	log.I.F("got available timezones, total %d", len(zoneNames))
	// next get the details for each named timezone
	for _, v := range zoneNames {
		if res, err = http.Get(getZoneInfo + v); chk.E(err) {
			return
		}
		if b, err = io.ReadAll(res.Body); chk.E(err) {
			return
		}
		var z timeZone
		if err = json.Unmarshal(b, &z); chk.E(err) {
			return
		}
		if z.TimeZone == "" {
			continue
		}
		log.I.F("got timezone %s", z.TimeZone)
		zi := TimeZone{
			Name:   z.TimeZone,
			Offset: z.StandardUtcOffset.Seconds,
		}
		if z.HasDayLightSaving {
			zi.DstName = z.DstInterval.DstName
			zi.DstOffset = z.DstInterval.DstOffsetToUtc.Seconds
			zi.DstStart = z.DstInterval.DstStart.Unix()
			zi.DstEnd = z.DstInterval.DstEnd.Unix()
		}
		tzs[z.TimeZone] = zi
	}
	var tzl TimeZoneList
	for _, v := range tzs {
		tzl = append(tzl, v)
	}
	sort.Sort(tzl)
	b = b[:0]
	if b, err = json.Marshal(&tzl); chk.E(err) {
		return
	}
	// cache the current version so we can avoid making it again any time too soon
	chk.E(os.WriteFile(jsonFile, b, 0660))
	log.I.F("finished creating timezones json %s", jsonFile)
	return
}

func GetNations(jsonFile string) (o string) {
	if jsonFile == "" {
		jsonFile = CountriesJsonFile
	}
	log.I.F("creating nations json %s", jsonFile)
	var err error
	var fi os.FileInfo
	if fi, err = os.Stat(jsonFile); err == nil {
		modTime := fi.ModTime()
		update := time.Now().Add(time.Hour * 24 * 90)
		if !modTime.After(update) {
			// int need to update more than once a season
			var b []byte
			if b, err = os.ReadFile(CountriesJsonFile); !chk.E(err) {
				o = string(b)
				return
			}
		}
	}
	var c Countries
	var res *http.Response
	var b []byte
	// next, get the country data
	if res, err = http.Get(countries); chk.E(err) {
		return
	}
	if b, err = io.ReadAll(res.Body); chk.E(err) {
		return
	}
	if err = json.Unmarshal(b, &c); chk.E(err) {
		return
	}
	// condense the data down to the essentials we require, and link the time zones
	// to the ones we gathered before (as many overlap over the same longitude).
	var cc []Nation
	var tzs TimeZones
	_, tzs, err = GetTimezones(TimezonesJsonFile)
	for _, country := range c {
		ccc := Nation{
			Id:             country.Id,
			Name:           country.Name,
			Iso3:           country.Iso3,
			Iso2:           country.Iso2,
			NumericCode:    country.NumericCode,
			PhoneCode:      country.PhoneCode,
			Capital:        country.Capital,
			Currency:       country.Currency,
			CurrencyName:   country.CurrencyName,
			CurrencySymbol: country.CurrencySymbol,
			Tld:            country.Tld,
			Native:         country.Native,
			Emoji:          country.Emoji,
			TimeZones:      make(TimeZones),
		}
		lang := GetLanguageByIso2(country.Iso2)
		if lang != nil {
			ccc.Languages = (*lang).Languages
		} else {
			// default to english if we don't know it from our language db
			ccc.Languages = append(ccc.Languages, "English")
		}
		for _, tz := range country.Tzs {
			var ok bool
			if ccc.TimeZones[tz.ZoneName], ok = tzs[tz.ZoneName]; ok {
				x := ccc.TimeZones[tz.ZoneName]
				x.Name = tz.Abbreviation
				ccc.TimeZones[tz.ZoneName] = x
			}
		}
		log.I.F("got nation %s", ccc.Name)
		cc = append(cc, ccc)
	}
	if b, err = json.Marshal(&cc); chk.E(err) {
		return
	}
	// cache the current version so we can avoid making it again any time too soon
	chk.E(os.WriteFile(jsonFile, b, 0660))
	log.I.F("finished creating nations json %s", jsonFile)
	return string(b)
}
