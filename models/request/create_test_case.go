package request

import (
	"github.com/ninteen19/testlink-go-api"
)

type CreateTestCase struct {
	TestCaseName    string
	TestProjectName string
	TestSuiteId     int
	TestProjectId   int
	AuthorLogin     string
	Summary         string
	Steps           []testlink.TestCaseStep
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

	//Validate Name, ProjectName, SuiteId is filled
	if len(m.TestCaseName) == 0 || len(m.TestProjectName) == 0 || m.TestSuiteId == 0 {
		return false
	}

	//Validate steps exist
	return len(m.Steps) > 0
}
