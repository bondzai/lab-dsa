//Input
const input = inputList[0].split(" ")

//Process
const month = Number(input[0])
const year = Number(input[1]) - 543
let febDays = getFebDays()
const dayNumbers = ["", 31, febDays, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
function getFebDays() {
    let output = 28
    if (year % 4 == 0 && year % 100 != 0) {
        output = 29
    }
    return output
}

//Output
console.log(dayNumbers[month])