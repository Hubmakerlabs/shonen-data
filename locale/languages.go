package locale

type Language struct {
	Country   string
	Iso2      string
	Languages []string
}

func GetLanguageByIso2(iso2 string) (v *Language) {
	for _, v = range LanguagesByNation {
		if v.Iso2 == iso2 {
			return
		}
	}
	return nil
}

var LanguagesByNation = []*Language{
	{
		Country: "Abkhazia",
		Iso2:    "AB",
		Languages: []string{
			"Abkhaz",
			"Russian",
		},
	},
	{
		Country: "Afghanistan",
		Iso2:    "AF",
		Languages: []string{
			"Persian",
			"Pashto",
		},
	},
	{
		Country: "Albania",
		Iso2:    "AL",
		Languages: []string{
			"Albanian",
		},
	},
	{
		Country: "Algeria",
		Iso2:    "DZ",
		Languages: []string{
			"Arabic",
			"Berber",
		},
	},
	{
		Country: "Andorra",
		Iso2:    "AD",
		Languages: []string{
			"Catalan",
		},
	},
	{
		Country: "Angola",
		Iso2:    "AO",
		Languages: []string{
			"Portuguese",
		},
	},
	{
		Country: "Antigua and Barbuda",
		Iso2:    "AG",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Argentina",
		Iso2:    "AR",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Armenia",
		Iso2:    "AM",
		Languages: []string{
			"Armenian",
		},
	},
	{
		Country: "Australia",
		Iso2:    "AU",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Austria",
		Iso2:    "AT",
		Languages: []string{
			"German",
		},
	},
	{
		Country: "Azerbaijan",
		Iso2:    "AZ",
		Languages: []string{
			"Azerbaijani",
		},
	},
	{
		Country: "Bahamas",
		Iso2:    "BS",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Bahrain",
		Iso2:    "BH",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Bangladesh",
		Iso2:    "BD",
		Languages: []string{
			"Bengali",
		},
	},
	{
		Country: "Barbados",
		Iso2:    "BB",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Belarus",
		Iso2:    "BY",
		Languages: []string{
			"Belarusian",
			"Russian",
		},
	},
	{
		Country: "Belgium",
		Iso2:    "BE",
		Languages: []string{
			"Dutch",
			"French",
			"German",
		},
	},
	{
		Country: "Belize",
		Iso2:    "BZ",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Benin",
		Iso2:    "BJ",
		Languages: []string{
			"French",
		},
	},
	{
		Country: "Bhutan",
		Iso2:    "BT",
		Languages: []string{
			"Dzongkha",
		},
	},
	{
		Country: "Bolivia",
		Iso2:    "BO",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Bosnia and Herzegovina",
		Iso2:    "BA",
		Languages: []string{
			"Bosnian",
			"Croatian",
			"Serbian",
		},
	},
	{
		Country: "Botswana",
		Iso2:    "BW",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Brazil",
		Iso2:    "BR",
		Languages: []string{
			"Portuguese",
		},
	},
	{
		Country: "Brunei",
		Iso2:    "BN",
		Languages: []string{
			"Malay",
		},
	},
	{
		Country: "Burkina Faso",
		Iso2:    "BF",
		Languages: []string{
			"English",
			"French",
		},
	},
	{
		Country: "Burundi",
		Iso2:    "BI",
		Languages: []string{
			"French",
			"Kirundi",
			"English",
		},
	},
	{
		Country: "Cambodia",
		Iso2:    "KH",
		Languages: []string{
			"Khmer",
		},
	},
	{
		Country: "Cameroon",
		Iso2:    "CM",
		Languages: []string{
			"English",
			"French",
		},
	},
	{
		Country: "Canada",
		Iso2:    "CA",
		Languages: []string{
			"English",
			"French",
		},
	},
	{
		Country: "Cape Verde",
		Iso2:    "CV",
		Languages: []string{
			"Portuguese",
		},
	},
	{
		Country: "Central African Republic",
		Iso2:    "CF",
		Languages: []string{
			"French",
			"Sango",
		},
	},
	{
		Country: "Chad",
		Iso2:    "TD",
		Languages: []string{
			"Arabic",
			"French",
		},
	},
	{
		Country: "Chile",
		Iso2:    "CL",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "China",
		Iso2:    "CN",
		Languages: []string{
			"Chinese",
		},
	},
	{
		Country: "Christmas Island",
		Iso2:    "CX",
		Languages: []string{
			"English",
			"Mandarin Chinese",
			"Malay",
		},
	},
	{
		Country: "Cocos Islands",
		Iso2:    "CC",
		Languages: []string{
			"English",
			"Cocos Malay",
		},
	},
	{
		Country: "Columbia",
		Iso2:    "CO",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Comoros",
		Iso2:    "KM",
		Languages: []string{
			"Arabic",
			"French",
			"Comorian",
		},
	},
	{
		Country: "Democratic Republic of the Congo",
		Iso2:    "CD",
		Languages: []string{
			"French",
		},
	},
	{
		Country: "Congo",
		Iso2:    "CG",
		Languages: []string{
			"French",
		},
	},
	{
		Country: "Cook Islands",
		Iso2:    "CK",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Costa Rica",
		Iso2:    "CR",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Croatia",
		Iso2:    "HR",
		Languages: []string{
			"Croatian",
		},
	},
	{
		Country: "Cuba",
		Iso2:    "CU",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Cyprus",
		Iso2:    "CY",
		Languages: []string{
			"Greek",
			"Turkish",
		},
	},
	{
		Country: "Czech Republic",
		Iso2:    "CZ",
		Languages: []string{
			"Czech",
			"Slovak",
		},
	},
	{
		Country: "Denmark",
		Iso2:    "DK",
		Languages: []string{
			"Danish",
		},
	},
	{
		Country: "Djibouti",
		Iso2:    "DJ",
		Languages: []string{
			"Arabic",
			"French",
		},
	},
	{
		Country: "Dominica",
		Iso2:    "DM",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Dominican Republic",
		Iso2:    "DO",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Timor-Leste",
		Iso2:    "TL",
		Languages: []string{
			"Portuguese",
		},
	},
	{
		Country: "Ecuador",
		Iso2:    "EC",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Egypt",
		Iso2:    "EG",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "El Salvador",
		Iso2:    "SV",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Equatorial Guinea",
		Iso2:    "GQ",
		Languages: []string{
			"French",
			"Portuguese",
			"Spanish",
		},
	},
	{
		Country: "Eritrea",
		Iso2:    "ER",
		Languages: []string{
			"Tigrinya",
		},
	},
	{
		Country: "Estonia",
		Iso2:    "EE",
		Languages: []string{
			"Estonian",
		},
	},
	{
		Country: "Eswatini",
		Iso2:    "SZ",
		Languages: []string{
			"English",
			"Swazi",
		},
	},
	{
		Country: "Ethiopia",
		Iso2:    "ET",
		Languages: []string{
			"Afar",
			"Amharic",
			"Oromo",
			"Somali",
			"Tigrinya",
		},
	},
	{
		Country: "Fiji",
		Iso2:    "FJ",
		Languages: []string{
			"English",
			"Fijian",
			"Hindi",
		},
	},
	{
		Country: "Finland",
		Iso2:    "FI",
		Languages: []string{
			"Finnish",
		},
	},
	{
		Country: "France",
		Iso2:    "FR",
		Languages: []string{
			"French",
		},
	},
	{
		Country: "Gabon",
		Iso2:    "GA",
		Languages: []string{
			"French",
		},
	},
	{
		Country: "Gambia",
		Iso2:    "GM",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Georgia",
		Iso2:    "GE",
		Languages: []string{
			"Georgian",
		},
	},
	{
		Country: "Germany",
		Iso2:    "DE",
		Languages: []string{
			"German",
		},
	},
	{
		Country: "Ghana",
		Iso2:    "GH",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Greece",
		Iso2:    "GR",
		Languages: []string{
			"Greek",
		},
	},
	{
		Country: "Grenada",
		Iso2:    "GD",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Guatemala",
		Iso2:    "GT",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Guinea",
		Iso2:    "GN",
		Languages: []string{
			"French",
		},
	},
	{
		Country: "Guinea-Bissau",
		Iso2:    "GW",
		Languages: []string{
			"Portuguese",
		},
	},
	{
		Country: "Guyana",
		Iso2:    "GY",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Haiti",
		Iso2:    "HT",
		Languages: []string{
			"French",
		},
	},
	{
		Country: "Honduras",
		Iso2:    "HN",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Hong Kong",
		Iso2:    "HK",
		Languages: []string{
			"Cantonese",
			"English",
			"Mandarin",
		},
	},
	{
		Country: "Hungary",
		Iso2:    "HU",
		Languages: []string{
			"Hungarian",
		},
	},
	{
		Country: "Iceland",
		Iso2:    "IS",
		Languages: []string{
			"Icelandic",
		},
	},
	{
		Country: "India",
		Iso2:    "IN",
		Languages: []string{
			"Assamese",
			"Bengali",
			"Bodo",
			"Dogri",
			"Gujarati",
			"Hindi",
			"Kannada",
			"Kashmiri",
			"Konkani",
			"Maithili",
			"Malayalam",
			"Meitei (Manipuri)",
			"Marathi",
			"Nepali",
			"Odia (Oriya),",
			"Punjabi",
			"Sanskrit",
			"Santali",
			"Sindhi",
			"Tamil",
			"Telugu",
			"Urdu",
		},
	},
	{
		Country: "Indonesia",
		Iso2:    "ID",
		Languages: []string{
			"Indonesian",
		},
	},
	{
		Country: "Iran",
		Iso2:    "IR",
		Languages: []string{
			"Persian",
		},
	},
	{
		Country: "Iraq",
		Iso2:    "IQ",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Ireland",
		Iso2:    "IE",
		Languages: []string{
			"Irish",
			"English",
		},
	},
	{
		Country: "Israel",
		Iso2:    "IL",
		Languages: []string{
			"Hebrew",
		},
	},
	{
		Country: "Italy",
		Iso2:    "IT",
		Languages: []string{
			"Italian",
		},
	},
	{
		Country: "Ivory Coast",
		Iso2:    "CI",
		Languages: []string{
			"French",
		},
	},
	{
		Country: "Jamaica",
		Iso2:    "JM",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Japan",
		Iso2:    "JP",
		Languages: []string{
			"Japanese",
		},
	},
	{
		Country: "Jordan",
		Iso2:    "JO",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Kazakhstan",
		Iso2:    "KZ",
		Languages: []string{
			"Kazakh",
			"Russian",
		},
	},
	{
		Country: "Kenya",
		Iso2:    "KE",
		Languages: []string{
			"English",
			"Swahili",
		},
	},
	{
		Country: "North Korea",
		Iso2:    "KP",
		Languages: []string{
			"Korean",
		},
	},
	{
		Country: "South Korea",
		Iso2:    "KR",
		Languages: []string{
			"KKorean",
		},
	},
	{
		Country: "Kosovo", // je Srbija
		Iso2:    "XK",
		Languages: []string{
			"Albanian",
			"Serbian",
		},
	},
	{
		Country: "Kuwait",
		Iso2:    "KW",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Kyrgyzstan",
		Iso2:    "KG",
		Languages: []string{
			"Kyrgyz",
			"Russian",
		},
	},
	{
		Country: "Laos",
		Iso2:    "LA",
		Languages: []string{
			"Lao",
		},
	},
	{
		Country: "Latvia",
		Iso2:    "LV",
		Languages: []string{
			"Latvian",
		},
	},
	{
		Country: "Lebanon",
		Iso2:    "LB",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Lesotho",
		Iso2:    "LS",
		Languages: []string{
			"Sotho",
			"English",
		},
	},
	{
		Country: "Liberia",
		Iso2:    "LR",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Libya",
		Iso2:    "LY",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Liechtenstein",
		Iso2:    "LI",
		Languages: []string{
			"German",
		},
	},
	{
		Country: "Lithuania",
		Iso2:    "LT",
		Languages: []string{
			"Lithuanian",
		},
	},
	{
		Country: "Luxembourg",
		Iso2:    "LU",
		Languages: []string{
			"French",
			"German",
		},
	},
	{
		Country: "Macau",
		Iso2:    "MO",
		Languages: []string{
			"Cantonese",
			"Portuguese",
		},
	},
	{
		Country: "Madagascar",
		Iso2:    "MG",
		Languages: []string{
			"French",
			"Malagasy",
		},
	},
	{
		Country: "Malawi",
		Iso2:    "MW",
		Languages: []string{
			"English",
			"Chichewa",
		},
	},
	{
		Country: "Malaysia",
		Iso2:    "My",
		Languages: []string{
			"Malay",
		},
	},
	{
		Country: "Maldives",
		Iso2:    "MV",
		Languages: []string{
			"Dhivehi",
		},
	},
	{
		Country: "Mali",
		Iso2:    "ML",
		Languages: []string{
			"French",
		},
	},
	{
		Country: "Malta",
		Iso2:    "MT",
		Languages: []string{
			"English",
			"Maltese",
		},
	},
	{
		Country: "Marshall Islands",
		Iso2:    "MH",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Mauritania",
		Iso2:    "MR",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Mauritius",
		Iso2:    "MU",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Mexico",
		Iso2:    "MX",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Micronesia",
		Iso2:    "FM",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Moldova",
		Iso2:    "MD",
		Languages: []string{
			"Romanian",
		},
	},
	{
		Country: "Monaco",
		Iso2:    "MC",
		Languages: []string{
			"French",
		},
	},
	{
		Country: "Mongolia",
		Iso2:    "MN",
		Languages: []string{
			"Mongolian",
		},
	},
	{
		Country: "Montenegro",
		Iso2:    "ME",
		Languages: []string{
			"Montenegrin",
		},
	},
	{
		Country: "Morocco",
		Iso2:    "MA",
		Languages: []string{
			"Arabic",
			"Berber",
		},
	},
	{
		Country: "Mozambique",
		Iso2:    "MZ",
		Languages: []string{
			"Portuguese",
		},
	},
	{
		Country: "Myanmar",
		Iso2:    "MM",
		Languages: []string{
			"Burmese",
		},
	},
	{
		Country: "Namibia",
		Iso2:    "NA",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Nauru",
		Iso2:    "NR",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Nepal",
		Iso2:    "NP",
		Languages: []string{
			"Nepali",
		},
	},
	{
		Country: "Netherlands",
		Iso2:    "NL",
		Languages: []string{
			"Dutch",
		},
	},
	{
		Country: "New Zealand",
		Iso2:    "NZ",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Nicaragua",
		Iso2:    "NI",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Niger",
		Iso2:    "NE",
		Languages: []string{
			"French",
		},
	},
	{
		Country: "Nigeria",
		Iso2:    "NG",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Niue",
		Iso2:    "NU",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Norfolk Island",
		Iso2:    "NF",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "North Macedonia",
		Iso2:    "MK",
		Languages: []string{
			"Macedonian",
			"Albanian",
		},
	},
	{
		Country: "Norway",
		Iso2:    "NO",
		Languages: []string{
			"Norwegian",
		},
	},
	{
		Country: "Oman",
		Iso2:    "OM",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Pakistan",
		Iso2:    "PK",
		Languages: []string{
			"Urdu",
			"English",
		},
	},
	{
		Country: "Palau",
		Iso2:    "PW",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Palestine",
		Iso2:    "PS",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Panama",
		Iso2:    "PA",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Papua New Guinea",
		Iso2:    "PG",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Paraguay",
		Iso2:    "PY",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Peru",
		Iso2:    "PE",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Philippines",
		Iso2:    "PH",
		Languages: []string{
			"Filipino",
			"English",
		},
	},
	{
		Country: "Poland",
		Iso2:    "PL",
		Languages: []string{
			"Polish",
		},
	},
	{
		Country: "Portugal",
		Iso2:    "PT",
		Languages: []string{
			"Portuguese",
		},
	},
	{
		Country: "Qatar",
		Iso2:    "QA",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Romania",
		Iso2:    "RO",
		Languages: []string{
			"Romanian",
		},
	},
	{
		Country: "Russia",
		Iso2:    "RU",
		Languages: []string{
			"Russian",
		},
	},
	{
		Country: "Rwanda",
		Iso2:    "RW",
		Languages: []string{
			"English",
			"French",
			"Swahili",
		},
	},
	{
		Country: "Western Sahara",
		Iso2:    "EH",
		Languages: []string{
			"Arabic",
			"Spanish",
		},
	},
	{
		Country: "Saint Kitts and Nevis",
		Iso2:    "KN",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Saint Lucia",
		Iso2:    "LC",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Saint Vincent and the Grenadines",
		Iso2:    "VC",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Samoa",
		Iso2:    "WS",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "San Marino",
		Iso2:    "SM",
		Languages: []string{
			"Italian",
		},
	},
	{
		Country: "Sao Tome and Principe",
		Iso2:    "ST",
		Languages: []string{
			"Portuguese",
		},
	},
	{
		Country: "Saudi Arabia",
		Iso2:    "SA",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Senegal",
		Iso2:    "SN",
		Languages: []string{
			"French",
		},
	},
	{
		Country: "Serbia",
		Iso2:    "RS",
		Languages: []string{
			"Serbian",
		},
	},
	{
		Country: "Seychelles",
		Iso2:    "SC",
		Languages: []string{
			"English",
			"French",
		},
	},
	{
		Country: "Singapore",
		Iso2:    "SG",
		Languages: []string{
			"English",
			"Malay",
			"Mandarin Chinese",
			"Tamil",
		},
	},
	{
		Country: "Slovakia",
		Iso2:    "SK",
		Languages: []string{
			"Slovak",
		},
	},
	{
		Country: "Slovenia",
		Iso2:    "SI",
		Languages: []string{
			"Slovene",
		},
	},
	{
		Country: "Somalia",
		Iso2:    "SO",
		Languages: []string{
			"Arabic",
			"Somali",
		},
	},
	{
		Country: "South Africa",
		Iso2:    "ZA",
		Languages: []string{
			"Afrikaans",
			"English",
		},
	},
	{
		Country: "South Sudan",
		Iso2:    "SS",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Spain",
		Iso2:    "ES",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Sri Lanka",
		Iso2:    "LK",
		Languages: []string{
			"Sinhala",
			"Tamil",
		},
	},
	{
		Country: "Suriname",
		Iso2:    "SR",
		Languages: []string{
			"Dutch",
		},
	},
	{
		Country: "Sweden",
		Iso2:    "SE",
		Languages: []string{
			"Swedish",
		},
	},
	{
		Country: "Switzerland",
		Iso2:    "CH",
		Languages: []string{
			"French",
			"German",
			"Italian",
		},
	},
	{
		Country: "Syria",
		Iso2:    "SY",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Taiwan",
		Iso2:    "TW",
		Languages: []string{
			"Mandarin Chinese",
		},
	},
	{
		Country: "Tajikistan",
		Iso2:    "TJ",
		Languages: []string{
			"Tajik",
		},
	},
	{
		Country: "Tanzania",
		Iso2:    "TZ",
		Languages: []string{
			"Swahili",
			"English",
		},
	},
	{
		Country: "Thailand",
		Iso2:    "TH",
		Languages: []string{
			"Thai",
		},
	},
	{
		Country: "Togo",
		Iso2:    "TG",
		Languages: []string{
			"French",
		},
	},
	{
		Country: "Tokelau",
		Iso2:    "TK",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Tonga",
		Iso2:    "TO",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Trinidad and Tobago",
		Iso2:    "TT",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Tunisia",
		Iso2:    "TN",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Turkey",
		Iso2:    "TR",
		Languages: []string{
			"Turkish",
		},
	},
	{
		Country: "Turkmenistan",
		Iso2:    "TM",
		Languages: []string{
			"Turkmen",
		},
	},
	{
		Country: "Tuvalu",
		Iso2:    "TV",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Uganda",
		Iso2:    "UG",
		Languages: []string{
			"English",
			"Swahili",
		},
	},
	{
		Country: "Ukraine",
		Iso2:    "UA",
		Languages: []string{
			"Ukrainian",
			"Russian",
		},
	},
	{
		Country: "United Kingdom",
		Iso2:    "UK",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "United States",
		Iso2:    "US",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Uzbekistan",
		Iso2:    "UZ",
		Languages: []string{
			"Uzbek",
		},
	},
	{
		Country: "Vanuatu",
		Iso2:    "VU",
		Languages: []string{
			"English",
			"French",
		},
	},
	{
		Country: "Venezuela",
		Iso2:    "VE",
		Languages: []string{
			"Spanish",
		},
	},
	{
		Country: "Vietnam",
		Iso2:    "VN",
		Languages: []string{
			"Vietnamese",
		},
	},
	{
		Country: "Yemen",
		Iso2:    "YE",
		Languages: []string{
			"Arabic",
		},
	},
	{
		Country: "Zambia",
		Iso2:    "ZM",
		Languages: []string{
			"English",
		},
	},
	{
		Country: "Zimbabwe",
		Iso2:    "ZW",
		Languages: []string{
			"English",
		},
	},
}
