package utils

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/atotto/clipboard"
	"github.com/ninteen19/confluence-to-test-link/enums"
	"github.com/ninteen19/confluence-to-test-link/models/request"
	"github.com/ninteen19/confluence-to-test-link/models/response"
	"github.com/ninteen19/testlink-go-api"
	"strconv"
	"strings"
)

func ConvertClipboardToConfluenceContent() *response.ConfluenceContent {
	confluenceContentTxts, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	confluenceContent := &response.ConfluenceContent{}

	err = json.Unmarshal([]byte(confluenceContentTxts), confluenceContent)
	if err != nil {
		fmt.Println("Error unmarshalling JSON, cause:", err)
		return nil
	}

	return confluenceContent
}

func ConvertConfluenceContentToCreateTestCase(content *response.ConfluenceContent) *request.CreateTestCase {
	contentVal := *content
	str := contentVal.Body.Storage.Value
	str = strings.ReplaceAll(str, "&nbsp;", " ")
	str = strings.ReplaceAll(str, "&rarr;", " ")

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(str))
	if err != nil {
		fmt.Println("Fail read document from reader, cause:", err)
		return nil
	}

	createTestCase := &request.CreateTestCase{
		Status:              testlink.TestCaseStatusDraft,
		Importance:          testlink.TestImportanceMedium,
		Order:               0,
		Execution:           testlink.ExecutionTypeManual,
		CheckDuplicatedName: false,
	}

	doc.Find("table").Each(func(i int, selection *goquery.Selection) {
		selection.Find("thead th").Each(func(i int, selection *goquery.Selection) {
			switch i {
			case enums.Title.Value():
				createTestCase.TestCaseName = selection.Text()
				createTestCase.Summary = selection.Text()
			case enums.TestProjectId.Value():
				createTestCase.TestProjectId, _ = strconv.Atoi(selection.Text())
			case enums.TestSuiteId.Value():
				createTestCase.TestSuiteId, _ = strconv.Atoi(selection.Text())
			}
		})

		selection.Find("tbody table").Each(func(i int, selection *goquery.Selection) {
			selection.Find("tbody td:nth-child(2)").Each(func(i int, selection *goquery.Selection) {
				s, _ := selection.Find("ul").Html()
				step := selection.Find("p").Text() + `</br><ul>` + s + `</ul>`

				createTestCase.Steps = append(createTestCase.Steps, *createTestCaseStep(i, step))
			})
		})
	})

	return createTestCase
}

func createTestCaseStep(number int, step string) *testlink.TestCaseStep {
	manual := testlink.ExecutionTypeManual
	return &testlink.TestCaseStep{
		Actions:         step,
		Number:          number,
		ExpectedResults: "",
		ExecutionType:   manual.Value(),
	}
}
