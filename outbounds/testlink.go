package outbounds

import (
	"github.com/ninteen19/confluence-to-test-link/models/request"
	"github.com/ninteen19/testlink-go-api/enums"
	"github.com/ninteen19/testlink-go-api/models"
	"github.com/ninteen19/testlink-go-api/testcase"
)

func CreateTestCase(request *request.CreateTestCase) (*models.TestCase, error) {
	return testcase.Create(
		request.TestCaseName,
		request.TestSuiteId,
		request.TestProjectId,
		request.AuthorLogin,
		request.Summary,
		request.Steps,
		"",
		enums.DRAFT,
		enums.MEDIUM,
		enums.MANUAL,
		0,
		request.InternalId,
		false,
		0,
	)
}
