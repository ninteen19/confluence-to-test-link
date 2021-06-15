package main

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// initialize a new api instance
	//api, err := goconfluence.NewAPI("https://confluence.gdn-app.com/rest/api", "andrian@gdn-commerce.com", "")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//res, err := api.GetContentByID("162002871", goconfluence.ContentQuery{
	//	Expand:   []string{"body.storage", "version"},
	//})
	//if err != nil {
	//	log.Fatal(res, err)
	//}
	//fmt.Printf("%+v\n", res)

	str := "<table class=\"relative-table wrapped\" style=\"width: 262.5px;\"><colgroup><col style=\"width: 0.0px;\" /></colgroup><thead><tr><th style=\"text-align: left;\"><p>Services</p></th></tr></thead><tbody><tr><td style=\"text-align: left;\"><ol><li>Apps</li></ol></td></tr></tbody></table><p><br /></p><p><strong>List of Test Cases</strong></p><p><ac:structured-macro ac:name=\"toc\" ac:schema-version=\"1\" ac:macro-id=\"796bcdb5-744b-4b20-bbfa-46640bea8e02\" /></p><h3>BM-2647 -&nbsp;[Mobile Apps] Enable instant delivery for Subscription on checkout one (shipping)</h3><ac:structured-macro ac:name=\"expand\" ac:schema-version=\"1\" ac:macro-id=\"1d2167d6-94cd-44c6-8ef1-1e667c37a3ba\"><ac:rich-text-body><table class=\"relative-table wrapped\" style=\"width: 100.0%;\"><colgroup><col style=\"width: 2.24429%;\" /><col style=\"width: 28.6997%;\" /><col style=\"width: 5.03264%;\" /><col style=\"width: 9.52122%;\" /><col style=\"width: 5.98477%;\" /><col style=\"width: 5.98477%;\" /><col style=\"width: 42.3694%;\" /></colgroup><tbody><tr><th class=\"numberingColumn\" style=\"text-align: center;\"><p>No</p></th><th style=\"text-align: center;\"><p>Test Case</p></th><th style=\"text-align: center;\">Reviewer</th><th style=\"text-align: center;\">Created by</th><th colspan=\"1\">Result Android</th><th colspan=\"1\">Result iOS</th><th colspan=\"1\">Notes</th></tr><tr><td class=\"numberingColumn\" style=\"text-align: center;\">1</td><td colspan=\"1\"><p><strong>Scenario</strong>: User checkout using STANDARD logistic for 1 product</p><ul><li>User choose product that available for STANDARD</li><li>Go to checkout 1</li><li>Check UI should be correct for the design and wording</li><li>Verify that product can use STANDARD on the logistic available</li><li>Verify that another logistic other than STANDARD is available / not available&nbsp;</li><li>Verify that another logistic that available can be chosen</li><li>checkout using STANDARD success and order created&nbsp;</li></ul></td><td colspan=\"1\"><br /></td><td style=\"text-align: center;\" colspan=\"1\"><div class=\"content-wrapper\"><p><ac:link><ri:user ri:userkey=\"2c90419550dd20fe01519e773d930054\" /></ac:link></p></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"a377df6f-e75d-4fa3-af1a-a2fb352da2d3\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>62</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>63</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"4dbcb517-31cd-4ee5-bdc3-6f2de0d6bb0c\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>64</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>65</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><br /></td></tr><tr><td class=\"numberingColumn\">2</td><td colspan=\"1\"><p><strong>Scenario</strong>: User checkout using EXPRESS logistic for 1 product</p><ul><li>User choose product that available for&nbsp;EXPRES</li><li>Go to checkout 1</li><li>Check UI should be correct for the design and wording</li><li>Verify that product can use EXPRESS on the logistic available</li><li>Verify that another logistic other than EXPRESSis available / not available&nbsp;</li><li>Verify that another logistic that available can be choosen</li><li>checkout using EXPRESS success and order created&nbsp;</li></ul></td><td colspan=\"1\"><br /></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><ac:link><ri:user ri:userkey=\"2c90419550dd20fe01519e773d930054\" /></ac:link></p></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"a5093e1b-b030-4249-b88f-3e75a8f5a59e\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>537</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>538</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"4f264dad-2f86-466f-b18c-d322d4fa8305\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>539</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>540</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><br /></td></tr><tr><td class=\"numberingColumn\">3</td><td><p><strong>Scenario</strong>: Verify that every logistic option that available for channel Subscription mobile and Subscription web</p><ul><li>Make EXPRESS not available for subs mobile and web</li><li>Go to checkout, express logistic option is not showing</li><li>Make EXPRESS available for subs mobile and web</li><li>Go to checkout, express logistic option is showing</li></ul></td><td><br /></td><td><div class=\"content-wrapper\"><p><ac:link><ri:user ri:userkey=\"2c90419550dd20fe01519e773d930054\" /></ac:link></p></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"c0e2f827-77d3-4bd0-9e44-1c94e4df9add\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>541</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>542</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"924a6fee-1253-4989-b07f-80e47e952c07\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>543</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>544</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><br /></td></tr><tr><td class=\"numberingColumn\" colspan=\"1\">4</td><td colspan=\"1\"><p><strong>Scenario</strong>: User checkout and STANDARD FREE SHIPPING</p><ul><li>Go to checkout 1</li><li>Verify that product can use STANDARD FREE SHIPPING</li><li>Check UI should be correct for the design and wording</li><li>Verify user changing to other logistic, free shipping wont be applied</li><li>Verify user changing to FREE SHIPPING, free shipping applied again</li></ul></td><td colspan=\"1\"><br /></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><ac:link><ri:user ri:userkey=\"2c90419550dd20fe01519e773d930054\" /></ac:link></p></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"07d48aa0-379f-49ea-89fd-abcf10dead31\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>545</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>546</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"27ffc3dd-d58b-4c38-8ffc-3f129c5f4ac8\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>547</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>548</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><br /></td></tr><tr><td class=\"numberingColumn\">5</td><td><p><strong>Scenario</strong>: Verify the default logistic showing on checkout 1</p><p>try to make STANDARD as default logistic</p><ul><li>User choose product subs&nbsp;that available for STANDARD or EXPRESS</li><li>User go to checkout 1</li><li>Check UI should be correct for the design and wording</li><li>Verify default logistic will be the same logic like retail&nbsp;&rarr; STANDARD</li><li>Verify that another logistic other than EXPRESS is available</li></ul></td><td><br /></td><td><div class=\"content-wrapper\"><p><ac:link><ri:user ri:userkey=\"2c90419550dd20fe01519e773d930054\" /></ac:link></p></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"c792e7b0-adb1-4775-bada-69950325be7a\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>549</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>550</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"6c352877-a0ae-4869-9be3-bcfdcee8991f\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>551</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>552</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><br /></td></tr><tr><td class=\"numberingColumn\" colspan=\"1\">6</td><td colspan=\"1\"><p><strong>Scenario</strong>: User checkout 2 PACKAGE all using STANDARD (different location)</p><ul><li>User choose product that available for both logistic</li><li>Go to checkout 1</li><li>Check UI should be correct for the design and wording</li><li>Verify that one package can use STANDARD</li><li>checkout using both logistic</li><li>success and order created&nbsp;</li></ul></td><td colspan=\"1\"><br /></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><ac:link><ri:user ri:userkey=\"2c90419550dd20fe01519e773d930054\" /></ac:link></p></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"2535ba13-4472-48da-ba5a-55348ac9498f\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>553</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>554</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"a08ecc0e-9bf7-4a7c-9efc-2732efa4502f\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>555</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>556</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><br /></td></tr><tr><td class=\"numberingColumn\" colspan=\"1\">7</td><td colspan=\"1\"><p><strong>Scenario</strong>: User checkout 2 PACKAGE using STANDARD and EXPRESS</p><ul><li>User choose product that available for both logistic</li><li>Go to checkout 1</li><li>Check UI should be correct for the design and wording</li><li>Verify that one package can use STANDARD and another package can use EXPRESS</li><li>checkout using both logistic</li><li>success and order created&nbsp;</li></ul></td><td colspan=\"1\"><br /></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><ac:link><ri:user ri:userkey=\"2c90419550dd20fe01519e773d930054\" /></ac:link></p></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"d857f387-eecf-40f3-9173-10f638a614d5\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>557</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>558</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"cc6bb0d2-8272-43c5-a941-032ef6746aab\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>559</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>560</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><br /></td></tr><tr><td class=\"numberingColumn\" colspan=\"1\">8</td><td colspan=\"1\"><p><strong>Scenario</strong>: User checkout 2 PACKAGE using STANDARD and EXPRESS</p><ul><li>User choose product that available for both logistic</li><li>Go to checkout 1</li><li>Check UI should be correct for the design and wording</li><li>Verify that one package can use STANDARD and another package can use EXPRESS</li></ul><p><strong>Try to edit one package EXPRESS become STANDARD&nbsp;</strong></p><ul><li>Verify that two package will be merged and use STANDARD</li></ul><p><strong>Try to edit one package STANDARD become&nbsp;EXPRESS</strong></p><ul><li>Verify that two package will be merged and use&nbsp;EXPRESS</li></ul></td><td colspan=\"1\"><br /></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><ac:link><ri:user ri:userkey=\"2c90419550dd20fe01519e773d930054\" /></ac:link></p></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"f0924c52-ae6c-49ee-a107-56d9db04e302\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>561</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>562</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"4c938168-1e34-47ed-bbcd-163dca5f3d9f\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>563</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>564</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><br /></td></tr><tr><td class=\"numberingColumn\" colspan=\"1\">9</td><td colspan=\"1\"><p><strong>Scenario</strong>: User checkout 1 PACKAGE 2 items, ALL STANDARD</p><ul><li>User choose product that available for mentioned logistic</li><li>Go to checkout 1</li><li>Check UI should be correct for the design and wording</li><li>Verify that one package can use STANDARD&nbsp;</li><li>checkout using both logistic</li></ul></td><td colspan=\"1\"><br /></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><ac:link><ri:user ri:userkey=\"2c90419550dd20fe01519e773d930054\" /></ac:link></p></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"dc33dc80-3bb2-439c-bc72-e01daad2f621\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>565</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>566</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"6b46a07f-d0db-4794-81d5-2f6a9acd358b\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>567</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>568</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><br /></td></tr><tr><td class=\"numberingColumn\" colspan=\"1\">10</td><td colspan=\"1\"><p><strong>Scenario</strong>: User checkout 1 PACKAGE 2 items, ALL EXPRESS</p><ul><li>User choose product that available for mentionedlogistic</li><li>Go to checkout 1</li><li>Check UI should be correct for the design and wording</li><li>Verify that one package can use EXPRESS</li><li>checkout using both logistic</li><li>success and order created&nbsp;</li></ul></td><td colspan=\"1\"><br /></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><ac:link><ri:user ri:userkey=\"2c90419550dd20fe01519e773d930054\" /></ac:link></p></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"7fb4d075-f0a8-441a-978a-b1776e7c9e94\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>569</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>570</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"8fcbf2ba-dd1c-4151-88c5-e3ee2571540d\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>571</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>572</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><br /></td></tr><tr><td class=\"numberingColumn\" colspan=\"1\">11</td><td colspan=\"1\"><p><strong>Scenario</strong>: Check the consolidation in checkout two warehouse one zone</p><ul><li>User choose product that available for logistic standard and<strong>&nbsp;item come from two different warehouse 1 zone</strong></li><li>Go to checkout 1</li><li>Check UI should be correct for the design and wording</li><li>Verify that one package can use standard and <strong>consolidation is happened</strong></li></ul></td><td colspan=\"1\"><br /></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><ac:link><ri:user ri:userkey=\"2c90419550dd20fe01519e773d930054\" /></ac:link></p></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"4459b935-381b-4540-819e-45739240e18b\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>573</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>574</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"91c91359-945c-436a-b485-9984b94e64e2\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>575</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>576</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><br /></td></tr><tr><td class=\"numberingColumn\" colspan=\"1\">12</td><td colspan=\"1\"><p><strong>Scenario</strong>: Check condition of consolidation in checkout if logistic using EXPRESS</p><ul><li>User choose product that available for logistic express and<strong>&nbsp;item come from two different warehouse 1 zone</strong></li><li>Go to checkout 1</li><li>Check UI should be correct for the design and wording</li><li>Verify that <strong>consolidation is not happened&nbsp;</strong>(because consolidation only available for STANDARD only)</li></ul></td><td colspan=\"1\"><br /></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><ac:link><ri:user ri:userkey=\"2c90419550dd20fe01519e773d930054\" /></ac:link></p></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"f70b2f22-3143-406c-b745-b89a6ddffc67\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>577</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>578</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><div class=\"content-wrapper\"><p><span class=\"placeholder-inline-tasks\"><ac:structured-macro ac:name=\"status\" ac:schema-version=\"1\" ac:macro-id=\"d2fd9852-df3d-4fc6-849b-0e5338fad7d0\"><ac:parameter ac:name=\"colour\">Grey</ac:parameter><ac:parameter ac:name=\"title\">PENDING</ac:parameter><ac:parameter ac:name=\"\" /></ac:structured-macro></span></p><ac:task-list>\n<ac:task>\n<ac:task-id>579</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">EN</span></ac:task-body>\n</ac:task>\n<ac:task>\n<ac:task-id>580</ac:task-id>\n<ac:task-status>incomplete</ac:task-status>\n<ac:task-body><span class=\"placeholder-inline-tasks\">ID</span></ac:task-body>\n</ac:task>\n</ac:task-list></div></td><td colspan=\"1\"><br /></td></tr></tbody></table></ac:rich-text-body></ac:structured-macro>"
	str = strings.ReplaceAll(str, "&nbsp;", " ")
	str = strings.ReplaceAll(str, "&rarr;", " ")
	regex := regexp.MustCompile("<table.*?>(.|\n)*?</table>")
	tables := regex.FindAllString(str, -1)

	data := &Table{}
	err := xml.Unmarshal([]byte(tables[1]), data)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return
	}

	//fmt.Println(tables[1])
	for _, row := range data.Body.Rows {
		if len(row.Data) > 0 {
			fmt.Println(row.Data[1])
		}
		fmt.Println("===================================")
	}

	//result, err := json.Marshal(data)
	//if nil != err {
	//	fmt.Println("Error marshalling to JSON", err)
	//	return
	//}

	//testlink.Conf.Key = ""
	//testlink.Conf.Url = ""
}

type Table struct {
	Body Body `xml:"tbody" json:"body"`
}

type Body struct {
	Rows []Row `xml:"tr" json:"rows"`
}

type Row struct {
	Data []Data `xml:"td" json:"data"`
}

type Data struct {
	Value string `xml:",innerxml"`
}
