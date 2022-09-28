//Output format example
//7 15 15 4 13 15 18 14 9 14 7 : goodmorning
//9 12 15 22 5 25 15 21 : i love you

//Input
const input = "7 15 15 4 13 15 18 14 9 14 7"

//Process
const inputWords = input.split(" ") 
const chars = ["", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
              "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"]
const deCoding = []
for (let i in inputWords) {
    deCoding.push(chars[inputWords[i]])
}

//Output
console.log(deCoding.join(""))