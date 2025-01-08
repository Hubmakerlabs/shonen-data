package locale

type Language struct {
	Country   st
	Iso2      st
	Languages []st
}

func GetLanguageByIso2(iso2 st) (v *Language) {
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
		Languages: []st{
			"Abkhaz",
			"Russian",
		},
	},
	{
		Country: "Afghanistan",
		Iso2:    "AF",
		Languages: []st{
			"Persian",
			"Pashto",
		},
	},
	{
		Country: "Albania",
		Iso2:    "AL",
		Languages: []st{
			"Albanian",
		},
	},
	{
		Country: "Algeria",
		Iso2:    "DZ",
		Languages: []st{
			"Arabic",
			"Berber",
		},
	},
	{
		Country: "Andorra",
		Iso2:    "AD",
		Languages: []st{
			"Catalan",
		},
	},
	{
		Country: "Angola",
		Iso2:    "AO",
		Languages: []st{
			"Portuguese",
		},
	},
	{
		Country: "Antigua and Barbuda",
		Iso2:    "AG",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Argentina",
		Iso2:    "AR",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Armenia",
		Iso2:    "AM",
		Languages: []st{
			"Armenian",
		},
	},
	{
		Country: "Australia",
		Iso2:    "AU",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Austria",
		Iso2:    "AT",
		Languages: []st{
			"German",
		},
	},
	{
		Country: "Azerbaijan",
		Iso2:    "AZ",
		Languages: []st{
			"Azerbaijani",
		},
	},
	{
		Country: "Bahamas",
		Iso2:    "BS",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Bahrain",
		Iso2:    "BH",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Bangladesh",
		Iso2:    "BD",
		Languages: []st{
			"Bengali",
		},
	},
	{
		Country: "Barbados",
		Iso2:    "BB",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Belarus",
		Iso2:    "BY",
		Languages: []st{
			"Belarusian",
			"Russian",
		},
	},
	{
		Country: "Belgium",
		Iso2:    "BE",
		Languages: []st{
			"Dutch",
			"French",
			"German",
		},
	},
	{
		Country: "Belize",
		Iso2:    "BZ",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Benin",
		Iso2:    "BJ",
		Languages: []st{
			"French",
		},
	},
	{
		Country: "Bhutan",
		Iso2:    "BT",
		Languages: []st{
			"Dzongkha",
		},
	},
	{
		Country: "Bolivia",
		Iso2:    "BO",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Bosnia and Herzegovina",
		Iso2:    "BA",
		Languages: []st{
			"Bosnian",
			"Croatian",
			"Serbian",
		},
	},
	{
		Country: "Botswana",
		Iso2:    "BW",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Brazil",
		Iso2:    "BR",
		Languages: []st{
			"Portuguese",
		},
	},
	{
		Country: "Brunei",
		Iso2:    "BN",
		Languages: []st{
			"Malay",
		},
	},
	{
		Country: "Burkina Faso",
		Iso2:    "BF",
		Languages: []st{
			"English",
			"French",
		},
	},
	{
		Country: "Burundi",
		Iso2:    "BI",
		Languages: []st{
			"French",
			"Kirundi",
			"English",
		},
	},
	{
		Country: "Cambodia",
		Iso2:    "KH",
		Languages: []st{
			"Khmer",
		},
	},
	{
		Country: "Cameroon",
		Iso2:    "CM",
		Languages: []st{
			"English",
			"French",
		},
	},
	{
		Country: "Canada",
		Iso2:    "CA",
		Languages: []st{
			"English",
			"French",
		},
	},
	{
		Country: "Cape Verde",
		Iso2:    "CV",
		Languages: []st{
			"Portuguese",
		},
	},
	{
		Country: "Central African Republic",
		Iso2:    "CF",
		Languages: []st{
			"French",
			"Sango",
		},
	},
	{
		Country: "Chad",
		Iso2:    "TD",
		Languages: []st{
			"Arabic",
			"French",
		},
	},
	{
		Country: "Chile",
		Iso2:    "CL",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "China",
		Iso2:    "CN",
		Languages: []st{
			"Chinese",
		},
	},
	{
		Country: "Christmas Island",
		Iso2:    "CX",
		Languages: []st{
			"English",
			"Mandarin Chinese",
			"Malay",
		},
	},
	{
		Country: "Cocos Islands",
		Iso2:    "CC",
		Languages: []st{
			"English",
			"Cocos Malay",
		},
	},
	{
		Country: "Columbia",
		Iso2:    "CO",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Comoros",
		Iso2:    "KM",
		Languages: []st{
			"Arabic",
			"French",
			"Comorian",
		},
	},
	{
		Country: "Democratic Republic of the Congo",
		Iso2:    "CD",
		Languages: []st{
			"French",
		},
	},
	{
		Country: "Congo",
		Iso2:    "CG",
		Languages: []st{
			"French",
		},
	},
	{
		Country: "Cook Islands",
		Iso2:    "CK",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Costa Rica",
		Iso2:    "CR",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Croatia",
		Iso2:    "HR",
		Languages: []st{
			"Croatian",
		},
	},
	{
		Country: "Cuba",
		Iso2:    "CU",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Cyprus",
		Iso2:    "CY",
		Languages: []st{
			"Greek",
			"Turkish",
		},
	},
	{
		Country: "Czech Republic",
		Iso2:    "CZ",
		Languages: []st{
			"Czech",
			"Slovak",
		},
	},
	{
		Country: "Denmark",
		Iso2:    "DK",
		Languages: []st{
			"Danish",
		},
	},
	{
		Country: "Djibouti",
		Iso2:    "DJ",
		Languages: []st{
			"Arabic",
			"French",
		},
	},
	{
		Country: "Dominica",
		Iso2:    "DM",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Dominican Republic",
		Iso2:    "DO",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Timor-Leste",
		Iso2:    "TL",
		Languages: []st{
			"Portuguese",
		},
	},
	{
		Country: "Ecuador",
		Iso2:    "EC",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Egypt",
		Iso2:    "EG",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "El Salvador",
		Iso2:    "SV",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Equatorial Guinea",
		Iso2:    "GQ",
		Languages: []st{
			"French",
			"Portuguese",
			"Spanish",
		},
	},
	{
		Country: "Eritrea",
		Iso2:    "ER",
		Languages: []st{
			"Tigrinya",
		},
	},
	{
		Country: "Estonia",
		Iso2:    "EE",
		Languages: []st{
			"Estonian",
		},
	},
	{
		Country: "Eswatini",
		Iso2:    "SZ",
		Languages: []st{
			"English",
			"Swazi",
		},
	},
	{
		Country: "Ethiopia",
		Iso2:    "ET",
		Languages: []st{
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
		Languages: []st{
			"English",
			"Fijian",
			"Hindi",
		},
	},
	{
		Country: "Finland",
		Iso2:    "FI",
		Languages: []st{
			"Finnish",
		},
	},
	{
		Country: "France",
		Iso2:    "FR",
		Languages: []st{
			"French",
		},
	},
	{
		Country: "Gabon",
		Iso2:    "GA",
		Languages: []st{
			"French",
		},
	},
	{
		Country: "Gambia",
		Iso2:    "GM",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Georgia",
		Iso2:    "GE",
		Languages: []st{
			"Georgian",
		},
	},
	{
		Country: "Germany",
		Iso2:    "DE",
		Languages: []st{
			"German",
		},
	},
	{
		Country: "Ghana",
		Iso2:    "GH",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Greece",
		Iso2:    "GR",
		Languages: []st{
			"Greek",
		},
	},
	{
		Country: "Grenada",
		Iso2:    "GD",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Guatemala",
		Iso2:    "GT",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Guinea",
		Iso2:    "GN",
		Languages: []st{
			"French",
		},
	},
	{
		Country: "Guinea-Bissau",
		Iso2:    "GW",
		Languages: []st{
			"Portuguese",
		},
	},
	{
		Country: "Guyana",
		Iso2:    "GY",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Haiti",
		Iso2:    "HT",
		Languages: []st{
			"French",
		},
	},
	{
		Country: "Honduras",
		Iso2:    "HN",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Hong Kong",
		Iso2:    "HK",
		Languages: []st{
			"Cantonese",
			"English",
			"Mandarin",
		},
	},
	{
		Country: "Hungary",
		Iso2:    "HU",
		Languages: []st{
			"Hungarian",
		},
	},
	{
		Country: "Iceland",
		Iso2:    "IS",
		Languages: []st{
			"Icelandic",
		},
	},
	{
		Country: "India",
		Iso2:    "IN",
		Languages: []st{
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
		Languages: []st{
			"Indonesian",
		},
	},
	{
		Country: "Iran",
		Iso2:    "IR",
		Languages: []st{
			"Persian",
		},
	},
	{
		Country: "Iraq",
		Iso2:    "IQ",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Ireland",
		Iso2:    "IE",
		Languages: []st{
			"Irish",
			"English",
		},
	},
	{
		Country: "Israel",
		Iso2:    "IL",
		Languages: []st{
			"Hebrew",
		},
	},
	{
		Country: "Italy",
		Iso2:    "IT",
		Languages: []st{
			"Italian",
		},
	},
	{
		Country: "Ivory Coast",
		Iso2:    "CI",
		Languages: []st{
			"French",
		},
	},
	{
		Country: "Jamaica",
		Iso2:    "JM",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Japan",
		Iso2:    "JP",
		Languages: []st{
			"Japanese",
		},
	},
	{
		Country: "Jordan",
		Iso2:    "JO",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Kazakhstan",
		Iso2:    "KZ",
		Languages: []st{
			"Kazakh",
			"Russian",
		},
	},
	{
		Country: "Kenya",
		Iso2:    "KE",
		Languages: []st{
			"English",
			"Swahili",
		},
	},
	{
		Country: "North Korea",
		Iso2:    "KP",
		Languages: []st{
			"Korean",
		},
	},
	{
		Country: "South Korea",
		Iso2:    "KR",
		Languages: []st{
			"KKorean",
		},
	},
	{
		Country: "Kosovo", // je Srbija
		Iso2:    "XK",
		Languages: []st{
			"Albanian",
			"Serbian",
		},
	},
	{
		Country: "Kuwait",
		Iso2:    "KW",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Kyrgyzstan",
		Iso2:    "KG",
		Languages: []st{
			"Kyrgyz",
			"Russian",
		},
	},
	{
		Country: "Laos",
		Iso2:    "LA",
		Languages: []st{
			"Lao",
		},
	},
	{
		Country: "Latvia",
		Iso2:    "LV",
		Languages: []st{
			"Latvian",
		},
	},
	{
		Country: "Lebanon",
		Iso2:    "LB",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Lesotho",
		Iso2:    "LS",
		Languages: []st{
			"Sotho",
			"English",
		},
	},
	{
		Country: "Liberia",
		Iso2:    "LR",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Libya",
		Iso2:    "LY",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Liechtenstein",
		Iso2:    "LI",
		Languages: []st{
			"German",
		},
	},
	{
		Country: "Lithuania",
		Iso2:    "LT",
		Languages: []st{
			"Lithuanian",
		},
	},
	{
		Country: "Luxembourg",
		Iso2:    "LU",
		Languages: []st{
			"French",
			"German",
		},
	},
	{
		Country: "Macau",
		Iso2:    "MO",
		Languages: []st{
			"Cantonese",
			"Portuguese",
		},
	},
	{
		Country: "Madagascar",
		Iso2:    "MG",
		Languages: []st{
			"French",
			"Malagasy",
		},
	},
	{
		Country: "Malawi",
		Iso2:    "MW",
		Languages: []st{
			"English",
			"Chichewa",
		},
	},
	{
		Country: "Malaysia",
		Iso2:    "My",
		Languages: []st{
			"Malay",
		},
	},
	{
		Country: "Maldives",
		Iso2:    "MV",
		Languages: []st{
			"Dhivehi",
		},
	},
	{
		Country: "Mali",
		Iso2:    "ML",
		Languages: []st{
			"French",
		},
	},
	{
		Country: "Malta",
		Iso2:    "MT",
		Languages: []st{
			"English",
			"Maltese",
		},
	},
	{
		Country: "Marshall Islands",
		Iso2:    "MH",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Mauritania",
		Iso2:    "MR",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Mauritius",
		Iso2:    "MU",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Mexico",
		Iso2:    "MX",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Micronesia",
		Iso2:    "FM",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Moldova",
		Iso2:    "MD",
		Languages: []st{
			"Romanian",
		},
	},
	{
		Country: "Monaco",
		Iso2:    "MC",
		Languages: []st{
			"French",
		},
	},
	{
		Country: "Mongolia",
		Iso2:    "MN",
		Languages: []st{
			"Mongolian",
		},
	},
	{
		Country: "Montenegro",
		Iso2:    "ME",
		Languages: []st{
			"Montenegrin",
		},
	},
	{
		Country: "Morocco",
		Iso2:    "MA",
		Languages: []st{
			"Arabic",
			"Berber",
		},
	},
	{
		Country: "Mozambique",
		Iso2:    "MZ",
		Languages: []st{
			"Portuguese",
		},
	},
	{
		Country: "Myanmar",
		Iso2:    "MM",
		Languages: []st{
			"Burmese",
		},
	},
	{
		Country: "Namibia",
		Iso2:    "NA",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Nauru",
		Iso2:    "NR",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Nepal",
		Iso2:    "NP",
		Languages: []st{
			"Nepali",
		},
	},
	{
		Country: "Netherlands",
		Iso2:    "NL",
		Languages: []st{
			"Dutch",
		},
	},
	{
		Country: "New Zealand",
		Iso2:    "NZ",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Nicaragua",
		Iso2:    "NI",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Niger",
		Iso2:    "NE",
		Languages: []st{
			"French",
		},
	},
	{
		Country: "Nigeria",
		Iso2:    "NG",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Niue",
		Iso2:    "NU",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Norfolk Island",
		Iso2:    "NF",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "North Macedonia",
		Iso2:    "MK",
		Languages: []st{
			"Macedonian",
			"Albanian",
		},
	},
	{
		Country: "Norway",
		Iso2:    "NO",
		Languages: []st{
			"Norwegian",
		},
	},
	{
		Country: "Oman",
		Iso2:    "OM",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Pakistan",
		Iso2:    "PK",
		Languages: []st{
			"Urdu",
			"English",
		},
	},
	{
		Country: "Palau",
		Iso2:    "PW",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Palestine",
		Iso2:    "PS",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Panama",
		Iso2:    "PA",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Papua New Guinea",
		Iso2:    "PG",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Paraguay",
		Iso2:    "PY",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Peru",
		Iso2:    "PE",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Philippines",
		Iso2:    "PH",
		Languages: []st{
			"Filipino",
			"English",
		},
	},
	{
		Country: "Poland",
		Iso2:    "PL",
		Languages: []st{
			"Polish",
		},
	},
	{
		Country: "Portugal",
		Iso2:    "PT",
		Languages: []st{
			"Portuguese",
		},
	},
	{
		Country: "Qatar",
		Iso2:    "QA",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Romania",
		Iso2:    "RO",
		Languages: []st{
			"Romanian",
		},
	},
	{
		Country: "Russia",
		Iso2:    "RU",
		Languages: []st{
			"Russian",
		},
	},
	{
		Country: "Rwanda",
		Iso2:    "RW",
		Languages: []st{
			"English",
			"French",
			"Swahili",
		},
	},
	{
		Country: "Western Sahara",
		Iso2:    "EH",
		Languages: []st{
			"Arabic",
			"Spanish",
		},
	},
	{
		Country: "Saint Kitts and Nevis",
		Iso2:    "KN",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Saint Lucia",
		Iso2:    "LC",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Saint Vincent and the Grenadines",
		Iso2:    "VC",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Samoa",
		Iso2:    "WS",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "San Marino",
		Iso2:    "SM",
		Languages: []st{
			"Italian",
		},
	},
	{
		Country: "Sao Tome and Principe",
		Iso2:    "ST",
		Languages: []st{
			"Portuguese",
		},
	},
	{
		Country: "Saudi Arabia",
		Iso2:    "SA",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Senegal",
		Iso2:    "SN",
		Languages: []st{
			"French",
		},
	},
	{
		Country: "Serbia",
		Iso2:    "RS",
		Languages: []st{
			"Serbian",
		},
	},
	{
		Country: "Seychelles",
		Iso2:    "SC",
		Languages: []st{
			"English",
			"French",
		},
	},
	{
		Country: "Singapore",
		Iso2:    "SG",
		Languages: []st{
			"English",
			"Malay",
			"Mandarin Chinese",
			"Tamil",
		},
	},
	{
		Country: "Slovakia",
		Iso2:    "SK",
		Languages: []st{
			"Slovak",
		},
	},
	{
		Country: "Slovenia",
		Iso2:    "SI",
		Languages: []st{
			"Slovene",
		},
	},
	{
		Country: "Somalia",
		Iso2:    "SO",
		Languages: []st{
			"Arabic",
			"Somali",
		},
	},
	{
		Country: "South Africa",
		Iso2:    "ZA",
		Languages: []st{
			"Afrikaans",
			"English",
		},
	},
	{
		Country: "South Sudan",
		Iso2:    "SS",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Spain",
		Iso2:    "ES",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Sri Lanka",
		Iso2:    "LK",
		Languages: []st{
			"Sinhala",
			"Tamil",
		},
	},
	{
		Country: "Suriname",
		Iso2:    "SR",
		Languages: []st{
			"Dutch",
		},
	},
	{
		Country: "Sweden",
		Iso2:    "SE",
		Languages: []st{
			"Swedish",
		},
	},
	{
		Country: "Switzerland",
		Iso2:    "CH",
		Languages: []st{
			"French",
			"German",
			"Italian",
		},
	},
	{
		Country: "Syria",
		Iso2:    "SY",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Taiwan",
		Iso2:    "TW",
		Languages: []st{
			"Mandarin Chinese",
		},
	},
	{
		Country: "Tajikistan",
		Iso2:    "TJ",
		Languages: []st{
			"Tajik",
		},
	},
	{
		Country: "Tanzania",
		Iso2:    "TZ",
		Languages: []st{
			"Swahili",
			"English",
		},
	},
	{
		Country: "Thailand",
		Iso2:    "TH",
		Languages: []st{
			"Thai",
		},
	},
	{
		Country: "Togo",
		Iso2:    "TG",
		Languages: []st{
			"French",
		},
	},
	{
		Country: "Tokelau",
		Iso2:    "TK",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Tonga",
		Iso2:    "TO",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Trinidad and Tobago",
		Iso2:    "TT",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Tunisia",
		Iso2:    "TN",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Turkey",
		Iso2:    "TR",
		Languages: []st{
			"Turkish",
		},
	},
	{
		Country: "Turkmenistan",
		Iso2:    "TM",
		Languages: []st{
			"Turkmen",
		},
	},
	{
		Country: "Tuvalu",
		Iso2:    "TV",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Uganda",
		Iso2:    "UG",
		Languages: []st{
			"English",
			"Swahili",
		},
	},
	{
		Country: "Ukraine",
		Iso2:    "UA",
		Languages: []st{
			"Ukrainian",
			"Russian",
		},
	},
	{
		Country: "United Kingdom",
		Iso2:    "UK",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "United States",
		Iso2:    "US",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Uzbekistan",
		Iso2:    "UZ",
		Languages: []st{
			"Uzbek",
		},
	},
	{
		Country: "Vanuatu",
		Iso2:    "VU",
		Languages: []st{
			"English",
			"French",
		},
	},
	{
		Country: "Venezuela",
		Iso2:    "VE",
		Languages: []st{
			"Spanish",
		},
	},
	{
		Country: "Vietnam",
		Iso2:    "VN",
		Languages: []st{
			"Vietnamese",
		},
	},
	{
		Country: "Yemen",
		Iso2:    "YE",
		Languages: []st{
			"Arabic",
		},
	},
	{
		Country: "Zambia",
		Iso2:    "ZM",
		Languages: []st{
			"English",
		},
	},
	{
		Country: "Zimbabwe",
		Iso2:    "ZW",
		Languages: []st{
			"English",
		},
	},
}
