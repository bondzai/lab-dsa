function solution(inputString) {
    let output = false
    const mirror = inputString.split("").reverse().join("")
    if (inputString === mirror) {
        output = true
    }
    return output
}

console.log(solution("aabaa"))