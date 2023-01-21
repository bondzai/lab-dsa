function initialize(element) {
    let output = null
    if (element >= -1000 && element <= 1000) {
        output = element
    }
    return output
}

function solution(param1, param2) {
    const a = initialize(param1)
    const b = initialize(param2)
    return a + b
}

console.log(solution(1, 3))