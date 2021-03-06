# Confluence to Testlink

Will read confluence content, Parse all the testcases and call create test case API to testlink.

## Installation

      go install github.com/ninteen19/confluence-to-test-link@latest

## Quick Start

1. Add env variable:
    - `TESTLINK_DEVKEY` = dev key / api key `*required`
    - `TESTLINK_LOGIN` = username, the one we use for login `*required`
    - `TESTLINK_URL` = override default testlink url `*optional`
2. Find confluence's page ID:
    - open the page, open dev tools and refresh
    - filter `pageId`
        - should be able to find something like
          this: https://confluence.gdn-app.com/rest/highlighting/1.0/panel-items?pageId=162592801&_=1625565605649
        - in this case the page id is `162592801`
    - get the confluence body content using this url (replace `{pageId}` from previous
      step) `https://confluence.gdn-app.com/rest/api/content/{pageId}?expand=body.storage`
    - copy all the content after it's loaded (ctrl A, ctrl C)
3. open terminal
4. type `confluence-to-test-link`
5. If you found any error, please contact me

## Run test

      go test ./...

### Check coverage

      go clean -testcache && go test -coverpkg ./... -coverprofile cover.out ./... && go tool cover -html=cover.out -o cover.html

then open the cover.html in browser

## Resources

- `example.json` confluence content json example
- `request.xml` xmlrpc requestbody generated for testlink

## TODO

- [x] unit tests
- [x] update testlink-go-api, read testProjectId by testprojectname
- connect to confluence API (unable to get the API key for now)
- update confluence structure
