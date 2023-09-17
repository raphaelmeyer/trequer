import { createPinia, setActivePinia } from 'pinia';
import { describe, it, expect, beforeEach } from 'vitest';
import { useSongStore } from './song';

describe('song store', () => {
  beforeEach(() => {
    setActivePinia(createPinia());
  });

  it('should return undefined for empty notes', () => {
    const store = useSongStore();
    store.patterns[0] = { tracks: { 0: { length: 64, notes: {} } } };

    const track = store.track(0, 0);

    expect(track).toHaveLength(64);
    expect(track.at(0)).toBeUndefined();
    expect(track.at(17)).toBeUndefined();
    expect(track.at(63)).toBeUndefined();
  });

  it('should return notes at correct positions', () => {
    const store = useSongStore();
    store.patterns[0] = {
      tracks: {
        0: {
          length: 64,
          notes: { 7: { key: 52, length: 4, volume: 64 }, 11: { key: 69, length: 7, volume: 77 } },
        },
      },
    };

    const track = store.track(0, 0);

    expect(track).toHaveLength(64);
    expect(track.at(6)).toBeUndefined();
    expect(track.at(7)).toStrictEqual(expect.objectContaining({ key: 52, length: 4, volume: 64 }));
    expect(track.at(8)).toBeUndefined();
    expect(track.at(11)).toStrictEqual(expect.objectContaining({ key: 69, length: 7, volume: 77 }));
  });
});
