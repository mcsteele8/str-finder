package urls

import (
	"encoding/json"
	"fmt"
	"net/url"
	"projects/review-finder/tools/web"
	"strings"
)

type UrlParams struct {
	protocol        string
	hostName        string
	Cities          []string `json:"cities"`
	Subdivisions    []string `json:"subdivisions"`
	HouseType       []string `json:"house_type"`
	Zoning          []string `json:"zoning"`
	ListPriceMin    string   `json:"list_price_min"`
	ListPriceMax    string   `json:"list_price_max"`
	YearBuiltMin    string   `json:"year_built_min"`
	AreaMin         string   `json:"area_min"`
	AcresMin        string   `json:"acres_min"`
	BedsMin         string   `json:"beds_min"`
	BathsMin        string   `json:"baths_min"`
	LevelsMax       string   `json:"levels_max"`
	Den             string   `json:"den"`
	Gated           string   `json:"gated"`
	Pool            string   `json:"pool"`
	OldPeople       string   `json:"55_community"`
	GarageSpacesMin string   `json:"garage_spaces_min"`
	VacationRental  string   `json:"vacation_rental"`
	Hoa             string   `json:"hoa"`
	queryBuilder    strings.Builder
}

func NewUrl(b []byte) (string, error) {
	url := &UrlParams{}

	err := json.Unmarshal(b, url)
	if err != nil {
		return "", fmt.Errorf("error getting url params | %s", err.Error())
	}

	return url.formatUrl(), nil
}

func (u *UrlParams) SendRequest() ([]byte, error) {
	return web.FetchUrlData(u.formatUrl())
}

func (u *UrlParams) formatUrl() string {
	formatted := url.URL{
		Scheme:   "https",
		Host:     "www.stgeorgeutrealestate.com",
		Path:     "search/results/13ha/",
		RawQuery: u.getQueryString(),
	}
	return formatted.String()
}

func (u *UrlParams) buildUrlQuery() {
	u.addCityQueryParams()
	u.addSubdivisionsQueryParams()
	u.addHouseTypeQueryParams()
	u.addListPriceMinQueryParams()
	u.addListPriceMaxQueryParams()
	u.addYearBuiltMinQueryParams()
	u.addAreaMinQueryParams()
	u.addAcresMinQueryParams()
	u.addBedsMinQueryParams()
	u.addBathsMinQueryParams()
	u.addLevelsMaxQueryParams()
	u.addDenQueryParams()
	u.addGatedQueryParams()
	u.addPoolQueryParams()
	u.addOldPeopleQueryParams()
	u.addGarageSpaceMinQueryParams()
	u.addZoningQueryParams()
	u.addVacationRentalQueryParams()
	u.addHoaQueryParams()
}

func (u *UrlParams) getQueryString() string {
	u.buildUrlQuery()
	return u.queryBuilder.String()
}

func (u *UrlParams) addCityQueryParams() {
	query := strings.Builder{}
	for i, city := range u.Cities {
		query.WriteString("city=" + city)
		if len(u.Cities)-1 > i {
			query.WriteString("&")
		}
	}
	u.queryBuilder.WriteString(query.String())
}

func (u *UrlParams) addSubdivisionsQueryParams() {
	query := strings.Builder{}
	query.WriteString("&")
	for i, subdivision := range u.Subdivisions {
		query.WriteString("subdivision=" + subdivision)
		if len(u.Subdivisions)-1 > i {
			query.WriteString("&")
		}
	}
	u.queryBuilder.WriteString(query.String())
}

func (u *UrlParams) addHouseTypeQueryParams() {
	query := strings.Builder{}
	query.WriteString("&")
	for i, hType := range u.HouseType {
		query.WriteString("type=" + hType)
		if len(u.HouseType)-1 > i {
			query.WriteString("&")
		}
	}
	u.queryBuilder.WriteString(query.String())
}

func (u *UrlParams) addZoningQueryParams() {
	query := strings.Builder{}
	query.WriteString("&")
	for i, ZoneType := range u.Zoning {
		query.WriteString("zoning=" + ZoneType)
		if len(u.Zoning)-1 > i {
			query.WriteString("&")
		}
	}
	u.queryBuilder.WriteString(query.String())
}

func (u *UrlParams) addListPriceMinQueryParams() {
	if u.ListPriceMin == "" {
		u.ListPriceMin = "all"
	}
	u.queryBuilder.WriteString("&list_price_min=" + u.ListPriceMin)
}

func (u *UrlParams) addListPriceMaxQueryParams() {
	if u.ListPriceMax == "" {
		u.ListPriceMax = "all"
	}
	u.queryBuilder.WriteString("&list_price_max=" + u.ListPriceMax)
}

func (u *UrlParams) addYearBuiltMinQueryParams() {
	if u.YearBuiltMin == "" {
		u.YearBuiltMin = "all"
	}
	u.queryBuilder.WriteString("&year_built_min=" + u.YearBuiltMin)
}

func (u *UrlParams) addAreaMinQueryParams() {
	if u.AreaMin == "" {
		u.AreaMin = "all"
	}
	u.queryBuilder.WriteString("&area_min=" + u.AreaMin)
}

func (u *UrlParams) addAcresMinQueryParams() {
	if u.AcresMin == "" {
		u.AcresMin = "all"
	}
	u.queryBuilder.WriteString("&acres_min=" + u.AcresMin)
}

func (u *UrlParams) addBedsMinQueryParams() {
	if u.BedsMin == "" {
		u.BedsMin = "all"
	}
	u.queryBuilder.WriteString("&beds_min=" + u.BedsMin)
}

func (u *UrlParams) addBathsMinQueryParams() {
	if u.BathsMin == "" {
		u.BathsMin = "all"
	}
	u.queryBuilder.WriteString("&baths_min=" + u.BathsMin)
}

func (u *UrlParams) addLevelsMaxQueryParams() {
	if u.LevelsMax == "" {
		u.LevelsMax = "all"
	}
	u.queryBuilder.WriteString("&levels_max=" + u.LevelsMax)
}

func (u *UrlParams) addDenQueryParams() {
	if u.Den == "" {
		u.Den = "all"
	}
	u.queryBuilder.WriteString("&den=" + u.Den)
}

func (u *UrlParams) addGatedQueryParams() {
	if u.Gated == "" {
		u.Gated = "all"
	}
	u.queryBuilder.WriteString("&gated=" + u.Gated)
}

func (u *UrlParams) addPoolQueryParams() {
	if u.Pool == "" {
		u.Pool = "all"
	}
	u.queryBuilder.WriteString("&pool=" + u.Pool)
}

func (u *UrlParams) addOldPeopleQueryParams() {
	if u.OldPeople == "" {
		u.OldPeople = "all"
	}
	u.queryBuilder.WriteString("&55_community=" + u.OldPeople)
}

func (u *UrlParams) addGarageSpaceMinQueryParams() {
	if u.GarageSpacesMin == "" {
		u.GarageSpacesMin = "all"
	}
	u.queryBuilder.WriteString("&garage_spaces_min=" + u.GarageSpacesMin)
}

func (u *UrlParams) addVacationRentalQueryParams() {
	if u.VacationRental == "" {
		u.VacationRental = "all"
	}
	u.queryBuilder.WriteString("&vacation_rental=" + u.VacationRental)
}

func (u *UrlParams) addHoaQueryParams() {
	if u.Hoa == "" {
		u.Hoa = "all"
	}
	u.queryBuilder.WriteString("&hoa=" + u.Hoa)
}
