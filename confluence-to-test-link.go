package main

import (
	"github.com/ninteen19/confluence-to-test-link/outbounds"
	"github.com/ninteen19/confluence-to-test-link/services"
	"github.com/ninteen19/testlink-go-api"
	"log"
	"os"
)

func main() {
	devKey := os.Getenv("TESTLINK_DEVKEY")
	authorLogin := os.Getenv("TESTLINK_LOGIN")
	testlinkUrl := os.Getenv("TESTLINK_URL")

	if len(devKey) < 0 || len(authorLogin) < 0 {
		log.Println("Unable to find testlink-devkey OR testlink-login from env variable")
		return
	}

	testlink.Conf.Key = devKey
	testlink.Conf.Url = "https://testlink.gdn-app.com/lib/api/xmlrpc/v1/xmlrpc.php"

	if len(testlinkUrl) > 0 {
		testlink.Conf.Url = testlinkUrl
		log.Printf("Overriding testlink url from env variable: %v\n", testlinkUrl)
	}

	testlinkOutbound := outbounds.NewTestlinkOutbound()
	testlinkService := services.NewTestlinkService(testlinkOutbound)

	testlinkService.CreateTestCaseFromConfluenceContentClipboard(authorLogin)

}
