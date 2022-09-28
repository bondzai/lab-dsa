//Input
const n = "CDL"

//Process
const x = n.split("")
const keys = ["I", "V", "X", "L", "C", "D", "M"]
const values = [1, 5, 10, 50, 100, 500, 1000]
const y = []
for (let i in x) {
    for (let j in keys) {
        if (x[i] == keys[j]) {
            y.push(values[j])
        }
    }
}
function calRoman(a, b) {
    if (a >= b) {
        return a + b
    } else {
        return b - a
    } 
}

//Output
console.log(y.reduce(calRoman))