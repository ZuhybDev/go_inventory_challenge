const freqLetter = (letters: string[]) => {
  const freq: Record<string, number> = {};

  for (const letter of letters) {
    freq[letter] = (freq[letter] || 0) + 1;
  }
  console.log(freq);
};

const letter = ["A", "B", "C", "D", "E", "A"];
console.log(freqLetter(letter));
