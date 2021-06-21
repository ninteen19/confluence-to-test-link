package main

import (
	"github.com/ninteen19/confluence-to-test-link/outbounds"
	"github.com/ninteen19/confluence-to-test-link/service"
)

func main() {

	testlinkOutbound := outbounds.NewTestlinkOutbound()
	testlinkService := service.NewTestlinkService(testlinkOutbound)

	testlinkService.CreateTestCaseFromConfluenceContentClipboard()
	//testlink.Conf.Key = ""
	//testlink.Conf.Url = ""
}
