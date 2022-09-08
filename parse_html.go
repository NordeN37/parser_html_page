package parser_html_page

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"parser_html_page/models"
)

func ParseHtml(res *http.Response, selection models.Selection) (*[]models.ParseSelectionResult, error) {
	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return parseSelection(doc.Selection, selection.Find), nil
}

func parseSelection(doc *goquery.Selection, selection []*models.Find) *[]models.ParseSelectionResult {
	var parsSelectResult []models.ParseSelectionResult

	for _, startFindValue := range selection {
		var find string
		if startFindValue.Tag != nil {
			find = *startFindValue.Tag
		}
		if startFindValue.Class != nil {
			find += " " + *startFindValue.Class
		}
		if startFindValue.Id != nil {
			find += " " + *startFindValue.Id
		}

		// Find the review items
		doc.Find(find).Each(func(i int, s *goquery.Selection) {
			var lineParse = models.ParseSelectionResult{}

			if startFindValue.GetValue {
				sText := s.Text()
				lineParse.Value = &sText
			}

			if startFindValue.Find != nil {
				lineParse.FoundValue = parseSelection(s, startFindValue.Find)
			}

			parsSelectResult = append(parsSelectResult, lineParse)
		})
	}

	return &parsSelectResult
}
