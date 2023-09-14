import { describe, it, expect } from 'vitest';

import { keyName } from './midi';

describe('Midi', () => {
  it.each([
    { key: 21, note: 'A0' },
    { key: 22, note: 'A#0' },
    { key: 35, note: 'B1' },
    { key: 48, note: 'C3' },
    { key: 61, note: 'C#4' },
    { key: 74, note: 'D5' },
    { key: 87, note: 'D#6' },
    { key: 100, note: 'E7' },
    { key: 113, note: 'F8' },
    { key: 126, note: 'F#9' },
    { key: 127, note: 'G9' },
  ])('converts midi key $key to note $note', ({ key, note }) => {
    expect(keyName(key)).toStrictEqual(note);
    expect(keyName(22)).toStrictEqual('A#0');
    expect(keyName(60)).toStrictEqual('C4');
    expect(keyName(75)).toStrictEqual('D#5');
    expect(keyName(127)).toStrictEqual('G9');
    expect(keyName(128)).toStrictEqual('');
  });

  it.each([{ key: 0 }, { key: 20 }, { key: 128 }])(
    'should return empty string for out of range value $key',
    ({ key }) => {
      expect(keyName(key)).toStrictEqual('');
    }
  );
});
