//Output format example
/*
         __
      __|
   __|
__|

*/

//Input
const n = 0

//Process
const gear = n * 3
const output = {}
function createUpper() {   
    const output = []
    for (let i = 0; i < gear; i++) {
        output.push(" ")
    }
    for (let i = 0; i < 2; i++) {
        output.push("_")
    }
    return output    
}
function createLower() {
    const output = []
    const paddle = gear - 3
    for (let i = 0; i < paddle; i++) {
        output.push(" ")
    }
    for (let i = 0; i < gear - paddle - 1; i++) {
        output.push("_")
    }
    output.push("|")
    return output
}

//Output
function createStairs() {
    const output = []
    const upper = createUpper()
    const lower = createLower()
    for (let i = 0; i < n; i++) {
        if (i ==0) {
            console.log(upper.join(""))
            console.log(lower.join(""))
        } else {
            lower.shift()
            lower.shift()
            lower.shift()
            console.log(lower.join(""))
        }
    }    
}
if (n == 0) {
    console.log("__")
} else {
    createStairs()
}