package estate_listing

import (
	"fmt"
	"strings"
	"time"
)

type EstateListing struct {
	Url      string `json:"url"`
	ImageSrc string `json:"image_src"`
	Price    string `json:"price"`
	Address  string `json:"address"`
}

type EstateListingDetails []EstateListing

func (e EstateListingDetails) EmailTemplateMessage() []byte {
	t := time.Now()
	subject := fmt.Sprintf("Subject: STR Alert %d/%d\n", t.Month(), t.Day())
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	b := strings.Builder{}
	b.WriteString("<html><body><h1>New Short Term Rental Listings!</h1>")
	for _, p := range e {
		b.WriteString("<div>")
		b.WriteString(fmt.Sprintf("<a href=\"%s\">\n<img src=\"%s\">\n</a>", p.Url, p.ImageSrc))
		b.WriteString(fmt.Sprintf("<a href=\"%s\">\n Price: %s\n | Address: %s</a>", p.Url, p.Price, p.Address))
		b.WriteString("</div>")
	}
	b.WriteString("</body></html>")

	return []byte(subject + mime + b.String())
}
