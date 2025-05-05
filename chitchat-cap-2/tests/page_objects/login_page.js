import { By } from 'selenium-webdriver';

export class LoginPage {
    constructor(driver){
        this.driver = driver;
        this.url = "http://localhost:8082/login";
    }

    isVisible(){
        return this.loginBtn().isDisplayed();
    }
    
    loginInput(){
        return this.driver.findElement(By.name("email"));
    }

    passwordInput(){
        return this.driver.findElement(By.name("password"));
    }

    loginBtn(){
        return this.driver.findElement(By.id("login-form-btn"));
    }
    
    goToPage(){
        return this.driver.get(this.url);
    }

    waitToLoad(){
        return this.driver.getTitle();
    }

    fillLogin(text){
        return this.loginInput().sendKeys(text);
    }

    async getLoginInputValue(){
        return await this.loginInput().getAttribute("value");
    }

    getPasswordInputValue(){
        return this.passwordInput().getAttribute("value");
    }

    fillPassword(text){
        return this.passwordInput().sendKeys(text);
    }

    clickLogInBtn(){
        return this.loginBtn().click();
    }
}
