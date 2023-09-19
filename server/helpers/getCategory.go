package helpers

import "strings"

var ListKeys = []map[string]string{
	{
		"key":  "data",
		"name": "data",
	},
	{
		"key":  "internet",
		"name": "data",
	},
	{
		"key":  "inject",
		"name": "data",
	}, {
		"key":  "kuota",
		"name": "data",
	}, {
		"key":  "byu",
		"name": "data",
	},

	{
		"key":  "reguler",
		"name": "pulsa",
	},
	{
		"key":  "denom unik",
		"name": "pulsa",
	},
	{
		"key":  "telkomsel khusus",
		"name": "pulsa",
	},
	{
		"key":  "indosat khusus",
		"name": "pulsa",
	},
	{
		"key":  "kikipu",
		"name": "pulsa",
	},
	{
		"key":  "pulsa",
		"name": "pulsa",
	}, {
		"key":  "tees",
		"name": "pulsa",
	},
	{
		"key":  "voucher",
		"name": "pulsa-voucher",
	},
	{
		"key":  "transfer",
		"name": "pulsa-transfer",
	}, {
		"key":  "telpon",
		"name": "pulsa",
	},
	{
		"key":  "masa aktif",
		"name": "pulsa",
	},
	{
		"key":  "pubg",
		"name": "game",
	},
	{
		"key":  "cod",
		"name": "game",
	},
	{
		"key":  "mobile legen",
		"name": "game",
	},
	{
		"key":  "free fire",
		"name": "game",
	},
	{
		"key":  "shopeepay",
		"name": "e-wallet",
	},
	{
		"key":  "ovo",
		"name": "e-wallet",
	},
	{
		"key":  "gojek",
		"name": "e-wallet",
	},
	{
		"key":  "grab",
		"name": "e-wallet",
	},
	{
		"key":  "dana",
		"name": "e-wallet",
	},
	{
		"key":  "link aja",
		"name": "e-wallet",
	},
	{
		"key":  "pln",
		"name": "pln",
	},
	{
		"key":  "maxim",
		"name": "e-wallet",
	}, {
		"key":  "e-money",
		"name": "e-money",
	},
}

func GetCategory(str1 string, str2 string) string {
	var res string
	for _, operator := range ListKeys {
		if strings.Contains(strings.ToLower(str1), operator["key"]) || strings.Contains(strings.ToLower(str2), operator["key"]) {
			res = operator["name"]
		}
	}
	return res
}
