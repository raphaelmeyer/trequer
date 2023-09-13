const names = ['C', 'C#', 'D', 'D#', 'E', 'F', 'F#', 'G', 'G#', 'A', 'A#', 'B'];

export function keyName(key: number): string {
  if (key < 21) {
    return '---';
  }
  if (key > 127) {
    return '---';
  }

  const note = key % 12;
  const octave = Math.floor(key / 12) - 1;

  return `${names[note]}${octave}`;
}
