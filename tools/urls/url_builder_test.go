package urls

import (
	"testing"
)

func Test_urlParams_formatUrl(t *testing.T) {
	type fields struct {
		Cities          []string
		Subdivisions    []string
		HouseType       []string
		Zoning          []string
		ListPriceMin    string
		ListPriceMax    string
		YearBuiltMin    string
		AreaMin         string
		AcresMin        string
		BedsMin         string
		BathsMin        string
		LevelsMax       string
		Den             string
		Gated           string
		Pool            string
		OldPeople       string
		GarageSpacesMin string
		VacationRental  string
		Hoa             string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			fields: fields{
				Cities:          []string{"Ivins", "Hurricane", "Pine+Valley", "St+George"},
				Subdivisions:    []string{"Cliff+View+Est+At+Copper+Rock", "Cross+Roads+At+Stucki+Farms", "Ledges+Of+St+George", "Shinava+Ridge", "Villas+At+Sand+Hollow"},
				HouseType:       []string{"res", "con"},
				Zoning:          []string{"all"},
				ListPriceMin:    "50000",
				ListPriceMax:    "",
				YearBuiltMin:    "",
				AreaMin:         "",
				AcresMin:        "",
				BedsMin:         "",
				BathsMin:        "",
				LevelsMax:       "",
				Den:             "",
				Gated:           "",
				Pool:            "",
				OldPeople:       "",
				GarageSpacesMin: "",
				VacationRental:  "",
				Hoa:             "",
			},
			want: "https://www.stgeorgeutrealestate.com/search/results/13ha/?city=Ivins&city=Hurricane&city=Pine+Valley&city=St+George&subdivision=Cliff+View+Est+At+Copper+Rock&subdivision=Cross+Roads+At+Stucki+Farms&subdivision=Ledges+Of+St+George&subdivision=Shinava+Ridge&subdivision=Villas+At+Sand+Hollow&type=res&type=con&list_price_min=50000&list_price_max=all&year_built_min=all&area_min=all&acres_min=all&beds_min=all&baths_min=all&levels_max=all&den=all&gated=all&pool=all&55_community=all&garage_spaces_min=all&zoning=all&vacation_rental=all&hoa=all",
		},
		{
			name: "success 2",
			fields: fields{
				Cities:          []string{"Dammeron+Valley", "Hurricane", "Ivins", "La+Verkin", "Leeds", "Pine+Valley", "Santa+Clara", "Springdale", "St+George", "Toquerville", "Washington"},
				Subdivisions:    []string{"Cliff+View+Est+At+Copper+Rock", "Cross+Roads+At+Stucki+Farms", "Encanto+Resort", "Escapes+At+The+Ledges", "Golf+View+Estates", "Inn+Of+Entrada", "Ledges+Of+St+George", "Lofts+At+Green+Valley", "Shinava+Ridge", "Sports+Village", "Villas+At+Sand+Hollow", "Vue+At+Green+Valley"},
				HouseType:       []string{"res", "con"},
				Zoning:          []string{"all"},
				ListPriceMin:    "50000",
				ListPriceMax:    "all",
				YearBuiltMin:    "all",
				AreaMin:         "all",
				AcresMin:        "all",
				BedsMin:         "all",
				BathsMin:        "all",
				LevelsMax:       "all",
				Den:             "all",
				Gated:           "all",
				Pool:            "all",
				OldPeople:       "all",
				GarageSpacesMin: "all",
				VacationRental:  "all",
				Hoa:             "all",
			},
			want: "https://www.stgeorgeutrealestate.com/search/results/13ha/?city=Dammeron+Valley&city=Hurricane&city=Ivins&city=La+Verkin&city=Leeds&city=Pine+Valley&city=Santa+Clara&city=Springdale&city=St+George&city=Toquerville&city=Washington&subdivision=Cliff+View+Est+At+Copper+Rock&subdivision=Cross+Roads+At+Stucki+Farms&subdivision=Encanto+Resort&subdivision=Escapes+At+The+Ledges&subdivision=Golf+View+Estates&subdivision=Inn+Of+Entrada&subdivision=Ledges+Of+St+George&subdivision=Lofts+At+Green+Valley&subdivision=Shinava+Ridge&subdivision=Sports+Village&subdivision=Villas+At+Sand+Hollow&subdivision=Vue+At+Green+Valley&type=res&type=con&list_price_min=50000&list_price_max=all&year_built_min=all&area_min=all&acres_min=all&beds_min=all&baths_min=all&levels_max=all&den=all&gated=all&pool=all&55_community=all&garage_spaces_min=all&zoning=all&vacation_rental=all&hoa=all",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UrlParams{
				Cities:          tt.fields.Cities,
				Subdivisions:    tt.fields.Subdivisions,
				HouseType:       tt.fields.HouseType,
				Zoning:          tt.fields.Zoning,
				ListPriceMin:    tt.fields.ListPriceMin,
				ListPriceMax:    tt.fields.ListPriceMax,
				YearBuiltMin:    tt.fields.YearBuiltMin,
				AreaMin:         tt.fields.AreaMin,
				AcresMin:        tt.fields.AcresMin,
				BedsMin:         tt.fields.BedsMin,
				BathsMin:        tt.fields.BathsMin,
				LevelsMax:       tt.fields.LevelsMax,
				Den:             tt.fields.Den,
				Gated:           tt.fields.Gated,
				Pool:            tt.fields.Pool,
				OldPeople:       tt.fields.OldPeople,
				GarageSpacesMin: tt.fields.GarageSpacesMin,
				VacationRental:  tt.fields.VacationRental,
				Hoa:             tt.fields.Hoa,
			}
			if got := u.formatUrl(); got != tt.want {
				t.Errorf("formatUrl() = \ngot:  %v, \nwant: %v", got, tt.want)
			}
		})
	}
}

var paulaRecommends = `{
   "cities":[
      "Dammeron+Valley",
      "Hurricane",
      "Ivins",
      "La+Verkin",
      "Leeds",
      "Pine+Valley",
      "Santa+Clara",
      "Springdale",
      "St+George",
      "Toquerville",
      "Washington"
   ],
   "subdivisions":[
      "Cliff+View+Est+At+Copper+Rock",
      "Cross+Roads+At+Stucki+Farms",
      "Encanto+Resort",
      "Escapes+At+The+Ledges",
      "Golf+View+Estates",
      "Inn+Of+Entrada",
      "Ledges+Of+St+George",
      "Lofts+At+Green+Valley",
      "Shinava+Ridge",
      "Sports+Village",
      "Villas+At+Sand+Hollow",
      "Vue+At+Green+Valley"
   ],
   "house_type":[
      "res",
      "con"
   ],
   "zoning":[
      "all"
   ],
   "list_price_min":"50000",
   "list_price_max":"all",
   "year_built_min":"all",
   "area_min":"all",
   "acres_min":"all",
   "beds_min":"all",
   "baths_min":"all",
   "levels_max":"all",
   "den":"all",
   "gated":"all",
   "pool":"all",
   "55_community":"all",
   "garage_spaces_min":"all",
   "vacation_rental":"all",
   "hoa":"all"
}`

var jsonBlob = `{
   "cities":[
      "Ivins",
      "Hurricane",
      "Pine+Valley",
      "St+George"
   ],
   "subdivisions":[
      "Cliff+View+Est+At+Copper+Rock",
      "Cross+Roads+At+Stucki+Farms",
      "Ledges+Of+St+George",
      "Shinava+Ridge",
      "Villas+At+Sand+Hollow"
   ],
   "house_type":[
      "res",
      "con"
   ],
   "zoning":[
      "all"
   ],
   "list_price_min":"50000",
   "list_price_max":"all",
   "year_built_min":"all",
   "area_min":"all",
   "acres_min":"all",
   "beds_min":"all",
   "baths_min":"all",
   "levels_max":"all",
   "den":"all",
   "gated":"all",
   "pool":"all",
   "55_community":"all",
   "garage_spaces_min":"all",
   "vacation_rental":"all",
   "hoa":"all"
}`

func TestNewUrl(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				b: []byte(paulaRecommends),
			},
			want:    "https://www.stgeorgeutrealestate.com/search/results/13ha/?city=Dammeron+Valley&city=Hurricane&city=Ivins&city=La+Verkin&city=Leeds&city=Pine+Valley&city=Santa+Clara&city=Springdale&city=St+George&city=Toquerville&city=Washington&subdivision=Cliff+View+Est+At+Copper+Rock&subdivision=Cross+Roads+At+Stucki+Farms&subdivision=Encanto+Resort&subdivision=Escapes+At+The+Ledges&subdivision=Golf+View+Estates&subdivision=Inn+Of+Entrada&subdivision=Ledges+Of+St+George&subdivision=Lofts+At+Green+Valley&subdivision=Shinava+Ridge&subdivision=Sports+Village&subdivision=Villas+At+Sand+Hollow&subdivision=Vue+At+Green+Valley&type=res&type=con&list_price_min=50000&list_price_max=all&year_built_min=all&area_min=all&acres_min=all&beds_min=all&baths_min=all&levels_max=all&den=all&gated=all&pool=all&55_community=all&garage_spaces_min=all&zoning=all&vacation_rental=all&hoa=all",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUrl(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
