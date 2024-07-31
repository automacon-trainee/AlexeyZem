package main

import (
	"regexp"
)

type Ad struct {
	Title       string
	Description string
}

func censorAds(ads []Ad, censor map[string]string) []Ad {
	for i, ad := range ads {
		for k, v := range censor {
			re := regexp.MustCompile("(?i)" + k)
			ad.Description = re.ReplaceAllString(ad.Description, v)
			ad.Title = re.ReplaceAllString(ad.Title, v)
			ads[i] = ad
		}
	}
	return ads
}

func main() {}
