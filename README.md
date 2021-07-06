# Confluence to Testlink

Will read confluence content (example: `...`), Insert all the testcases to testlink, as per testProjectId and
testSuiteId defined

## Installation

1. git clone https://github.com/ninteen19/confluence-to-test-link.git
2. cd confluence-to-test-link
3. go install

## Quick Start

1. Add env variable:
    - `testlink-devkey` = dev key / api key `*required`
    - `testlink-login` = username, the one we use for login `*required`
    - `testlink-url` = override default testlink url `*optional`
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

## Resources

- `example.json` confluence content json example
- `request.xml` xmlrpc requestbody generated for testlink

## TODO

- unit tests
- update testlink-go-api, read testProjectId by testprojectkey
- connect to confluence API (unable to get the API key for now)
- update confluence structure
