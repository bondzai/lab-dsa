//Input
const x = 2
const y = 6
const raw = [-4605.16,4768.87,-494.21,6548.12,
               9198.68,-5335.41,-7206.89,-9559.76, 
              -731.66,2229.36,-8204.81,1053.96]

//Process & Output
const eng = []
for (const i in raw) {
    eng.push(scaling(raw[i]))
}
const matrix = []
for (let i = 0; i < y; i++) {
    matrix[i] = []
    for (let j = 0; j < x; j++) {
        matrix[i].push(eng[0])
        eng.splice(0,1)
    }
    console.log(matrix[i].join(" "))
}

function scaling(x) {
    const input = {}
        input.max =  Math.max.apply(Math, raw)
        input.min = Math.min.apply(Math, raw)
        input.value = x
    const output = {}
        output.max = 1
        output.min = 0
    const m = (output.max - output.min) / (input.max- input.min)
    const c = output.min - (m * input.min)
    output.value = (m * x) + c
    return output.value.toFixed(4)
}