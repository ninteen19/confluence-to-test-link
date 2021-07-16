package services

import (
	"github.com/atotto/clipboard"
	"github.com/ninteen19/confluence-to-test-link/models/request"
	"github.com/ninteen19/testlink-go-api"
	"github.com/stretchr/testify/mock"
	"testing"
)

type TestlinkOutboundMock struct {
	mock.Mock
}

func (m *TestlinkOutboundMock) CreateTestCase(request *request.CreateTestCase) error {
	args := m.Called(request)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func TestCreateTestCaseFromConfluenceContentClipboard(t *testing.T) {
	outboundMock := new(TestlinkOutboundMock)
	testlinkService := NewTestlinkService(outboundMock)

	validRequests := []*request.CreateTestCase{
		{
			TestCaseName:  "BM-2638 - [Subscription Center] As a user, I want to bulk upload product data to Subscription Service",
			TestSuiteId:   37591638,
			TestProjectId: 37591635,
			AuthorLogin:   "andrian",
			Summary:       "BM-2638 - [Subscription Center] As a user, I want to bulk upload product data to Subscription Service",
			Steps: []testlink.TestCaseStep{
				{
					Id:                0,
					TestCaseVersionId: 0,
					Number:            1,
					Actions:           "<p>Scenario: User want to do bulk upload happy flow</p><ul><li>User goes to subscription center</li><li>User click on tambah produk then click on bulk upload</li><li>Verify UI is correct based on design. </li></ul><p><ac:image ac:height=\"250\"><ri:attachment ri:filename=\"image2021-7-13_14-4-23.png\"></ri:attachment></ac:image></p><ul><li>User upload file to do bulk upload and click &#34;konfirmasi&#34;</li><li>UI will show confirmation pop up</li></ul><p><ac:image ac:height=\"250\"><ri:attachment ri:filename=\"image2021-7-12_15-3-32.png\"></ri:attachment></ac:image></p><ul><li>Click kembali, it will show previous pop up</li><li>Click konfirmasi again and click OK</li><li>Verify User will see pop up that giving ensure that file successfully uploaded and button &#34;konfirmasi&#34; is enabled<ac:image ac:height=\"250\"><ri:attachment ri:filename=\"image2021-7-12_15-6-28.png\"></ri:attachment></ac:image></li><li>Verify file successfully uploaded and ticker will show <ac:emoticon ac:name=\"question\"></ac:emoticon></li><li>Verify data uploaded is correct based on file being uploaded </li></ul>",
					ExpectedResults:   "<p><br/></p><ul><li>Verify UI is correct based on design. </li></ul><ul><li>Verify User will see pop up that giving ensure that file successfully uploaded and button &#34;konfirmasi&#34; is enabled</li><li>Verify file successfully uploaded and ticker will show <ac:emoticon ac:name=\"question\"></ac:emoticon></li><li>Verify data uploaded is correct based on file being uploaded </li></ul>",
					Active:            false,
					ExecutionType:     1,
				},
			},
			Status:              testlink.TestCaseStatusDraft,
			Importance:          testlink.TestImportanceMedium,
			Execution:           testlink.ExecutionTypeManual,
			Order:               0,
			InternalId:          0,
			CheckDuplicatedName: false,
		},
		{
			TestCaseName:  "BM-2639 - [Subscription Center] As a user, I want to bulk upload product data to Subscription Service",
			TestSuiteId:   37591638,
			TestProjectId: 37591635,
			AuthorLogin:   "andrian",
			Summary:       "BM-2639 - [Subscription Center] As a user, I want to bulk upload product data to Subscription Service",
			Steps: []testlink.TestCaseStep{
				{
					Id:                0,
					TestCaseVersionId: 0,
					Number:            1,
					Actions:           "<p>Scenario: User want to do bulk upload happy flow2639</p><ul><li>User goes to subscription center</li><li>User click on tambah produk then click on bulk upload</li><li>Verify UI is correct based on design. </li></ul><p><ac:image ac:height=\"250\"><ri:attachment ri:filename=\"image2021-7-13_14-4-23.png\"></ri:attachment></ac:image></p><ul><li>User upload file to do bulk upload and click &#34;konfirmasi&#34;</li><li>UI will show confirmation pop up</li></ul><p><ac:image ac:height=\"250\"><ri:attachment ri:filename=\"image2021-7-12_15-3-32.png\"></ri:attachment></ac:image></p><ul><li>Click kembali, it will show previous pop up</li><li>Click konfirmasi again and click OK</li><li>Verify User will see pop up that giving ensure that file successfully uploaded and button &#34;konfirmasi&#34; is enabled<ac:image ac:height=\"250\"><ri:attachment ri:filename=\"image2021-7-12_15-6-28.png\"></ri:attachment></ac:image></li><li>Verify file successfully uploaded and ticker will show <ac:emoticon ac:name=\"question\"></ac:emoticon></li><li>Verify data uploaded is correct based on file being uploaded </li></ul>",
					ExpectedResults:   "<p>2639</p><ul><li>Verify UI is correct based on design. </li></ul><ul><li>Verify User will see pop up that giving ensure that file successfully uploaded and button &#34;konfirmasi&#34; is enabled</li><li>Verify file successfully uploaded and ticker will show <ac:emoticon ac:name=\"question\"></ac:emoticon></li><li>Verify data uploaded is correct based on file being uploaded </li></ul>",
					Active:            false,
					ExecutionType:     1,
				},
			},
			Status:              testlink.TestCaseStatusDraft,
			Importance:          testlink.TestImportanceMedium,
			Execution:           testlink.ExecutionTypeManual,
			Order:               0,
			InternalId:          0,
			CheckDuplicatedName: false,
		},
	}

	for _, r := range validRequests {
		outboundMock.On("CreateTestCase", r).Return(nil)
	}
	type args struct {
		authorLogin string
		json        string
	}
	tests := []struct {
		name              string
		args              args
		outboundCallCount int
	}{
		{
			name: "Success create 2 test case",
			args: args{
				authorLogin: "andrian",
				json:        "{\n  \"id\": \"162592801\",\n  \"type\": \"page\",\n  \"status\": \"current\",\n  \"title\": \"Test\",\n  \"body\": {\n    \"storage\": {\n      \"value\": \"<table class=\\\"relative-table wrapped\\\" style=\\\"width: 262.5px;\\\"><colgroup><col style=\\\"width: 0.0px;\\\" /></colgroup><thead><tr><th style=\\\"text-align: left;\\\"><p>Services</p></th></tr></thead><tbody><tr><td style=\\\"text-align: left;\\\"><ol><li>Apps</li></ol></td></tr></tbody></table><p><br /></p><p><strong>List of Test Cases</strong></p><p><ac:structured-macro ac:name=\\\"toc\\\" ac:schema-version=\\\"1\\\" ac:macro-id=\\\"796bcdb5-744b-4b20-bbfa-46640bea8e02\\\" /></p><p><br /></p><p><br /></p><ac:structured-macro ac:name=\\\"info\\\" ac:schema-version=\\\"1\\\" ac:macro-id=\\\"34075dff-994a-45d7-9d84-bbf219559b43\\\"><ac:parameter ac:name=\\\"title\\\">Subscription Center</ac:parameter><ac:rich-text-body><h4><ac:structured-macro ac:name=\\\"jira\\\" ac:schema-version=\\\"1\\\" ac:macro-id=\\\"9e02f6dd-b62b-434f-ac98-f3a27a3bf563\\\"><ac:parameter ac:name=\\\"server\\\">Jira for GDN</ac:parameter><ac:parameter ac:name=\\\"serverId\\\">c4cc978b-cc62-3ae1-bbaf-f724a64b626b</ac:parameter><ac:parameter ac:name=\\\"key\\\">BM-2638</ac:parameter></ac:structured-macro></h4><table class=\\\"relative-table wrapped\\\" style=\\\"width: 87.5%;\\\"><colgroup><col style=\\\"width: 69.1923%;\\\" /><col style=\\\"width: 16.7977%;\\\" /><col style=\\\"width: 14.01%;\\\" /></colgroup><thead><tr><th class=\\\"highlight-grey\\\" data-highlight-colour=\\\"grey\\\"><h4><span style=\\\"color: rgb(153,204,0);\\\"><strong>BM-2638 -&nbsp;[Subscription Center] As a user, I want to bulk upload product data to Subscription Service</strong></span></h4></th><th colspan=\\\"1\\\"><p><span style=\\\"color: rgb(153,204,0);\\\">37591635</span></p></th><th colspan=\\\"1\\\"><span style=\\\"color: rgb(153,204,0);\\\"><a style=\\\"text-decoration: none;text-align: left;\\\" href=\\\"https://testlink.gdn-app.com/linkto.php?tprojectPrefix=TAL&amp;item=testsuite&amp;id=37591638\\\">37591638</a></span></th></tr></thead><tbody><tr><td colspan=\\\"3\\\"><div class=\\\"content-wrapper\\\"><p>other notes</p><ac:structured-macro ac:name=\\\"expand\\\" ac:schema-version=\\\"1\\\" ac:macro-id=\\\"c48e0c50-e4ec-4634-8d10-e415602934df\\\"><ac:rich-text-body><p>more notes</p><table class=\\\"relative-table wrapped\\\" style=\\\"width: 100.0%;\\\"><colgroup><col style=\\\"width: 2.24429%;\\\" /><col style=\\\"width: 28.6997%;\\\" /><col /><col style=\\\"width: 5.03264%;\\\" /><col style=\\\"width: 9.52122%;\\\" /><col style=\\\"width: 5.98477%;\\\" /><col style=\\\"width: 5.98477%;\\\" /><col style=\\\"width: 42.3694%;\\\" /></colgroup><tbody><tr><th class=\\\"numberingColumn\\\" style=\\\"text-align: center;\\\"><p>No</p></th><th style=\\\"text-align: center;\\\"><p><span style=\\\"color: rgb(153,204,0);\\\">Test Case</span></p></th><th colspan=\\\"1\\\"><span style=\\\"color: rgb(153,204,0);\\\">Expectation</span></th><th style=\\\"text-align: center;\\\">Reviewer</th><th style=\\\"text-align: center;\\\">Created by</th><th colspan=\\\"1\\\">Result Android</th><th colspan=\\\"1\\\">Result iOS</th><th colspan=\\\"1\\\">Notes</th></tr><tr><td class=\\\"numberingColumn\\\" style=\\\"text-align: center;\\\">1</td><td colspan=\\\"1\\\"><p>Scenario: User want to do bulk upload happy flow</p><ul><li>User goes to subscription center</li><li>User click on tambah produk then click on bulk upload</li><li>Verify UI is correct based on design.&nbsp;</li></ul><p><ac:image ac:height=\\\"250\\\"><ri:attachment ri:filename=\\\"image2021-7-13_14-4-23.png\\\" /></ac:image></p><ul><li>User upload file to do bulk upload and click &quot;konfirmasi&quot;</li><li>UI will show confirmation pop up</li></ul><p><ac:image ac:height=\\\"250\\\"><ri:attachment ri:filename=\\\"image2021-7-12_15-3-32.png\\\" /></ac:image></p><ul><li>Click kembali, it will show previous pop up</li><li>Click konfirmasi again and click OK</li><li>Verify User will see pop up that giving ensure that file successfully uploaded and button &quot;konfirmasi&quot; is enabled<ac:image ac:height=\\\"250\\\"><ri:attachment ri:filename=\\\"image2021-7-12_15-6-28.png\\\" /></ac:image></li><li>Verify file successfully uploaded and ticker will show&nbsp;<ac:emoticon ac:name=\\\"question\\\" /></li><li>Verify data uploaded is correct based on file being uploaded&nbsp;</li></ul></td><td colspan=\\\"1\\\"><p><br /></p><ul><li>Verify UI is correct based on design.&nbsp;</li></ul><ul><li>Verify User will see pop up that giving ensure that file successfully uploaded and button &quot;konfirmasi&quot; is enabled</li><li>Verify file successfully uploaded and ticker will show&nbsp;<ac:emoticon ac:name=\\\"question\\\" /></li><li>Verify data uploaded is correct based on file being uploaded&nbsp;</li></ul></td><td colspan=\\\"1\\\"><br /></td><td style=\\\"text-align: center;\\\" colspan=\\\"1\\\"><br /></td><td colspan=\\\"1\\\"><br /></td><td colspan=\\\"1\\\"><br /></td><td colspan=\\\"1\\\"><br /></td></tr></tbody></table></ac:rich-text-body></ac:structured-macro></div></td></tr></tbody></table><h4><ac:structured-macro ac:name=\\\"jira\\\" ac:schema-version=\\\"1\\\" ac:macro-id=\\\"13b801f3-be4f-4944-b7cf-149e2a6764d2\\\"><ac:parameter ac:name=\\\"server\\\">Jira for GDN</ac:parameter><ac:parameter ac:name=\\\"serverId\\\">c4cc978b-cc62-3ae1-bbaf-f724a64b626b</ac:parameter><ac:parameter ac:name=\\\"key\\\">BM-2638</ac:parameter></ac:structured-macro></h4><table class=\\\"relative-table wrapped\\\" style=\\\"width: 87.5%;\\\"><colgroup><col style=\\\"width: 69.1923%;\\\" /><col style=\\\"width: 16.7977%;\\\" /><col style=\\\"width: 14.01%;\\\" /></colgroup><thead><tr><th class=\\\"highlight-grey\\\" data-highlight-colour=\\\"grey\\\"><h4><span style=\\\"color: rgb(153,204,0);\\\"><strong>BM-2639 -&nbsp;[Subscription Center] As a user, I want to bulk upload product data to Subscription Service</strong></span></h4></th><th colspan=\\\"1\\\"><p><span style=\\\"color: rgb(153,204,0);\\\">37591635</span></p></th><th colspan=\\\"1\\\"><span style=\\\"color: rgb(153,204,0);\\\"><a style=\\\"text-decoration: none;text-align: left;\\\" href=\\\"https://testlink.gdn-app.com/linkto.php?tprojectPrefix=TAL&amp;item=testsuite&amp;id=37591638\\\">37591638</a></span></th></tr></thead><tbody><tr><td colspan=\\\"3\\\"><div class=\\\"content-wrapper\\\"><p>other notes</p><ac:structured-macro ac:name=\\\"expand\\\" ac:schema-version=\\\"1\\\" ac:macro-id=\\\"09034970-3110-4033-9221-260c9d1620cd\\\"><ac:rich-text-body><p>more notes</p><table class=\\\"relative-table wrapped\\\" style=\\\"width: 100.0%;\\\"><colgroup><col style=\\\"width: 2.24429%;\\\" /><col style=\\\"width: 28.6997%;\\\" /><col /><col style=\\\"width: 5.03264%;\\\" /><col style=\\\"width: 9.52122%;\\\" /><col style=\\\"width: 5.98477%;\\\" /><col style=\\\"width: 5.98477%;\\\" /><col style=\\\"width: 42.3694%;\\\" /></colgroup><tbody><tr><th class=\\\"numberingColumn\\\" style=\\\"text-align: center;\\\"><p>No</p></th><th style=\\\"text-align: center;\\\"><p><span style=\\\"color: rgb(153,204,0);\\\">Test Case</span></p></th><th colspan=\\\"1\\\"><span style=\\\"color: rgb(153,204,0);\\\">Expectation</span></th><th style=\\\"text-align: center;\\\">Reviewer</th><th style=\\\"text-align: center;\\\">Created by</th><th colspan=\\\"1\\\">Result Android</th><th colspan=\\\"1\\\">Result iOS</th><th colspan=\\\"1\\\">Notes</th></tr><tr><td class=\\\"numberingColumn\\\" style=\\\"text-align: center;\\\">1</td><td colspan=\\\"1\\\"><p>Scenario: User want to do bulk upload happy flow2639</p><ul><li>User goes to subscription center</li><li>User click on tambah produk then click on bulk upload</li><li>Verify UI is correct based on design.&nbsp;</li></ul><p><ac:image ac:height=\\\"250\\\"><ri:attachment ri:filename=\\\"image2021-7-13_14-4-23.png\\\" /></ac:image></p><ul><li>User upload file to do bulk upload and click &quot;konfirmasi&quot;</li><li>UI will show confirmation pop up</li></ul><p><ac:image ac:height=\\\"250\\\"><ri:attachment ri:filename=\\\"image2021-7-12_15-3-32.png\\\" /></ac:image></p><ul><li>Click kembali, it will show previous pop up</li><li>Click konfirmasi again and click OK</li><li>Verify User will see pop up that giving ensure that file successfully uploaded and button &quot;konfirmasi&quot; is enabled<ac:image ac:height=\\\"250\\\"><ri:attachment ri:filename=\\\"image2021-7-12_15-6-28.png\\\" /></ac:image></li><li>Verify file successfully uploaded and ticker will show&nbsp;<ac:emoticon ac:name=\\\"question\\\" /></li><li>Verify data uploaded is correct based on file being uploaded&nbsp;</li></ul></td><td colspan=\\\"1\\\"><p>2639</p><ul><li>Verify UI is correct based on design.&nbsp;</li></ul><ul><li>Verify User will see pop up that giving ensure that file successfully uploaded and button &quot;konfirmasi&quot; is enabled</li><li>Verify file successfully uploaded and ticker will show&nbsp;<ac:emoticon ac:name=\\\"question\\\" /></li><li>Verify data uploaded is correct based on file being uploaded&nbsp;</li></ul></td><td colspan=\\\"1\\\"><br /></td><td style=\\\"text-align: center;\\\" colspan=\\\"1\\\"><br /></td><td colspan=\\\"1\\\"><br /></td><td colspan=\\\"1\\\"><br /></td><td colspan=\\\"1\\\"><br /></td></tr></tbody></table></ac:rich-text-body></ac:structured-macro></div></td></tr></tbody></table></ac:rich-text-body></ac:structured-macro>\",\n      \"representation\": \"storage\",\n      \"_expandable\": {\n        \"content\": \"/rest/api/content/162592801\"\n      }\n    },\n    \"_expandable\": {\n      \"editor\": \"\",\n      \"view\": \"\",\n      \"export_view\": \"\",\n      \"styled_view\": \"\",\n      \"anonymous_export_view\": \"\"\n    }\n  },\n  \"extensions\": {\n    \"position\": \"none\"\n  },\n  \"_links\": {\n    \"webui\": \"/display/~andrian/Test\",\n    \"edit\": \"/pages/resumedraft.action?draftId=162592801\",\n    \"tinyui\": \"/x/IfiwCQ\",\n    \"collection\": \"/rest/api/content\",\n    \"base\": \"https://confluence.gdn-app.com\",\n    \"context\": \"\",\n    \"self\": \"https://confluence.gdn-app.com/rest/api/content/162592801\"\n  },\n  \"_expandable\": {\n    \"container\": \"/rest/api/space/~andrian\",\n    \"metadata\": \"\",\n    \"operations\": \"\",\n    \"children\": \"/rest/api/content/162592801/child\",\n    \"restrictions\": \"/rest/api/content/162592801/restriction/byOperation\",\n    \"history\": \"/rest/api/content/162592801/history\",\n    \"ancestors\": \"\",\n    \"version\": \"\",\n    \"descendants\": \"/rest/api/content/162592801/descendant\",\n    \"space\": \"/rest/api/space/~andrian\"\n  }\n}",
			},
			outboundCallCount: 2,
		},
		{
			name: "Failed no test invalid json",
			args: args{
				authorLogin: "andrian",
				json:        "",
			},
			outboundCallCount: 0,
		},
		{
			name: "Failed no test case found",
			args: args{
				authorLogin: "andrian",
				json:        "{}",
			},
			outboundCallCount: 0,
		},
	}

	totalCallCount := 0
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totalCallCount += tt.outboundCallCount
			_ = clipboard.WriteAll(tt.args.json)
			testlinkService.CreateTestCaseFromConfluenceContentClipboard(tt.args.authorLogin)
			outboundMock.AssertNumberOfCalls(t, "CreateTestCase", totalCallCount)
		})
	}
}
