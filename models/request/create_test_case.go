package request

import (
	"github.com/ninteen19/testlink-go-api/models"
)

type CreateTestCase struct {
	TestCaseName  string
	TestSuiteId   int
	TestProjectId int
	AuthorLogin   string
	Summary       string
	Steps         []models.TestCaseStep
	//Preconditions string
	//Status                 enums.TestCaseStatus
	//Importance             enums.TestImportance
	//Execution              enums.ExecutionType
	//Order      int
	InternalId int
	//CheckDuplicatedName    bool
	//ActionOnDuplicatedName enums.ActionOnDuplicate
}
