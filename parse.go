package parser_html_page

import (
	"parser_html_page/models"
)

func GetResultParseHtml(parse models.Parse) (*[]models.ParseSelectionResult, error) {
	response, err := GetResponse(parse.Url, parse.HeaderSets)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resultParse, err := ParseHtml(response, parse.Selection)
	if err != nil {
		return nil, err
	}

	return resultParse, nil
}
