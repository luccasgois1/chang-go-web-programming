import { By } from 'selenium-webdriver';

export class PublicHomePage {
    constructor(driver){
        this.driver = driver;
        this.url = "http://localhost:8082/";
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

    navBarLoginBtn(){
        return this.driver.findElement(By.id("login-navbar-btn"));
    }

    clickNavBarLoginBtn(){
        return this.navBarLoginBtn().click();
    }
}
