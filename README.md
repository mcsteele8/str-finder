# Estate Finder

Estate Finder is a small Go program used for scraping real estate websites for properties that match the configuration query params.

### Building & Running:
 - Configuration
   - config should be created in `configs/cfg.json`
   - config value definitions can be found in `configs/config.go`
   - Example Config:
```
{
  "alert_settings": [
    {
    "display_name": "Find me a STR",
    "cadence_in_minutes": 15,
    "send_alerts_to": [
      "mitch@gmail.com"
    ],
    "alert_query_params": {
      "cities": [
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
      "subdivisions": [
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
      "house_type": [
        "res",
        "con"
      ],
      "zoning": [
        "all"
      ],
      "list_price_min": "50000",
      "list_price_max": "all",
      "year_built_min": "all",
      "area_min": "all",
      "acres_min": "all",
      "beds_min": "all",
      "baths_min": "all",
      "levels_max": "all",
      "den": "all",
      "gated": "all",
      "pool": "all",
      "55_community": "all",
      "garage_spaces_min": "all",
      "vacation_rental": "all",
      "hoa": "all"
    }
  }],
  "email_settings":{
    "email":"dummy@gmail.com",
    "password":"itsMeMom!",
    "host":"smtp.gmail.com",
    "port":"587"
  }
}
```
 - Using Docker: 
   - Install docker: https://docs.docker.com/get-docker/
   - Build image `` docker build --tag docker-review-finder .``
   - Run image ``docker run docker-review-finder``
 - Using Go: 
   - Install Go: https://golang.org/doc/install
   - Terminal command: ``go run main.go``
### Testing
 - Using Go:
   - Terminal command: `./testing/test.sh`
