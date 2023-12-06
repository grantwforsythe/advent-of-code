const fs = require('fs');
const path = require('path');

const digits = {
  one: 1,
  two: 2,
  three: 3,
  four: 4,
  five: 5,
  six: 6,
  seven: 7,
  eight: 8,
  nine: 9,
};

const regex = /one|two|three|four|five|six|seven|eight|nine|\d/gi;

const mapMatchToDigit = (match) => {
  return match[0].length > 1 ? digits[match[0]] : Number(match[0]);
};

try {
  const data = fs.readFileSync(path.join(__dirname, 'data.txt'), 'utf8');
  const lines = data.trim().split('\n');

  const answer = lines.reduce((accumulator, line) => {
    const matches = Array.from(line.matchAll(regex));

    if (matches.length === 0) return accumulator;

    const firstNumber = mapMatchToDigit(matches[0]);
    const lastNumber = mapMatchToDigit(matches[matches.length - 1]);
    const currentValue = isNaN(firstNumber * 10 + lastNumber)
      ? 0
      : firstNumber * 10 + lastNumber;
    console.log(currentValue);

    // TODO: Get correct answer
    return accumulator + currentValue;
  }, 0);

  console.log(answer);
} catch (error) {
  console.error(error);
}
