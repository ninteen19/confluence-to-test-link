package services

import (
	"github.com/ninteen19/confluence-to-test-link/models/request"
	"github.com/ninteen19/confluence-to-test-link/utils"
	"github.com/ninteen19/testlink-go-api/testproject"
	"log"
	"strconv"
	"sync"
)

type ITestlinkOutbound interface {
	CreateTestCase(request *request.CreateTestCase) error
	GetTestProject(testProjectName string) (*testproject.Response, error)
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

	s.updateTestCaseId(createTestCaseRequests)

	var wg sync.WaitGroup
	wg.Add(len(createTestCaseRequests))
	for _, testCaseRequest := range createTestCaseRequests {
		testCaseRequest := testCaseRequest
		go func() {
			defer wg.Done()
			err := s.ITestlinkOutbound.CreateTestCase(testCaseRequest)
			if err != nil {
				log.Println("Failed create test case, cause:", err)
				return
			}
		}()
	}
	wg.Wait()

	log.Printf("Success create %v testcase(s)", len(createTestCaseRequests))
}

func (s *TestlinkService) updateTestCaseId(createTestCaseRequests []*request.CreateTestCase) {
	groupByProjectName := map[string][]*request.CreateTestCase{}

	for _, createTestCaseRequest := range createTestCaseRequests {
		groupByProjectName[createTestCaseRequest.TestProjectName] = append(groupByProjectName[createTestCaseRequest.TestProjectName], createTestCaseRequest)
	}

	var wg sync.WaitGroup
	wg.Add(len(groupByProjectName))
	for key, val := range groupByProjectName {
		projectName := key
		groupedRequests := val
		go func() {
			defer wg.Done()
			testProject, err := s.ITestlinkOutbound.GetTestProject(projectName)
			if err != nil {
				log.Printf("Failed to find test project for projectName: %s, cause: %s\n", projectName, err)
			}

			for _, createTestCaseRequest := range groupedRequests {
				createTestCaseRequest.TestProjectId, _ = strconv.Atoi(testProject.Id)
			}
		}()
	}
	wg.Wait()

}
