import { createPinia, setActivePinia } from 'pinia';
import { describe, it, expect, beforeEach } from 'vitest';
import { useEditorStore } from './editor';
import { TrackLength } from '@/models/song';

describe('editor store', () => {
  beforeEach(() => {
    setActivePinia(createPinia());
  });

  describe('navigation', () => {
    it('should select the next tick as current on arrow down', () => {
      const store = useEditorStore();
      store.currentTick = 10;

      const event = new KeyboardEvent('keyup', { key: 'ArrowDown' });
      store.handleKey(event);

      expect(store.currentTick).toStrictEqual(11);
    });

    it('should select the previous tick as current on arrow up', () => {
      const store = useEditorStore();
      store.currentTick = 10;

      const event = new KeyboardEvent('keyup', { key: 'ArrowUp' });
      store.handleKey(event);

      expect(store.currentTick).toStrictEqual(9);
    });

    it('should not set the current tick outside of the defined range', () => {
      const store = useEditorStore();

      store.currentTick = 0;
      store.handleKey(new KeyboardEvent('keyup', { key: 'ArrowUp' }));
      expect(store.currentTick).toStrictEqual(0);

      store.currentTick = TrackLength - 1;
      store.handleKey(new KeyboardEvent('keyup', { key: 'ArrowDown' }));
      expect(store.currentTick).toStrictEqual(TrackLength - 1);
    });
  });
});
