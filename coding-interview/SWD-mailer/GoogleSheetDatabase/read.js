function doGet(e) {
    let action = e.parameter.action
    if (action == 'readData') {
        return readData(e)
      }
}

function readData(e) {
    const ss = SpreadsheetApp.openById('1yJZXdZqUGh0RcZnQQyO72z3h3_LdKk2wt3yQF1kPi3c')
    const sheet = ss.getSheetByName('USER')
    const columnRange = sheet.getRange(1, 1, 1, sheet.getLastColumn());
    const columnRangeValues = columnRange.getValues();
    const keys = columnRangeValues[0]
    const rows = sheet.getRange(2,1,sheet.getLastRow()-1,sheet.getLastColumn()).getValues();
    const data = [];
    for (const i in rows) {
        let row = rows[i]
        const record = {}
        for (const j in row) {
          let key = keys[j]
          let value = row[j]
          record[key]=value
        }
        data.push(record)
    }
     let result = JSON.stringify(data)
    return ContentService.createTextOutput(result).setMimeType(ContentService.MimeType.JSON)
}