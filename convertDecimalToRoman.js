//Input
const input = 1905

//Process
const output = []
const roman = {}
    roman.char = ["M", "D", "C", "L", "X", "V", "I"]
    roman.value = [1000, 500, 100, 50, 10, 5, 1]
    roman.stack = [3, 3, 3, 3, 3, 3, 3]
const digits = getDigit()
function getDigit(){
    const digit = {}
    digit.raw = []
    digit.value = []
    for (const i of input.toString()) {
        digit.raw.push(i)
    }
    digit.raw.reverse()
    let x = "1"
    for (const i in digit.raw) {
        if (i == 0) {
            digit.value.push(digit.raw[i])
        } else {
            digit.value.push(digit.raw[i] * Number(x))
        }
        x += 0
    }
    return digit.value.reverse()
}
const minusCase = /(4|9)/
let current = 0
let buffer = 0
for (const i in digits){
    current = digits[i]
    if (minusCase.test(digits[i]) == false) {
        while (current > 0) {
            buffer = getValuePlus(current)
            output.push(buffer)
            current = current - roman.value[roman.char.indexOf(buffer)]
        }
    } else {
        buffer = getValueMinus(current)
        output.push(buffer)
    }
}
function getValuePlus(currentValue) {
    let output = null
    for (const i in roman.value) {
        if (currentValue >= roman.value[i] && roman.stack[i] > 0) {
            output = roman.char[i]
            roman.stack[i]--
            break
        }
    }   
    return output
}
function getValueMinus(currentValue) {
    let output = null
    for (const i in roman.value) {
        if (currentValue >= roman.value[i]) {
            output = roman.char[Number(i) + 1] + roman.char[i - 1]
            break
        }
    }
    return output
}

//Output
console.log(output)