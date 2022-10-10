//Output format example
//[5,4,50] => false

//INput
const input = [1,1,2,1]

//Process
const temporarySum = []
const sum = []
let buffer = 0
for (const i in input) {
    for (const j in input) {
        if (i != j) {
            temporarySum.push(input[j])
        }
    }
    buffer = temporarySum.reduce((a, b) => a + b)
    sum.push(buffer)
    temporarySum.splice(0, temporarySum.length)
}
const output = []
for (const i in sum) {
    if (input[i] < sum[i]) {
        output.push(true)
    } else {
        output.push(false)
    }
}
let result = "true"
for (const i in output) {
    if (output[i] != true) {
        result = "false"
        break
    }
}

//Output
console.log(result)