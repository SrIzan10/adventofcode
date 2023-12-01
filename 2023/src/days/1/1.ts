// REALLY UNCLEAN CODE BUT IT WORKS

const regex = /\D/g
const array: Array<number> = []
const file = (await (Bun.file('./src/days/1/input.txt')).text()).split('\n')
file.forEach((line) => {
    let arr = Array.from(line.replace(regex, '')).map(Number)
    console.log(`${arr} ${line}`)
    if (arr.length === 0) return
    if (arr.length === 1) {
        // lol
        arr = [Number(`${arr[0]}${arr[0]}`)]
    }
    if (arr.length >= 2) {
        arr = [Number(`${arr[0]}${arr[arr.length - 1]}`)]
    }
    array.push(...arr)
})
let sum = 0

array.forEach(num => sum += num)

console.log(sum)