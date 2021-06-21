package enums

import "strconv"

type TableHead int

const (
	Title TableHead = iota
	TestProjectId
	TestSuiteId
)

func (e TableHead) Value() int {
	return int(e)
}

func (e TableHead) StringValue() string {
	return strconv.Itoa(e.Value())
}
