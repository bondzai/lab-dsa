//Input
const inputList = []
inputList[0] = "6 2 8"

//Process
const raw = inputList[0].split(" ")
const cal = {}
cal.n = Number(raw[0])
cal.m = Number(raw[1])
cal.y = Number(raw[2])
cal.k = (cal.m * cal.y) + cal.n - cal.y
let b = cal.k / (cal.m - 1)
let a = b + cal.n

//Output
console.log(a + " " + b)