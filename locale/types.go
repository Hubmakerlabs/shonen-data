package locale

import (
	"time"
)

// countries

type Tz struct {
	ZoneName      st `json:"zoneName"`
	GmtOffset     no `json:"gmtOffset"`
	GmtOffsetName st `json:"gmtOffsetName"`
	Abbreviation  st `json:"abbreviation"`
	TzName        st `json:"tzName"`
}

type Tzs []Tz

type Translations struct {
	Ko   st `json:"ko"`
	PtBR st `json:"pt-BR"`
	Pt   st `json:"pt"`
	Nl   st `json:"nl"`
	Hr   st `json:"hr"`
	Fa   st `json:"fa"`
	De   st `json:"de"`
	Es   st `json:"es"`
	Fr   st `json:"fr"`
	Ja   st `json:"ja"`
	It   st `json:"it"`
	ZhCN st `json:"zh-CN"`
	Tr   st `json:"tr"`
	Ru   st `json:"ru"`
	Uk   st `json:"uk"`
	Pl   st `json:"pl"`
}

type City struct {
	Id        no `json:"id"`
	Name      st `json:"name"`
	Latitude  st `json:"latitude"`
	Longitude st `json:"longitude"`
}

type Cities []City

type State struct {
	Id        no `json:"id"`
	Name      st `json:"name"`
	StateCode st `json:"state_code"`
	Latitude  st `json:"latitude"`
	Longitude st `json:"longitude"`
	Type      st `json:"type"`
	Cities    `json:"cities,omitempty"`
}

type States []State

type Country struct {
	Id             no `json:"id"`
	Name           st `json:"name"`
	Iso3           st `json:"iso3"`
	Iso2           st `json:"iso2"`
	NumericCode    st `json:"numeric_code"`
	PhoneCode      st `json:"phonecode"`
	Capital        st `json:"capital"`
	Currency       st `json:"currency"`
	CurrencyName   st `json:"currency_name"`
	CurrencySymbol st `json:"currency_symbol"`
	Tld            st `json:"tld"`
	Native         st `json:"native"`
	Region         st `json:"region"`
	RegionId       no `json:"region_id"`
	Subregion      st `json:"subregion"`
	SubregionId    no `json:"subregion_id"`
	Nationality    st `json:"nationality"`
	Tzs            `json:"timezones"`
	Translations   `json:"translations"`
	Latitude       st `json:"latitude"`
	Longitude      st `json:"longitude"`
	Emoji          st `json:"emoji"`
	EmojiU         st `json:"emojiU"`
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
