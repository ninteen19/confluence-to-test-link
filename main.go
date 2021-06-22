package main

import (
	"github.com/ninteen19/confluence-to-test-link/outbounds"
	"github.com/ninteen19/confluence-to-test-link/service"
	"github.com/ninteen19/testlink-go-api"
)

func main() {
	testlink.Conf.Key = "ab3b331ab9a3996cad5fe61d629bfd31"
	testlink.Conf.Url = "https://testlink.gdn-app.com/lib/api/xmlrpc/v1/xmlrpc.php"

	testlinkOutbound := outbounds.NewTestlinkOutbound()
	testlinkService := service.NewTestlinkService(testlinkOutbound)

	testlinkService.CreateTestCaseFromConfluenceContentClipboard()

}
