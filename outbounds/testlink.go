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
		request.TestCaseName,
		request.TestSuiteId,
		request.TestProjectId,
		request.AuthorLogin,
		request.Summary,
		request.Steps,
		"",
		testlink.TestCaseStatusDraft,
		testlink.TestImportanceMedium,
		testlink.ExecutionTypeManual,
		0,
		request.InternalId,
		false,
		0,
	)
}
