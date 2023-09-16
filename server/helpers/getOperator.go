package helpers

import "strings"

var ListOperators = []map[string]string{
	{
		"key":  "tsel",
		"name": "telkomsel",
	},
	{
		"key":  "telkomsel",
		"name": "telkomsel",
	},

	{
		"key":  "isat",
		"name": "indosat",
	},
	{
		"key":  "indosat",
		"name": "indosat",
	},
	{
		"key":  "xl",
		"name": "xl",
	},
	{
		"key":  "three",
		"name": "three",
	},
	{
		"key":  "tri",
		"name": "three",
	},
	{
		"key":  "axis",
		"name": "axis",
	},
	{
		"key":  "aigo",
		"name": "axis",
	},
	{
		"key":  "by.u",
		"name": "by.u",
	},
	{
		"key":  "smartfren",
		"name": "smartfren",
	},
	{
		"key":  "smart",
		"name": "smartfren",
	},
	{
		"key":  "gopay",
		"name": "gopay",
	},
	{
		"key":  "ovo",
		"name": "ovo",
	},
	{
		"key":  "grab",
		"name": "ovo",
	},
	{
		"key":  "dana",
		"name": "dana",
	},
	{
		"key":  "link aja",
		"name": "linkaja",
	},
	{
		"key":  "shopeepay",
		"name": "shopeepay",
	},
	{
		"key":  "maxim",
		"name": "maxim",
	},
	{
		"key":  "pln",
		"name": "pln",
	},
	{
		"key":  "mobile legend",
		"name": "game-ml",
	},
	{
		"key":  "brizzi",
		"name": "brizzi",
	},
	{
		"key":  "mandiri",
		"name": "mandiri",
	},
	{
		"key":  "free fire",
		"name": "game-ff",
	},
	{
		"key":  "pubg",
		"name": "game-pubg",
	},
	{
		"key":  "tapcash",
		"name": "tapcash",
	},
	{
		"key":  "call of duty mobile",
		"name": "game-cod",
	},
}

func GetOperator(str string) string {
	var res string
	for _, operator := range ListOperators {
		if strings.Contains(strings.ToLower(str), operator["key"]) {
			res = operator["name"]
		}
	}
	return res
}
