import { By } from 'selenium-webdriver';

export class PrivateHomePage {
    
    constructor(driver){
        this.driver = driver;
        this.url = "http://localhost:8082/";
    }

    isVisible(){
        return Promise.all([
            this.startThreadBtn().isDisplayed(),
            this.chitChatBtn().isDisplayed(),
            this.navBarHomeBtn().isDisplayed(),
            this.navBarLogoutBtn().isDisplayed(),
        ])
    }
    
    goToPage(){
        return this.driver.get(this.url);
    }

    startThreadBtn(){
        return this.driver.findElement(By.id("start-thread-home"));
    }

    clickStartThreadBtn(){
        return this.startThreadBtn().click();
    }

    chitChatBtn(){
        return this.driver.findElement(By.id("chitchat-headed-btn"));
    }

    clickChitChatBtn(){
        return this.chitChatBtn().click();
    }

    navBarHomeBtn(){
        return this.driver.findElement(By.id("home-navbar-btn"));
    }

    clickNavBarHomeBtn(){
        return this.navBarHomeBtn().click();
    }

    navBarLogoutBtn(){
        return this.driver.findElement(By.id("logout-navbar-btn"));
    }

    clickNavBarLogoutBtn(){
        return this.navBarLogoutBtn().click();
    }
}
