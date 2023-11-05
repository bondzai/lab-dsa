//Input
const startDay = inputList[0]
const n = inputList[1]

//Process
const days = ["sunday","monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"]
let i = days.indexOf(startDay)
let period = Number(n) + Number(i)
let index = period
if (period > 7) {
    index = (period % 7)
}
let endDay = days[index]

//Output
console.log(endDay)