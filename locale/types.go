package locale

import (
	"time"
)

// countries

type Tz struct {
	ZoneName      string `json:"zoneName"`
	GmtOffset     int    `json:"gmtOffset"`
	GmtOffsetName string `json:"gmtOffsetName"`
	Abbreviation  string `json:"abbreviation"`
	TzName        string `json:"tzName"`
}

type Tzs []Tz

type Translations struct {
	Ko   string `json:"ko"`
	PtBR string `json:"pt-BR"`
	Pt   string `json:"pt"`
	Nl   string `json:"nl"`
	Hr   string `json:"hr"`
	Fa   string `json:"fa"`
	De   string `json:"de"`
	Es   string `json:"es"`
	Fr   string `json:"fr"`
	Ja   string `json:"ja"`
	It   string `json:"it"`
	ZhCN string `json:"zh-CN"`
	Tr   string `json:"tr"`
	Ru   string `json:"ru"`
	Uk   string `json:"uk"`
	Pl   string `json:"pl"`
}

type City struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Cities []City

type State struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	StateCode string `json:"state_code"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Type      string `json:"type"`
	Cities    `json:"cities,omitempty"`
}

type States []State

type Country struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Iso3           string `json:"iso3"`
	Iso2           string `json:"iso2"`
	NumericCode    string `json:"numeric_code"`
	PhoneCode      string `json:"phonecode"`
	Capital        string `json:"capital"`
	Currency       string `json:"currency"`
	CurrencyName   string `json:"currency_name"`
	CurrencySymbol string `json:"currency_symbol"`
	Tld            string `json:"tld"`
	Native         string `json:"native"`
	Region         string `json:"region"`
	RegionId       int    `json:"region_id"`
	Subregion      string `json:"subregion"`
	SubregionId    int    `json:"subregion_id"`
	Nationality    string `json:"nationality"`
	Tzs            `json:"timezones"`
	Translations   `json:"translations"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	Emoji          string `json:"emoji"`
	EmojiU         string `json:"emojiU"`
	States         `json:"states"`
}

type Countries []Country

// timezone

type timeZone struct {
	TimeZone         string `json:"timeZone"`
	CurrentLocalTime string `json:"currentLocalTime"`
	CurrentUtcOffset struct {
		Seconds      int   `json:"seconds"`
		Milliseconds int   `json:"milliseconds"`
		Ticks        int64 `json:"ticks"`
		Nanoseconds  int64 `json:"nanoseconds"`
	} `json:"currentUtcOffset"`
	StandardUtcOffset struct {
		Seconds      int   `json:"seconds"`
		Milliseconds int   `json:"milliseconds"`
		Ticks        int64 `json:"ticks"`
		Nanoseconds  int64 `json:"nanoseconds"`
	} `json:"standardUtcOffset"`
	HasDayLightSaving      bool `json:"hasDayLightSaving"`
	IsDayLightSavingActive bool `json:"isDayLightSavingActive"`
	DstInterval            struct {
		DstName        string `json:"dstName"`
		DstOffsetToUtc struct {
			Seconds      int   `json:"seconds"`
			Milliseconds int   `json:"milliseconds"`
			Ticks        int64 `json:"ticks"`
			Nanoseconds  int64 `json:"nanoseconds"`
		} `json:"dstOffsetToUtc"`
		DstOffsetToStandardTime struct {
			Seconds      int   `json:"seconds"`
			Milliseconds int   `json:"milliseconds"`
			Ticks        int64 `json:"ticks"`
			Nanoseconds  int64 `json:"nanoseconds"`
		} `json:"dstOffsetToStandardTime"`
		DstStart    time.Time `json:"dstStart"`
		DstEnd      time.Time `json:"dstEnd"`
		DstDuration struct {
			Days                 int     `json:"days"`
			NanosecondOfDay      int     `json:"nanosecondOfDay"`
			Hours                int     `json:"hours"`
			Minutes              int     `json:"minutes"`
			Seconds              int     `json:"seconds"`
			Milliseconds         int     `json:"milliseconds"`
			SubsecondTicks       int     `json:"subsecondTicks"`
			SubsecondNanoseconds int     `json:"subsecondNanoseconds"`
			BclCompatibleTicks   int64   `json:"bclCompatibleTicks"`
			TotalDays            float64 `json:"totalDays"`
			TotalHours           float64 `json:"totalHours"`
			TotalMinutes         float64 `json:"totalMinutes"`
			TotalSeconds         int     `json:"totalSeconds"`
			TotalMilliseconds    int64   `json:"totalMilliseconds"`
			TotalTicks           int64   `json:"totalTicks"`
			TotalNanoseconds     int64   `json:"totalNanoseconds"`
		} `json:"dstDuration"`
	} `json:"dstInterval"`
}
