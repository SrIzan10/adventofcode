export default async function getInput(day: number) {
    return await (Bun.file(`./src/days/${day}/input.txt`)).text()
}