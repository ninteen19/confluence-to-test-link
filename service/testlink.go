package service

import (
	"github.com/ninteen19/confluence-to-test-link/models/request"
	"github.com/ninteen19/confluence-to-test-link/utils"
	"log"
	"sync"
)

type ITestlinkOutbound interface {
	CreateTestCase(request *request.CreateTestCase) error
}

type TestlinkService struct {
	ITestlinkOutbound ITestlinkOutbound
}

func NewTestlinkService(outbound ITestlinkOutbound) *TestlinkService {
	return &TestlinkService{
		ITestlinkOutbound: outbound,
	}
}

func (s *TestlinkService) CreateTestCaseFromConfluenceContentClipboard(authorLogin string) {
	confluenceContent := utils.ConvertClipboardToConfluenceContent()

	if confluenceContent == nil {
		log.Println("Confluence content not valid!")
		return
	}

	createTestCaseRequests := utils.ConvertConfluenceContentToCreateTestCase(confluenceContent, authorLogin)

	var wg sync.WaitGroup

	wg.Add(len(createTestCaseRequests))
	for _, testCaseRequest := range createTestCaseRequests {
		testCaseRequest := testCaseRequest
		go func() {
			defer wg.Done()
			err := s.ITestlinkOutbound.CreateTestCase(testCaseRequest)
			if err != nil {
				log.Println("Failed creating test case, cause:", err)
				return
			}
		}()
	}
	wg.Wait()

	log.Printf("Success creating %v testcase(s)", len(createTestCaseRequests))
}
