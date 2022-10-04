function solution(inputArray) {
    let output = null
    const products = []
    let buffer = 0
    for (const i in inputArray) {
        buffer = inputArray[i] * inputArray[Number(i)+1]
        if (isNaN(buffer) == false) {
            products.push(buffer)
        }
    }
    output = Math.max.apply(Math, products)
    return output
}

console.log(solution([3, 6, -2, -5, 7, 3]))