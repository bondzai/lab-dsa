//Input
const rawData = "4/10/2022"
let inputBuffer = rawData.split("/")
const input = {}
    input.date = Number(inputBuffer[0])
    input.month = Number(inputBuffer[1])
    input.year = Number(inputBuffer[2])

//Process
function calLeapYear(currentYear) {
    const output = {}
    output.totalDays = 365
    output.febDays = 28
    if ((currentYear % 4 === 0) && (currentYear % 100 != 0) || (currentYear % 400 == 0)) {
        output.totalDays = 366
        output.febDays = 29
    }
    return output
}
function getWeekDay(dateDuration) {
    const days = ["Sunday","Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"]
    const startDay = "Monday"
    const startDayIndex = days.indexOf(startDay)
    const period = Number(dateDuration) + Number(startDayIndex)
    let index = period
    if (period > 7) {
        index = (period % 7)
    }
    let endDay = days[index]
    return endDay
}
function getDateDuration() {
    const months = {}
    months.name = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]
    months.dayNumbers = [31, calLeapYear().febDays, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    const dateDuration = []
    const yearDuration = input.year - 1900
    let buffer = 0
    for (let i = 0; i < yearDuration; i++) {
        buffer = 1900 + Number(i)
        dateDuration.push(calLeapYear(buffer).totalDays)
    }
    for (let i = 0; i < input.month - 1; i++) {
        dateDuration.push(months.dayNumbers[i])
    }
    dateDuration.push(input.date)
    return dateDuration.reduce((a, b) => Number(a) + Number(b)) 
}

let dateDuration = getDateDuration() - 1
let result = getWeekDay(dateDuration)
console.log(result)