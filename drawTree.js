//Input
const input = 3
const n = Number(input)

//Process
function createBase() {
    const trunk = []; const root = [];
    for (let i = 0; i < n; i++) { root.push("="); trunk.push(" "); }
    root.push("V"); trunk.push("|");
    for (let i = 0; i < n; i++) { root.push("="); }
    console.log(trunk.join("")); console.log(root.join(""));
}

function createStack() {
    const stack = []; const buffer = []; const range = (2 * n) + 1;
    for (let i = 0; i < range; i++) { buffer.push("*"); }
    for (let i = 0; i < n + 1; i++) {
        stack.unshift(buffer.join("")); buffer.pop(); buffer.pop(); buffer.unshift(" ");
    }
    return stack.reverse()
}

const fullStack = []
const stack = createStack()
for (let i = 0; i < n; i++) {
    for (let j in stack) {
        fullStack.push(stack[j])
    }
stack.shift()
}
fullStack.reverse()

//Output
for (let i in fullStack) {
    console.log(fullStack[i])
}
createBase()