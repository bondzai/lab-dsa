//Input
const x = "pupe"

//Process
function pupeToBinary(e) {
    let binary = e.replace(/u/g, "0")
    binary = binary.replace(/e/g, "1")
    binary = binary.replace(/p/g, "")
    const a = binary.split("")
    a.unshift("")
    const y = []
    let buffer = ""
    for(const i in a) {
        buffer += a[i]
        if (i % 8 == 0 && i > 1) {
            y.push(buffer)
            buffer = ""
        }
    }
    return y.join(" ")
}
function binaryToText(e) {
    let x = e.split(" ")
    const y = []
    for (const i in x) {
        y.push(String.fromCharCode(parseInt(x[i], 2)))
    }
    return y.join("")
}
let binary = pupeToBinary(x)
let output = binaryToText(binary)

//Output
console.log(binary)
console.log(output)