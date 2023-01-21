//input
const inputList = []
inputList[0] = "3h^4 + 6h^3 - 3h^2 + h + 4"
const input = inputList[0]

//Process
const reg = {}
reg.sign = /(\+|\-)/g
reg.var = /[A-z]/
reg.num = /\d+/g
const split = input.split(reg.sign)
function diferential(word) {
    const getNum = /\d+/g
    const mainVar = reg.var.exec(input)[0]
    let output = null  
    let buffer = ""
    let coefficient = ""
    let digiPatern = ""
    let degree = ""
    if (word.match(getNum) != null) {
        buffer = word.match(getNum)
        if (buffer.length > 1) {
            degree = buffer[1] - 1
            digiPatern = mainVar + "^"
            coefficient = buffer[0] * buffer[1]      
            if (degree == 1) {
                degree = ""
                digiPatern = mainVar
                coefficient = buffer[0] * buffer[1]
            }
        }
        if (buffer.length == 1) {
            degree = ""
            digiPatern = ""
            coefficient = buffer[0]
        }
        output = coefficient + digiPatern + degree
    }
    else {
        output = 1
    }
   return output
}
const output = []
for (const i in split) {
    if (reg.var.test(split[i]) == true) {
        output.push(diferential(split[i]))
    } 
    else if (reg.num.test(split[i]) == false) {
        output.push(split[i])
    }
}
if (reg.num.test(output[output.length - 1]) == false) {
    output.pop()
}

//Output
console.log(output.join(" "))