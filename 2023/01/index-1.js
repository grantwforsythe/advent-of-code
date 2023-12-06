const fs = require('fs');
const path = require('path');

try {
  const data = fs.readFileSync(path.join(__dirname, 'data.txt'), 'utf8');
  const lines = data.trim().split('\n');

  const answer = lines.reduce((accumulator, line) => {
    // Remove all non-digits from the line
    const strippedLine = line.replace(/\D/g, '');
    // Get the first and last number from the stripped line
    const firstNumber = Number(strippedLine[0]);
    const lastNumber = Number(strippedLine[strippedLine.length - 1]);

    const currentValue =
      isNaN(firstNumber) || isNaN(lastNumber)
        ? accumulator
        : firstNumber * 10 + lastNumber;

    return accumulator + currentValue;
  }, 0);

  console.log(answer);
} catch (error) {
  console.error(error);
}
