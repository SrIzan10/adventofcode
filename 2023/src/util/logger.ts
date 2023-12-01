import kleur from "kleur";

export default function logger(type: 'success' | 'error', message: string) {
    switch (type) {
        case 'success':
            console.log(`${kleur.bold().bgGreen().white('✓ SUCCESS')} ${message}`);
            break;
        case 'error':
            console.log(`${kleur.bold().bgRed().white('✗ ERROR')} ${message}`);
            break;
    }
}