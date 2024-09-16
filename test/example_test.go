package test

import (
	"bytes"
	"github.com/NordeN37/parser_html_page"
	models "github.com/NordeN37/parser_html_page/models"
	"io/ioutil"
	"reflect"
	"testing"
)

var (
	tagDiv = "div"
	tagH   = "h1"
	tagP   = "p"

	resultH    = "Example Domain"
	resultOneP = "This domain is for use in illustrative examples in documents. You may use this\n    domain in literature without prior coordination or asking for permission."
	resultTwoP = "More information..."
)

type addTest struct {
	parse  models.Parse
	result *[]models.ParseSelectionResult
}

var addTests = []addTest{
	{
		parse: models.Parse{
			Selection: models.Selection{
				Find: []*models.Find{
					{
						Tag:      &tagDiv,
						GetValue: false,
						Find: []*models.Find{
							{
								Tag:      &tagH,
								GetValue: true,
							},
							{
								Tag:      &tagP,
								GetValue: true,
							},
						},
					},
				},
			},
		},
		result: &[]models.ParseSelectionResult{
			{
				Value: nil,
				FoundValue: []models.ParseSelectionResult{
					{Value: &resultH, FoundValue: nil},
					{Value: &resultOneP, FoundValue: nil},
					{Value: &resultTwoP, FoundValue: nil},
					{},
				},
			},
		},
	},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		reader := bytes.NewReader([]byte(testHtml))
		readerCloser := ioutil.NopCloser(reader)
		output, err := parser_html_page.ParseHtml(readerCloser, test.parse.Selection)
		if err != nil {
			t.Errorf("[ERROR]: %s", err.Error())
			continue
		}
		if reflect.DeepEqual(output, test.result) {
			t.Errorf("[ERROR]: output : %+v\\n not test.result : %+v\\n", output, test.result)
			continue
		}

		t.Log("[INFO]: result ok")

	}
}

var testHtml = `
<!doctype html>
<html>
<head>
    <title>Example Domain</title>

    <meta charset="utf-8" />
    <meta http-equiv="Content-type" content="text/html; charset=utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <style type="text/css">
    body {
        background-color: #f0f0f2;
        margin: 0;
        padding: 0;
        font-family: -apple-system, system-ui, BlinkMacSystemFont, "Segoe UI", "Open Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;
        
    }
    div {
        width: 600px;
        margin: 5em auto;
        padding: 2em;
        background-color: #fdfdff;
        border-radius: 0.5em;
        box-shadow: 2px 3px 7px 2px rgba(0,0,0,0.02);
    }
    a:link, a:visited {
        color: #38488f;
        text-decoration: none;
    }
    @media (max-width: 700px) {
        div {
            margin: 0 auto;
            width: auto;
        }
    }
    </style>    
</head>

<body>
<div>
    <h1>Example Domain</h1>
    <p>This domain is for use in illustrative examples in documents. You may use this
    domain in literature without prior coordination or asking for permission.</p>
    <p><a href="https://www.iana.org/domains/example">More information...</a></p>
</div>
</body>
</html>
`
