function solution(year) {
    let output = 0
    output = Math.ceil(year / 100)
    return output
}

console.log(solution(1905))