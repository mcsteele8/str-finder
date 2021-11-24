package parser

import (
	"projects/review-finder/tools/estate_listing"
	"regexp"
	"strings"
)

func GetPropertyResults(b []byte) estate_listing.EstateListingDetails {
	match := regexp.MustCompile("property-thumb(.|\n)*?property-description")
	propertyBlobs := match.FindAllString(string(b), -1)
	var listings []estate_listing.EstateListing
	for _, blob := range propertyBlobs {
		listings = append(listings, getEstateDetails(blob))
	}
	return listings
}

func getEstateDetails(blob string) estate_listing.EstateListing {
	return estate_listing.EstateListing{
		Url:      getUrl(blob),
		ImageSrc: getImageSrc(blob),
		Price:    getPrice(blob),
		Address:  getAddress(blob),
	}
}

func getUrl(blob string) string {
	match := regexp.MustCompile("href=\"(.|\n)*?\"")
	url := trimString(match.FindString(blob), 6, 1)
	return "https://www.stgeorgeutrealestate.com" + url
}

func getPrice(blob string) string {
	return trimString(regexp.MustCompile("price\">(.|\n)*?<").FindString(blob), 7, 1)
}

func getImageSrc(blob string) string {
	return trimString(regexp.MustCompile("src=\"(.|\n)*?\"").FindString(blob), 5, 0)
}

func getAddress(blob string) string {
	return trimString(regexp.MustCompile("address\">(.|\n)*?</a>").FindString(blob), 9, 4)
}

func trimString(cleanMe string, trimLeftLength, trimRightLength int) string {
	trimLeft := cleanMe[trimLeftLength:]
	trimRight := trimLeft[:len(trimLeft)-trimRightLength]
	clean := strings.ReplaceAll(trimRight, "\n", "")
	clean = strings.ReplaceAll(clean, "\t", "")
	clean = strings.ReplaceAll(clean, "\"", "")
	clean = strings.TrimSpace(clean)
	return clean
}
