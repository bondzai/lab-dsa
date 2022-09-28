//Input
const raw = "12 5"

//Process
const input = raw.split(" ")
const n = input[0] / 2
const r = input[1]
const output = {}
function createBase() {
    const output = []
    for (let i = 0; i < n; i++) {
        output.push("*")
    }
    return output
}
function createLeft() {
    const output = []
    for (let i = 0; i < n; i++) {
        if (i < r) {
            output.push("-")
        }
        else {
            output.push("*")
        }
    }
    return output
}
output.left = createLeft()
function createHigh() {
    const output = []
    const left = createLeft()
    const right = createLeft().reverse()
    const base = createBase()
    for (let i = 0; i < r; i++) {
        if (i == 0) {
            output.push(left.join("") + right.join(""))
        } else {
            left.shift()
            left.push("*")
            right.pop()
            right.unshift("*")
            output.push(left.join("") + right.join(""))           
        }
    }
    for (let i = 0; i < n - r; i++) {
        output.push(base.join("") + base.join(""))
    }
    return output
}

//Output
output.high = createHigh()
for(const i in output.high) {
    console.log(output.high[i])
}
output.base = createHigh().reverse()
for(const i in output.base) {
    console.log(output.base[i])
}