const input = require('fs').readFileSync('/dev/stdin', 'utf8');
const [testCases, ...tests] = input.split('\n');

const [SAND, L_DIAMOND] = ['.', '<'];


for (let i = 0; i < testCases; ++i) {
    let diamonds = 0;
    let stack = [];

    for (const element of tests[i]) {
        if (element === SAND)
            continue;


        if (element === L_DIAMOND) {
            stack.push(element);

        } else if (stack.length > 0) {
            ++diamonds;
            stack.pop();
        }
    }

    console.log(diamonds)
}