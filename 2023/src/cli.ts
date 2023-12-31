import { readdir, mkdir } from "fs/promises";
import * as commander from 'commander';
import logger from "./util/logger.ts";

const program = new commander.Command;
program
	.command('run <day>')
	.description('Runs the given day and part')
	.option('-p, --part <part>', 'The part of the day to run', '1')
	.action(async (day: string, option: { part: string }) => {
		const dirs = await readdir(`./src/days/${day}`, { withFileTypes: true });
		const file = dirs.find((dir) => dir.name === `${option.part}.ts`);
		if (file) {
			logger('success', `Running day ${day} part ${option.part}`)
			await import(`./days/${day}/${option.part}.ts`);
		} else {
			logger('error', `Day ${day} part ${option.part} not found!`);
		}
	});
program
	.command('input <day>')
	.description('Fetches and downloads the input for the given day')
	.action(async (day) => {
		const inputFetch = await fetch(`https://adventofcode.com/${import.meta.dir.split('/').at(-2)}/day/${day}/input`, {
			headers: {
				cookie: `session=${process.env.AOC_SESSION}`
			}
		});
		const input = await inputFetch.text();

		await Bun.write(`./src/days/${day}/input.txt`, input);
		logger('success', `Fetched input for day ${day}`);
	});
program
	.command('init <day>')
	.description('Initializes a new day')
	.action(async (day) => {
		await mkdir(`./src/days/${day}`);
		await Bun.write(`./src/days/${day}/1.ts`, `import getInput from '../../util/getInput.js';\n\n`);
		await Bun.write(`./src/days/${day}/2.ts`, `import getInput from '../../util/getInput.js';\n\n`);
		logger('success', `Initialized day ${day}! Happy hacking!`)
	});

program.parse(process.argv);