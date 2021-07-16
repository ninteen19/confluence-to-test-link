package outbounds

import (
	"github.com/ninteen19/confluence-to-test-link/models/request"
	"github.com/ninteen19/testlink-go-api"
	"github.com/ninteen19/testlink-go-api/testcase"
)

type TestlinkOutbound struct {
}

func NewTestlinkOutbound() *TestlinkOutbound {
	return &TestlinkOutbound{}
}

func (o *TestlinkOutbound) CreateTestCase(request *request.CreateTestCase) error {
	return testcase.Create(
		&testcase.CreateRequest{
			TestCaseName:  request.TestCaseName,
			TestSuiteId:   request.TestSuiteId,
			TestProjectId: request.TestProjectId,
			AuthorLogin:   request.AuthorLogin,
			Summary:       request.Summary,
			Steps:         request.Steps,
			Status:        testlink.TestCaseStatusDraft,
			Importance:    testlink.TestImportanceMedium,
			Execution:     testlink.ExecutionTypeManual,
			InternalId:    request.InternalId,
		},
	)
}
