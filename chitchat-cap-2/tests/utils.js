export function createTestUser(driver){
    return driver.get("http://localhost:8082/test/user/create");
}

export function deleteTestUser(driver){
    return driver.get("http://localhost:8082/test/user/delete");
}

export function setDriverTimeouts(driver){
    return driver.manage().setTimeouts({implicit: 10_000, script: 10_000, pageLoad: 10_00})
}
