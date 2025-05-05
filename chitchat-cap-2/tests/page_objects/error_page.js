import { By } from 'selenium-webdriver';

export class ErrorPage {
    constructor(driver){
        this.driver = driver;
        this.url = "http://localhost:8082/err";
    }

    showsMessage(message){
        return this.driver.findElement(By.xpath(`\/\/*[text()='${message}']`)).isDisplayed();
    }
}
