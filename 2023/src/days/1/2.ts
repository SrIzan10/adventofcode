import getInput from "../../util/getInput.ts"

// REALLY UNCLEAN CODE BUT IT WORKS
const regex = /\D/g
const array: Array<number> = []

let file = await getInput(1)
file = file.replaceAll("twone", "21")
file = file.replaceAll("sevenine", "79")
file = file.replaceAll("oneight", "18")
file = file.replaceAll("threeight", "38")
file = file.replaceAll("nineight", "98")
file = file.replaceAll("fiveight", "58")
file = file.replaceAll("eighthree", "83")
file = file.replaceAll("eightwo", "82")
file = file.replaceAll("one", "1")
file = file.replaceAll("two", "2")
file = file.replaceAll("three", "3")
file = file.replaceAll("four", "4")
file = file.replaceAll("five", "5")
file = file.replaceAll("six", "6")
file = file.replaceAll("seven", "7")
file = file.replaceAll("eight", "8")
file = file.replaceAll("nine", "9")
const fileSplit = file.split("\n");

fileSplit.forEach((line) => {
    let arr = Array.from(line.replace(regex, '')).map(Number)
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