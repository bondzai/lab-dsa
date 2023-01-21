//Input
const input = inputList[0]

//Process
const elements = ["Fire", "Water", "Wind", "Ground", "Light", "Dark"]  
let elementIndex = elements.indexOf(input)
const output = []
    if (elementIndex == -1) {
        console.log("No have this mahou in your library.")
    } else {
        let decimal = elementIndex.toString(2)
        for (let i of decimal) {
            output.push(i)
        }
        for (let i = 0; i < 20 - decimal.length; i++){
            output.unshift("0")
        }
    }

//Output
console.log(output.join(""))