//Input
const input = inputList[0].split("")

//Process & Output
const real = inputList[0].split("")
const mirror = input.reverse()
if (real.join("") === mirror.join("")) {
    console.log("Yes")
} else {
    console.log("No")
}