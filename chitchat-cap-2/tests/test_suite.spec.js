import { Builder, Browser } from 'selenium-webdriver';
import assert from "assert";
import { LoginPage } from './page_objects/login_page.js';
import { ErrorPage } from './page_objects/error_page.js';
import { PublicHomePage } from './page_objects/public_home_page.js';
import { PrivateHomePage } from './page_objects/private_home_page.js';
import { createTestUser, deleteTestUser, setDriverTimeouts } from './utils.js';


describe('E2E Tests', function () {
  let driver;
  let loginPage;
  let errorPage;
  let publicHomePage;
  let privateHomePage;
  let testUser;
  let unexistingTestUser;

  before(async function () {
    testUser = {
        name: "test",
        email: "test@test.com",
        password: "test",
    };
    unexistingTestUser = {
      email: "exist@doesnot.com",
      password: "aosfpbefibef",
    };
  });

  beforeEach(async function(){
    driver = await new Builder().forBrowser(Browser.CHROME).build();
    loginPage = new LoginPage(driver);
    errorPage = new ErrorPage(driver);
    publicHomePage = new PublicHomePage(driver);
    privateHomePage = new PrivateHomePage(driver);
    await createTestUser(driver);
    await setDriverTimeouts(driver);
  })

  afterEach(async function (){
    await deleteTestUser(driver);
    await driver.quit()
    driver = null;
    loginPage = null;
    errorPage = null;
    publicHomePage = null;
    privateHomePage = null;
  })

  it('Existing User is redirected to Private Home page.', async function () {
    this.timeout(60_000)
    await loginPage.goToPage();
    await loginPage.fillLogin(testUser.email);
    await loginPage.fillPassword(testUser.password);
    assert.equal(await loginPage.getLoginInputValue(), testUser.email);
    assert.equal(await loginPage.getPasswordInputValue(), testUser.password)
    
    await loginPage.clickLogInBtn()
    await privateHomePage.isVisible();
  });
  
  it('Unexisting User is redirected to Error page.', async function () {
    await loginPage.goToPage();
    await loginPage.fillLogin(unexistingTestUser.email);
    await loginPage.fillPassword(unexistingTestUser.password);
    assert.equal(await loginPage.getLoginInputValue(), unexistingTestUser.email);
    assert.equal(await loginPage.getPasswordInputValue(), unexistingTestUser.password)
    await loginPage.clickLogInBtn()
    await errorPage.showsMessage("Incorrect credentials.");
  });

  it('Existing User with wrong password is redirected to Error page.', async function () {
    this.timeout(60_000)
    await loginPage.goToPage();
    await loginPage.fillLogin(testUser.email);
    await loginPage.fillPassword(unexistingTestUser.password);
    assert.equal(await loginPage.getLoginInputValue(), testUser.email);
    assert.equal(await loginPage.getPasswordInputValue(), unexistingTestUser.password)
    
    await loginPage.clickLogInBtn()
    await errorPage.showsMessage("Incorrect credentials.");
  });

  it('Unlogged user is redirected to Login page when try to start a thread.', async function () {
    await publicHomePage.goToPage();
    await publicHomePage.clickStartThreadBtn();
    await loginPage.isVisible();
  });

  it('Unlogged user is redirected to Login page when clicks on the login navbar button', async function () {
    await publicHomePage.goToPage();
    await publicHomePage.clickNavBarLoginBtn();
    await loginPage.isVisible();
  });
  
  after(async () => {
    testUser = null;
    unexistingTestUser= null;
  });
});
