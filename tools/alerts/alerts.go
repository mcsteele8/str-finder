package alerts

import (
	"fmt"
	"projects/review-finder/tools/email"
	"projects/review-finder/tools/parser"
	"projects/review-finder/tools/urls"
	"strings"
	"time"
)

type AlertSettings struct {
	DisplayName      string         `json:"display_name"`
	CadenceInMinutes int            `json:"cadence_in_minutes"`
	SendAlertsTo     []string       `json:"send_alerts_to"`
	AlertQueryParams urls.UrlParams `json:"alert_query_params"`
}

func (a *AlertSettings) CreateNewJob(emailService email.Email) {
	fmt.Printf("\n      Registering New Job: %s\n", a.DisplayName)
	a.EmailYouHaveBeenEnrolledTemplate(emailService)
	cadence := time.NewTicker(time.Duration(a.CadenceInMinutes) * time.Minute)
	for {
		select {
		case <-cadence.C:
			b, err := a.AlertQueryParams.SendRequest()
			if err != nil {
				fmt.Printf("failed to get data, err: " + err.Error())
			}

			properties := parser.GetPropertyResults(b)

			if len(properties) > 0 {
				err = emailService.SendEmail(a.SendAlertsTo, properties.EmailTemplateMessage())
				if err == nil {
					fmt.Printf("\nEmail sent for %s", a.DisplayName)
				} else {
					fmt.Printf("ERROR!!!! %v", err)
				}
			}
		}
	}
}

func (a *AlertSettings) EmailYouHaveBeenEnrolledTemplate(emailService email.Email) {
	subject := fmt.Sprintf("Subject: Welcome to STR Alerts\n")
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := fmt.Sprintf(`<html>
					<body>
						<div style="text-align: center; margin-left: 10%%; margin-right: 10%%;">
						<h1>Welcome!</h1>
						<p>You have been enrolled in a STR alert finder. The property bot scrapes realestate websites looking for qualified short term rental properties and will notify you based on the alert settings and search query listed below. If you would like to be removed from this alert please reply with "STOP" or contact Mitch</p>
						<div>
							<table style="margin: 0 auto; border: 1px solid black;">
								<tr>
									<th colspan="2">Alert Settings</th>
								</tr>
								<tr>
									<td style="border: 1px solid black;">Alert Name</td>
									<td style="border: 1px solid black;">%s</td>
								</tr>
								<tr>
									<td style="border: 1px solid black;">Alert Scrape Cadence</td>
									<td style="border: 1px solid black;">%d min</td>
								</tr>			
							</table>
						</div>
						<div>
							<table style="margin: 0 auto; border: 1px solid black;">
								<tr>
									<th colspan="2">Search Query</th>
								</tr>
								%s %s %s %s %s
							</table>
						</div>
						</div>
					</body>
					</html>`, a.DisplayName, a.CadenceInMinutes, a.getCitiesHtml(), a.getSubdivisionsHtml(), a.getHouseTypeHtml(), a.getZoningHtml(), a.getAllOtherQueryHtml())

	for _, email := range a.SendAlertsTo {
		err := emailService.SendEmail([]string{email}, []byte(subject+mime+message))
		if err != nil {
			fmt.Println("Error sending welcome email: err | " + err.Error())
		}
	}

}

func (a *AlertSettings) getCitiesHtml() string {
	s := strings.Builder{}

	for _, c := range a.AlertQueryParams.Cities {
		s.WriteString("<tr>")
		s.WriteString(`<td style="border: 1px solid black;">City</td>`)
		s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, strings.ReplaceAll(c, "+", " ")))
		s.WriteString("</tr>")

	}

	return s.String()
}

func (a *AlertSettings) getSubdivisionsHtml() string {
	s := strings.Builder{}

	for _, sub := range a.AlertQueryParams.Subdivisions {
		s.WriteString("<tr>")
		s.WriteString(`<td style="border: 1px solid black;">Subdivision</td>`)
		s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, strings.ReplaceAll(sub, "+", " ")))
		s.WriteString("</tr>")

	}

	return s.String()
}

func (a *AlertSettings) getHouseTypeHtml() string {
	s := strings.Builder{}

	for _, c := range a.AlertQueryParams.HouseType {
		s.WriteString("<tr>")
		s.WriteString(`<td style="border: 1px solid black;">House Type</td>`)
		s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, strings.ReplaceAll(c, "+", " ")))
		s.WriteString("</tr>")
	}

	return s.String()
}

func (a *AlertSettings) getZoningHtml() string {
	s := strings.Builder{}

	for _, c := range a.AlertQueryParams.Zoning {
		s.WriteString("<tr>")
		s.WriteString(`<td style="border: 1px solid black;">Zoning</td>`)
		s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, strings.ReplaceAll(c, "+", " ")))
		s.WriteString("</tr>")

	}

	return s.String()
}

func (a *AlertSettings) getAllOtherQueryHtml() string {
	s := strings.Builder{}

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">Min Price</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.ListPriceMin))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">Max Price</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.ListPriceMax))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">Year Built Min</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.YearBuiltMin))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">Area Min</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.AreaMin))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">Acres Min</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.AcresMin))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">Beds Min</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.BedsMin))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">Baths Min</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.BathsMin))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">Levels Max</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.LevelsMax))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">Den</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.Den))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">Gated</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.Gated))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">Pool</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.Pool))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">55+</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.OldPeople))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">Garage Spaces Min</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.GarageSpacesMin))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">Vacation Rental</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.VacationRental))
	s.WriteString("</tr>")

	s.WriteString("<tr>")
	s.WriteString(`<td style="border: 1px solid black;">HOA</td>`)
	s.WriteString(fmt.Sprintf(`<td style="border: 1px solid black;">%s</td>`, a.AlertQueryParams.Hoa))
	s.WriteString("</tr>")

	return s.String()
}
