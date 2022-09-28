//Input
const input = "I Love You"

//Process
const output = []
const digitShift = 4
let buffer = 0
for (const i in input) {
  buffer = digitShift + input.charCodeAt(i)
  output.push(String.fromCharCode(buffer))
}

//Output
console.log(output.join(""))