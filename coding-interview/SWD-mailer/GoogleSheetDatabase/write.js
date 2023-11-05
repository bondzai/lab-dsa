function doPost(e) {
    let action = e.parameter.action
    if (action == 'approve') {
        return approveData(e)
    }
    if (action == 'login') {
        return checkLogin(e)
    }
}

function approveData(e) {
    let data = JSON.parse(e.postData.contents)
    let id = data.id;
    const ss = SpreadsheetApp.openById('1yJZXdZqUGh0RcZnQQyO72z3h3_LdKk2wt3yQF1kPi3c')
    const sheet = ss.getSheetByName('USER')
    if (id === 'user1') {
      sheet.getRange('E2').setValue("true")
    } else if (id === 'user2') {
      sheet.getRange('E3').setValue("true")
    } else if (id === 'user3') {
      sheet.getRange('E4').setValue("true")
    }
  return ContentService.createTextOutput("approve success").setMimeType(ContentService.MimeType.TEXT)
}

function checkLogin(e) {
    let data = JSON.parse(e.postData.contents)
    let id = data.id;
    const ss = SpreadsheetApp.openById('1yJZXdZqUGh0RcZnQQyO72z3h3_LdKk2wt3yQF1kPi3c')
    const sheet = ss.getSheetByName('USER')
    if (id === 'user1') {
      sheet.getRange('D2').setValue("true")
    } else if (id === 'user2') {
      sheet.getRange('D3').setValue("true")
    } else if (id === 'user3') {
      sheet.getRange('D4').setValue("true")
    }
  return ContentService.createTextOutput("login success").setMimeType(ContentService.MimeType.TEXT)
}

function testApprove() {
  approveData('user2');
}
function testLogin() {
  checkLogin('user2');
}