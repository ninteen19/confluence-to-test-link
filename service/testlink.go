package service

import (
	"fmt"
	"github.com/ninteen19/confluence-to-test-link/models/request"
	"github.com/ninteen19/confluence-to-test-link/utils"
	"github.com/ninteen19/testlink-go-api"
)

type ITestlinkOutbound interface {
	CreateTestCase(request *request.CreateTestCase) (*testlink.TestCase, error)
}

type TestlinkService struct {
	ITestlinkOutbound ITestlinkOutbound
}

func NewTestlinkService(outbound ITestlinkOutbound) *TestlinkService {
	return &TestlinkService{
		ITestlinkOutbound: outbound,
	}
}

func (s *TestlinkService) CreateTestCaseFromConfluenceContentClipboard() {
	confluenceContent := utils.ConvertClipboardToConfluenceContent()
	createTestCaseRequest := utils.ConvertConfluenceContentToCreateTestCase(confluenceContent)

	fmt.Println(createTestCaseRequest)
}
