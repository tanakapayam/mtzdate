package main

import (
	"fmt"
	"os"
	"regexp"

	docopt "github.com/docopt/docopt-go"
	"github.com/fatih/color"
)

func init() {
	var err error

	usage := `Usage:
  ` + prog + ` (-h | --version)
  ` + prog + `
  ` + prog + ` --loop

Description
  This command-line utility displays Unix date in multiple time zones
  based on environment variables.

  With MTZDATE_LOOP=1 or --loop, ` + prog + ` will refresh the screen once a second.
  Control-C will break the loop.

Options:
  -h, --help
  -l, --loop    # Loop until Control-C is trapped
  --version

Installation
  go get -u -v -ldflags="-s -w" github.com/tanakapayam/` + prog + `

Environment
  Set MTZDATE_TIMEZONES to a comma-separated list of time zones. If desired, preface each time zone with a UTF-8-encoded city name or alias and a colon.

  To see emoji flags, set MTZDATE_FLAGS to a comma-separated map of UTF-8-encoded city names or aliases followd by two-letter country code -- separated by a colon.

  ` + prog + ` defaults to coloring workhours green and to coloring pre- and post-workhours yellow. The behavior is controlled by the following environment variables (with their default values):

  export MTZDATE_WORKDAYS='Mon,Tue,Wed,Thu,Fri'
  export MTZDATE_GREEN_HOURS='8-17'
  export MTZDATE_YELLOW_HOURS='7-8,17-18'
  export MTZDATE_FAINT_HOURS='0-7,22-24'

  To opt out of the feature, set MTZDATE_WORKDAYS='':

  export MTZDATE_WORKDAYS=''

  MTZDATE_FORMAT can be set to a sequence of "d" (date), "f" (flag) and "c" (city) to signify the display format. (If unset, "dfc" is assumed.) Naturally, it's most meaningful if it's three letters, but there are no restrictions.

Examples
  $ export MTZDATE_TIMEZONES='America/Chicago,Europe/Paris'
  $ export MTZDATE_FLAGS='Chicago:US,Paris:FR'

  $ date
  Sun Jul 29 18:10:33 PDT 2018

  $ mtzdate
  Sun Jul 29 20:10:33 CDT   🇺🇸  Chicago
  Mon Jul 30 03:10:33 CEST  🇫🇷  Paris

  $ export MTZDATE_TIMEZONES='San Francisco:America/Los_Angeles,UTC,München:Europe/Berlin,काठमाडौं:Asia/Kathmandu,東京:Asia/Tokyo'
  $ export MTZDATE_FLAGS='San Francisco:US,München:DE,काठमाडौं:NP,東京:JP'

  $ mtzdate
  Sun Jul 29 18:10:33 PDT   🇺🇸  San Francisco
  Mon Jul 30 01:10:33 UTC   ☁️
  Mon Jul 30 03:10:33 CEST  🇩🇪  München
  Mon Jul 30 06:55:33 +0545 🇳🇵  काठमाडौं
  Mon Jul 30 10:10:33 JST   🇯🇵  東京

  $ export MTZDATE_FORMAT='fd'
  $ mtzdate
  🇺🇸  Sun Jul 29 18:10:33 PDT
  ☁️  Mon Jul 30 01:10:33 UTC
  🇩🇪  Mon Jul 30 03:10:33 CEST
  🇳🇵  Mon Jul 30 06:55:33 +0545
  🇯🇵  Mon Jul 30 10:10:33 JST

See Also
  /usr/share/zoneinfo
  https://github.com/tanakapayam/mtzdate
`

	// http://docopt.org/
	parser := &docopt.Parser{
		HelpHandler: func(err error, usage string) {
			var re = regexp.MustCompile(`((?:^|\n)\S+.*\n)`)
			usage = re.ReplaceAllString(usage, bold("$1"))

			if err != nil {
				_, err = fmt.Fprintln(os.Stderr, usage)
				die(err)
				os.Exit(1)
			} else {
				fmt.Println(usage)
				os.Exit(0)
			}
		},
	}

	args, err = parser.ParseArgs(
		usage,
		os.Args[1:],
		version,
	)
	die(err)

	if args["--loop"].(bool) {
		err := os.Setenv("MTZDATE_LOOP", "1")
		die(err)
	}

	_, set := os.LookupEnv("MTZDATE_FORMAT")
	if !set {
		err := os.Setenv("MTZDATE_FORMAT", mtzdateFormat)
		die(err)
	}
}

const (
	utc                = "UTC"
	mtzdateWorkdays    = "Mon,Tue,Wed,Thu,Fri"
	mtzdateGreenHours  = "8-17"
	mtzdateYellowHours = "7-8,17-18"
	mtzdateFaintHours  = "0-7,22-24"
	mtzdateFormat      = "dfc"
)

var (
	args docopt.Opts

	prog    = os.Args[0]
	version = "1.0.0"

	bold   = color.New(color.Bold).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	faint  = color.New(color.Faint).SprintFunc()

	workday     map[string]bool
	greenHours  [][]int
	yellowHours [][]int
	faintHours  [][]int

	// https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2
	countryCode = map[string]string{
		"AF": "Afghanistan",
		"AX": "Åland Islands",
		"AL": "Albania",
		"DZ": "Algeria",
		"AS": "American Samoa",
		"AD": "Andorra",
		"AO": "Angola",
		"AI": "Anguilla",
		"AQ": "Antarctica",
		"AG": "Antigua & Barbuda",
		"AR": "Argentina",
		"AM": "Armenia",
		"AW": "Aruba",
		"AC": "Ascension Island",
		"AU": "Australia",
		"AT": "Austria",
		"AZ": "Azerbaijan",
		"BS": "Bahamas",
		"BH": "Bahrain",
		"BD": "Bangladesh",
		"BB": "Barbados",
		"BY": "Belarus",
		"BE": "Belgium",
		"BZ": "Belize",
		"BJ": "Benin",
		"BM": "Bermuda",
		"BT": "Bhutan",
		"BO": "Bolivia",
		"BA": "Bosnia & Herzegovina",
		"BW": "Botswana",
		"BV": "Bouvet Island",
		"BR": "Brazil",
		"IO": "British Indian Ocean Territory",
		"VG": "British Virgin Islands",
		"BN": "Brunei",
		"BG": "Bulgaria",
		"BF": "Burkina Faso",
		"BI": "Burundi",
		"KH": "Cambodia",
		"CM": "Cameroon",
		"CA": "Canada",
		"IC": "Canary Islands",
		"CV": "Cape Verde",
		"BQ": "Caribbean Netherlands",
		"KY": "Cayman Islands",
		"CF": "Central African Republic",
		"EA": "Ceuta & Melilla",
		"TD": "Chad",
		"CL": "Chile",
		"CN": "China",
		"CX": "Christmas Island",
		"CP": "Clipperton Island",
		"CC": "Cocos (Keeling) Islands",
		"CO": "Colombia",
		"KM": "Comoros",
		"CG": "Congo - Brazzaville",
		"CD": "Congo - Kinshasa",
		"CK": "Cook Islands",
		"CR": "Costa Rica",
		"CI": "Côte D’Ivoire",
		"HR": "Croatia",
		"CU": "Cuba",
		"CW": "Curaçao",
		"CY": "Cyprus",
		"CZ": "Czechia",
		"DK": "Denmark",
		"DG": "Diego Garcia",
		"DJ": "Djibouti",
		"DM": "Dominica",
		"DO": "Dominican Republic",
		"EC": "Ecuador",
		"EG": "Egypt",
		"SV": "El Salvador",
		"GQ": "Equatorial Guinea",
		"ER": "Eritrea",
		"EE": "Estonia",
		"ET": "Ethiopia",
		"EU": "European Union",
		"FK": "Falkland Islands",
		"FO": "Faroe Islands",
		"FJ": "Fiji",
		"FI": "Finland",
		"FR": "France",
		"GF": "French Guiana",
		"PF": "French Polynesia",
		"TF": "French Southern Territories",
		"GA": "Gabon",
		"GM": "Gambia",
		"GE": "Georgia",
		"DE": "Germany",
		"GH": "Ghana",
		"GI": "Gibraltar",
		"GR": "Greece",
		"GL": "Greenland",
		"GD": "Grenada",
		"GP": "Guadeloupe",
		"GU": "Guam",
		"GT": "Guatemala",
		"GG": "Guernsey",
		"GW": "Guinea-Bissau",
		"GN": "Guinea",
		"GY": "Guyana",
		"HT": "Haiti",
		"HM": "Heard & McDonald Islands",
		"HN": "Honduras",
		"HK": "Hong Kong SAR China",
		"HU": "Hungary",
		"IS": "Iceland",
		"IN": "India",
		"ID": "Indonesia",
		"IR": "Iran",
		"IQ": "Iraq",
		"IE": "Ireland",
		"IM": "Isle of Man",
		"IL": "Israel",
		"IT": "Italy",
		"JM": "Jamaica",
		"JP": "Japan",
		"JE": "Jersey",
		"JO": "Jordan",
		"KZ": "Kazakhstan",
		"KE": "Kenya",
		"KI": "Kiribati",
		"XK": "Kosovo",
		"KW": "Kuwait",
		"KG": "Kyrgyzstan",
		"LA": "Laos",
		"LV": "Latvia",
		"LB": "Lebanon",
		"LS": "Lesotho",
		"LR": "Liberia",
		"LY": "Libya",
		"LI": "Liechtenstein",
		"LT": "Lithuania",
		"LU": "Luxembourg",
		"MO": "Macau SAR China",
		"MK": "Macedonia",
		"MG": "Madagascar",
		"MW": "Malawi",
		"MY": "Malaysia",
		"MV": "Maldives",
		"ML": "Mali",
		"MT": "Malta",
		"MH": "Marshall Islands",
		"MQ": "Martinique",
		"MR": "Mauritania",
		"MU": "Mauritius",
		"YT": "Mayotte",
		"MX": "Mexico",
		"FM": "Micronesia",
		"MD": "Moldova",
		"MC": "Monaco",
		"MN": "Mongolia",
		"ME": "Montenegro",
		"MS": "Montserrat",
		"MA": "Morocco",
		"MZ": "Mozambique",
		"MM": "Myanmar",
		"NA": "Namibia",
		"NR": "Nauru",
		"NP": "Nepal",
		"NL": "Netherlands",
		"NC": "New Caledonia",
		"NZ": "New Zealand",
		"NI": "Nicaragua",
		"NE": "Niger",
		"NG": "Nigeria",
		"NU": "Niue",
		"NF": "Norfolk Island",
		"KP": "North Korea",
		"MP": "Northern Mariana Islands",
		"NO": "Norway",
		"OM": "Oman",
		"PK": "Pakistan",
		"PW": "Palau",
		"PS": "Palestinian Territories",
		"PA": "Panama",
		"PG": "Papua New Guinea",
		"PY": "Paraguay",
		"PE": "Peru",
		"PH": "Philippines",
		"PN": "Pitcairn Islands",
		"PL": "Poland",
		"PT": "Portugal",
		"PR": "Puerto Rico",
		"QA": "Qatar",
		"RE": "Réunion",
		"RO": "Romania",
		"RU": "Russia",
		"RW": "Rwanda",
		"WS": "Samoa",
		"SM": "San Marino",
		"ST": "São Tomé & Príncipe",
		"SA": "Saudi Arabia",
		"SN": "Senegal",
		"RS": "Serbia",
		"SC": "Seychelles",
		"SL": "Sierra Leone",
		"SG": "Singapore",
		"SX": "Sint Maarten",
		"SK": "Slovakia",
		"SI": "Slovenia",
		"SB": "Solomon Islands",
		"SO": "Somalia",
		"ZA": "South Africa",
		"GS": "South Georgia & South Sandwich Islands",
		"KR": "South Korea",
		"SS": "South Sudan",
		"ES": "Spain",
		"LK": "Sri Lanka",
		"BL": "St. Barthélemy",
		"SH": "St. Helena",
		"KN": "St. Kitts & Nevis",
		"LC": "St. Lucia",
		"MF": "St. Martin",
		"PM": "St. Pierre & Miquelon",
		"VC": "St. Vincent & Grenadines",
		"SD": "Sudan",
		"SR": "Suriname",
		"SJ": "Svalbard & Jan Mayen",
		"SZ": "Swaziland",
		"SE": "Sweden",
		"CH": "Switzerland",
		"SY": "Syria",
		"TW": "Taiwan",
		"TJ": "Tajikistan",
		"TZ": "Tanzania",
		"TH": "Thailand",
		"TL": "Timor-Leste",
		"TG": "Togo",
		"TK": "Tokelau",
		"TO": "Tonga",
		"TT": "Trinidad & Tobago",
		"TN": "Tunisia",
		"TR": "Turkey",
		"TM": "Turkmenistan",
		"TC": "Turks & Caicos Islands",
		"TV": "Tuvalu",
		"UG": "Uganda",
		"UA": "Ukraine",
		"AE": "United Arab Emirates",
		"GB": "United Kingdom",
		"US": "United States",
		"UY": "Uruguay",
		"UM": "U.S. Outlying Islands",
		"VI": "U.S. Virgin Islands",
		"UZ": "Uzbekistan",
		"VU": "Vanuatu",
		"VA": "Vatican City",
		"VE": "Venezuela",
		"VN": "Vietnam",
		"WF": "Wallis & Futuna",
		"EH": "Western Sahara",
		"YE": "Yemen",
		"ZM": "Zambia",
		"ZW": "Zimbabwe",
	}

	// https://emojipedia.org/flags/
	flag = map[string]string{
		"Afghanistan":          "🇦🇫 ",
		"Åland Islands":        "🇦🇽 ",
		"Albania":              "🇦🇱 ",
		"Algeria":              "🇩🇿 ",
		"American Samoa":       "🇦🇸 ",
		"Andorra":              "🇦🇩 ",
		"Angola":               "🇦🇴 ",
		"Anguilla":             "🇦🇮 ",
		"Antarctica":           "🇦🇶 ",
		"Antigua & Barbuda":    "🇦🇬 ",
		"Argentina":            "🇦🇷 ",
		"Armenia":              "🇦🇲 ",
		"Aruba":                "🇦🇼 ",
		"Ascension Island":     "🇦🇨 ",
		"Australia":            "🇦🇺 ",
		"Austria":              "🇦🇹 ",
		"Azerbaijan":           "🇦🇿 ",
		"Bahamas":              "🇧🇸 ",
		"Bahrain":              "🇧🇭 ",
		"Bangladesh":           "🇧🇩 ",
		"Barbados":             "🇧🇧 ",
		"Belarus":              "🇧🇾 ",
		"Belgium":              "🇧🇪 ",
		"Belize":               "🇧🇿 ",
		"Benin":                "🇧🇯 ",
		"Bermuda":              "🇧🇲 ",
		"Bhutan":               "🇧🇹 ",
		"Bolivia":              "🇧🇴 ",
		"Bosnia & Herzegovina": "🇧🇦 ",
		"Botswana":             "🇧🇼 ",
		"Bouvet Island":        "🇧🇻 ",
		"Brazil":               "🇧🇷 ",
		"British Indian Ocean Territory": "🇮🇴 ",
		"British Virgin Islands":         "🇻🇬 ",
		"Brunei":                         "🇧🇳 ",
		"Bulgaria":                       "🇧🇬 ",
		"Burkina Faso":                   "🇧🇫 ",
		"Burundi":                        "🇧🇮 ",
		"Cambodia":                       "🇰🇭 ",
		"Cameroon":                       "🇨🇲 ",
		"Canada":                         "🇨🇦 ",
		"Canary Islands":                 "🇮🇨 ",
		"Cape Verde":                     "🇨🇻 ",
		"Caribbean Netherlands":          "🇧🇶 ",
		"Cayman Islands":                 "🇰🇾 ",
		"Central African Republic":       "🇨🇫 ",
		"Ceuta & Melilla":                "🇪🇦 ",
		"Chad":                           "🇹🇩 ",
		"Chile":                          "🇨🇱 ",
		"China":                          "🇨🇳 ",
		"Christmas Island":               "🇨🇽 ",
		"Clipperton Island":              "🇨🇵 ",
		"Cocos (Keeling) Islands":        "🇨🇨 ",
		"Colombia":                       "🇨🇴 ",
		"Comoros":                        "🇰🇲 ",
		"Congo - Brazzaville":            "🇨🇬 ",
		"Congo - Kinshasa":               "🇨🇩 ",
		"Cook Islands":                   "🇨🇰 ",
		"Costa Rica":                     "🇨🇷 ",
		"Côte D’Ivoire":                  "🇨🇮 ",
		"Croatia":                        "🇭🇷 ",
		"Cuba":                           "🇨🇺 ",
		"Curaçao":                        "🇨🇼 ",
		"Cyprus":                         "🇨🇾 ",
		"Czechia":                        "🇨🇿 ",
		"Denmark":                        "🇩🇰 ",
		"Diego Garcia":                   "🇩🇬 ",
		"Djibouti":                       "🇩🇯 ",
		"Dominica":                       "🇩🇲 ",
		"Dominican Republic":             "🇩🇴 ",
		"Ecuador":                        "🇪🇨 ",
		"Egypt":                          "🇪🇬 ",
		"El Salvador":                    "🇸🇻 ",
		"England":                        "🏴󠁧󠁢󠁥󠁮󠁧󠁿",
		"Equatorial Guinea":              "🇬🇶 ",
		"Eritrea":                        "🇪🇷 ",
		"Estonia":                        "🇪🇪 ",
		"Ethiopia":                       "🇪🇹 ",
		"European Union":                 "🇪🇺 ",
		"Falkland Islands":               "🇫🇰 ",
		"Faroe Islands":                  "🇫🇴 ",
		"Fiji":                           "🇫🇯 ",
		"Finland":                        "🇫🇮 ",
		"France":                         "🇫🇷 ",
		"French Guiana":                  "🇬🇫 ",
		"French Polynesia":               "🇵🇫 ",
		"French Southern Territories":    "🇹🇫 ",
		"Gabon":                                  "🇬🇦 ",
		"Gambia":                                 "🇬🇲 ",
		"Georgia":                                "🇬🇪 ",
		"Germany":                                "🇩🇪 ",
		"Ghana":                                  "🇬🇭 ",
		"Gibraltar":                              "🇬🇮 ",
		"Greece":                                 "🇬🇷 ",
		"Greenland":                              "🇬🇱 ",
		"Grenada":                                "🇬🇩 ",
		"Guadeloupe":                             "🇬🇵 ",
		"Guam":                                   "🇬🇺 ",
		"Guatemala":                              "🇬🇹 ",
		"Guernsey":                               "🇬🇬 ",
		"Guinea-Bissau":                          "🇬🇼 ",
		"Guinea":                                 "🇬🇳 ",
		"Guyana":                                 "🇬🇾 ",
		"Haiti":                                  "🇭🇹 ",
		"Heard & McDonald Islands":               "🇭🇲 ",
		"Honduras":                               "🇭🇳 ",
		"Hong Kong SAR China":                    "🇭🇰 ",
		"Hungary":                                "🇭🇺 ",
		"Iceland":                                "🇮🇸 ",
		"India":                                  "🇮🇳 ",
		"Indonesia":                              "🇮🇩 ",
		"Iran":                                   "🇮🇷 ",
		"Iraq":                                   "🇮🇶 ",
		"Ireland":                                "🇮🇪 ",
		"Isle of Man":                            "🇮🇲 ",
		"Israel":                                 "🇮🇱 ",
		"Italy":                                  "🇮🇹 ",
		"Jamaica":                                "🇯🇲 ",
		"Japan":                                  "🇯🇵 ",
		"Jersey":                                 "🇯🇪 ",
		"Jordan":                                 "🇯🇴 ",
		"Kazakhstan":                             "🇰🇿 ",
		"Kenya":                                  "🇰🇪 ",
		"Kiribati":                               "🇰🇮 ",
		"Kosovo":                                 "🇽🇰 ",
		"Kuwait":                                 "🇰🇼 ",
		"Kyrgyzstan":                             "🇰🇬 ",
		"Laos":                                   "🇱🇦 ",
		"Latvia":                                 "🇱🇻 ",
		"Lebanon":                                "🇱🇧 ",
		"Lesotho":                                "🇱🇸 ",
		"Liberia":                                "🇱🇷 ",
		"Libya":                                  "🇱🇾 ",
		"Liechtenstein":                          "🇱🇮 ",
		"Lithuania":                              "🇱🇹 ",
		"Luxembourg":                             "🇱🇺 ",
		"Macau SAR China":                        "🇲🇴 ",
		"Macedonia":                              "🇲🇰 ",
		"Madagascar":                             "🇲🇬 ",
		"Malawi":                                 "🇲🇼 ",
		"Malaysia":                               "🇲🇾 ",
		"Maldives":                               "🇲🇻 ",
		"Mali":                                   "🇲🇱 ",
		"Malta":                                  "🇲🇹 ",
		"Marshall Islands":                       "🇲🇭 ",
		"Martinique":                             "🇲🇶 ",
		"Mauritania":                             "🇲🇷 ",
		"Mauritius":                              "🇲🇺 ",
		"Mayotte":                                "🇾🇹 ",
		"Mexico":                                 "🇲🇽 ",
		"Micronesia":                             "🇫🇲 ",
		"Moldova":                                "🇲🇩 ",
		"Monaco":                                 "🇲🇨 ",
		"Mongolia":                               "🇲🇳 ",
		"Montenegro":                             "🇲🇪 ",
		"Montserrat":                             "🇲🇸 ",
		"Morocco":                                "🇲🇦 ",
		"Mozambique":                             "🇲🇿 ",
		"Myanmar (Burma)":                        "🇲🇲 ",
		"Namibia":                                "🇳🇦 ",
		"Nauru":                                  "🇳🇷 ",
		"Nepal":                                  "🇳🇵 ",
		"Netherlands":                            "🇳🇱 ",
		"New Caledonia":                          "🇳🇨 ",
		"New Zealand":                            "🇳🇿 ",
		"Nicaragua":                              "🇳🇮 ",
		"Niger":                                  "🇳🇪 ",
		"Nigeria":                                "🇳🇬 ",
		"Niue":                                   "🇳🇺 ",
		"Norfolk Island":                         "🇳🇫 ",
		"North Korea":                            "🇰🇵 ",
		"Northern Mariana Islands":               "🇲🇵 ",
		"Norway":                                 "🇳🇴 ",
		"Oman":                                   "🇴🇲 ",
		"Pakistan":                               "🇵🇰 ",
		"Palau":                                  "🇵🇼 ",
		"Palestinian Territories":                "🇵🇸 ",
		"Panama":                                 "🇵🇦 ",
		"Papua New Guinea":                       "🇵🇬 ",
		"Paraguay":                               "🇵🇾 ",
		"Peru":                                   "🇵🇪 ",
		"Philippines":                            "🇵🇭 ",
		"Pitcairn Islands":                       "🇵🇳 ",
		"Poland":                                 "🇵🇱 ",
		"Portugal":                               "🇵🇹 ",
		"Puerto Rico":                            "🇵🇷 ",
		"Qatar":                                  "🇶🇦 ",
		"Réunion":                                "🇷🇪 ",
		"Romania":                                "🇷🇴 ",
		"Russia":                                 "🇷🇺 ",
		"Rwanda":                                 "🇷🇼 ",
		"Samoa":                                  "🇼🇸 ",
		"San Marino":                             "🇸🇲 ",
		"São Tomé & Príncipe":                    "🇸🇹 ",
		"Saudi Arabia":                           "🇸🇦 ",
		"Scotland":                               "🏴󠁧󠁢󠁳󠁣󠁴󠁿",
		"Senegal":                                "🇸🇳 ",
		"Serbia":                                 "🇷🇸 ",
		"Seychelles":                             "🇸🇨 ",
		"Sierra Leone":                           "🇸🇱 ",
		"Singapore":                              "🇸🇬 ",
		"Sint Maarten":                           "🇸🇽 ",
		"Slovakia":                               "🇸🇰 ",
		"Slovenia":                               "🇸🇮 ",
		"Solomon Islands":                        "🇸🇧 ",
		"Somalia":                                "🇸🇴 ",
		"South Africa":                           "🇿🇦 ",
		"South Georgia & South Sandwich Islands": "🇬🇸 ",
		"South Korea":                            "🇰🇷 ",
		"South Sudan":                            "🇸🇸 ",
		"Spain":                                  "🇪🇸 ",
		"Sri Lanka":                              "🇱🇰 ",
		"St. Barthélemy":                         "🇧🇱 ",
		"St. Helena":                             "🇸🇭 ",
		"St. Kitts & Nevis":                      "🇰🇳 ",
		"St. Lucia":                              "🇱🇨 ",
		"St. Martin":                             "🇲🇫 ",
		"St. Pierre & Miquelon":                  "🇵🇲 ",
		"St. Vincent & Grenadines":               "🇻🇨 ",
		"Sudan":                                  "🇸🇩 ",
		"Suriname":                               "🇸🇷 ",
		"Svalbard & Jan Mayen":                   "🇸🇯 ",
		"Swaziland":                              "🇸🇿 ",
		"Sweden":                                 "🇸🇪 ",
		"Switzerland":                            "🇨🇭 ",
		"Syria":                                  "🇸🇾 ",
		"Taiwan":                                 "🇹🇼 ",
		"Tajikistan":                             "🇹🇯 ",
		"Tanzania":                               "🇹🇿 ",
		"Thailand":                               "🇹🇭 ",
		"Timor-Leste":                            "🇹🇱 ",
		"Togo":                                   "🇹🇬 ",
		"Tokelau":                                "🇹🇰 ",
		"Tonga":                                  "🇹🇴 ",
		"Trinidad & Tobago":                      "🇹🇹 ",
		"Tristan Da Cunha":                       "🇹🇦 ",
		"Tunisia":                                "🇹🇳 ",
		"Turkey":                                 "🇹🇷 ",
		"Turkmenistan":                           "🇹🇲 ",
		"Turks & Caicos Islands":                 "🇹🇨 ",
		"Tuvalu":                                 "🇹🇻 ",
		"Uganda":                                 "🇺🇬 ",
		"Ukraine":                                "🇺🇦 ",
		"United Arab Emirates":                   "🇦🇪 ",
		"United Kingdom":                         "🇬🇧 ",
		"United States":                          "🇺🇸 ",
		"Uruguay":                                "🇺🇾 ",
		"U.S. Outlying Islands":                  "🇺🇲 ",
		"U.S. Virgin Islands":                    "🇻🇮 ",
		"UTC":             "☁️ ",
		"Uzbekistan":      "🇺🇿 ",
		"Vanuatu":         "🇻🇺 ",
		"Vatican City":    "🇻🇦 ",
		"Venezuela":       "🇻🇪 ",
		"Vietnam":         "🇻🇳 ",
		"Wales":           "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
		"Wallis & Futuna": "🇼🇫 ",
		"Western Sahara":  "🇪🇭 ",
		"Yemen":           "🇾🇪 ",
		"Zambia":          "🇿🇲 ",
		"Zimbabwe":        "🇿🇼 ",
		"Chequered":       "🏁",
		"Triangular":      "🚩",
		"Crossed":         "🎌",
		"Black":           "🏴",
		"White":           "🏳 ",
		"Rainbow":         "🏳️‍🌈 ",
	}
)
