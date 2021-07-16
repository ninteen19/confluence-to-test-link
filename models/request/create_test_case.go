package request

import (
	"github.com/ninteen19/testlink-go-api"
)

type CreateTestCase struct {
	TestCaseName  string
	TestSuiteId   int
	TestProjectId int
	AuthorLogin   string
	Summary       string
	Steps         []testlink.TestCaseStep
	//Preconditions string
	Status              testlink.TestCaseStatus
	Importance          testlink.TestImportance
	Execution           testlink.ExecutionType
	Order               int
	InternalId          int
	CheckDuplicatedName bool
	//ActionOnDuplicatedName enums.ActionOnDuplicate
}

func (m *CreateTestCase) IsValid() bool {

	//Validate Name, ProjectId, SuiteId is filled
	if len(m.TestCaseName) == 0 || m.TestProjectId == 0 || m.TestSuiteId == 0 {
		return false
	}

	//Validate steps exist
	return len(m.Steps) > 0
}
