//Input
const input = 1317919

//Process
const prime = filterPrime(input)
const primeRange = 120
const component = filterComponet(input)
const primeComponent = []
const output = []
function checkPrime(n, primeRange) {
    const fraction = []
    let result = true
    for (let i = 2; i <= primeRange; i++) { if (n % i == 0) { fraction.push(i) } }
    if (fraction.length > 1 || n == 1) { result = false }
    return result
}
function filterPrime(primeRange) {
    const result = []
    for (let i = 2; i <= primeRange; i++) { if (checkPrime(i, primeRange) == true) { result.push(i) } }
    return result
}
function filterComponet(n) {
    const result = []
    for (let i = 0; i <= n; i++) {
        if (n % i == 0) {
            result.push(i)
        }
  }
  return result
}
for (const i in component) {
    for (const j in prime) {
        if (component[i] == prime[j]) {
            primeComponent.push(component[i])
        }
    }
}
primeComponent.sort((a, b) => b - a)
let current = input
let k = 0
while (current > 1) {
    k = getPrimeComponent(current)
    output.push(k)
    current = current / k
}
function getPrimeComponent(num) {
    for (const i in primeComponent) {
        if (num % primeComponent[i] == 0) {
            return primeComponent[i]
        }
    }
}

//Output
output.sort((a, b) => a - b)
for (const i in output) {
    console.log(output[i])
}