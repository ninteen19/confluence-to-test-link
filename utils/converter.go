package utils

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/atotto/clipboard"
	"github.com/ninteen19/confluence-to-test-link/enums"
	"github.com/ninteen19/confluence-to-test-link/models/request"
	"github.com/ninteen19/confluence-to-test-link/models/response"
	"github.com/ninteen19/testlink-go-api"
	"log"
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
		log.Println("Error unmarshalling JSON, cause:", err)
		return nil
	}

	return confluenceContent
}

func ConvertConfluenceContentToCreateTestCase(content *response.ConfluenceContent, authorLogin string) []*request.CreateTestCase {
	contentVal := *content
	str := contentVal.Body.Storage.Value
	str = strings.ReplaceAll(str, "&nbsp;", " ")
	str = strings.ReplaceAll(str, "&rarr;", " ")

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(str))
	if err != nil {
		log.Println("Fail read document from reader, cause:", err)
		return nil
	}

	var testCases []*request.CreateTestCase

	doc.Find("table").Each(func(i int, selection *goquery.Selection) {
		createTestCase := &request.CreateTestCase{
			Status:              testlink.TestCaseStatusDraft,
			Importance:          testlink.TestImportanceMedium,
			Order:               0,
			Execution:           testlink.ExecutionTypeManual,
			CheckDuplicatedName: false,
			AuthorLogin:         authorLogin,
		}
		selection.Find("thead th").Each(func(i int, selection *goquery.Selection) {
			switch i {
			case enums.Title.Value():
				createTestCase.TestCaseName = selection.Text()
				createTestCase.Summary = selection.Text()
			case enums.TestProjectName.Value():
				createTestCase.TestProjectName = selection.Text()
			case enums.TestSuiteId.Value():
				createTestCase.TestSuiteId, _ = strconv.Atoi(selection.Text())
			}
		})

		selection.Find("tbody table").Each(func(i int, selection *goquery.Selection) {
			selection.Find("tbody").Contents().Each(func(i int, selection *goquery.Selection) {
				no, err := strconv.Atoi(selection.Find("td:nth-child(1)").Text())

				if err != nil {
					return
				}

				step, _ := selection.Find("td:nth-child(2)").Html()
				expectedResults, _ := selection.Find("td:nth-child(3)").Html()

				createTestCase.Steps = append(createTestCase.Steps, *createTestCaseStep(no, step, expectedResults))
			})
		})

		if createTestCase.IsValid() {
			testCases = append(testCases, createTestCase)
		}
	})

	return testCases
}

func createTestCaseStep(number int, step string, expectedResults string) *testlink.TestCaseStep {
	manual := testlink.ExecutionTypeManual
	return &testlink.TestCaseStep{
		Actions:         step,
		Number:          number,
		ExpectedResults: expectedResults,
		ExecutionType:   manual.Value(),
	}
}
