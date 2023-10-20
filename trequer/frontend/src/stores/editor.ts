import { ref } from 'vue';
import { defineStore } from 'pinia';

import { TrackLength, type ChannelId, type PatternId, type Tick } from '@/models/song';
import { useSongStore } from './song';

export const useEditorStore = defineStore('editor', () => {
  const currentPattern = ref<PatternId>(0);
  const currentChannel = ref<ChannelId>(0);
  const currentTick = ref<Tick>(0);

  const song = useSongStore();

  const keys: Record<string, number> = {
    a: 12,
    w: 13,
    s: 14,
    e: 15,
    d: 16,
    f: 17,
    t: 18,
    g: 19,
    y: 20,
    h: 21,
    u: 22,
    j: 23,
  };

  function handleKey(event: KeyboardEvent): void {
    if (!event.shiftKey && !event.altKey && !event.ctrlKey) {
      switch (event.key) {
        case 'a':
        case 'w':
        case 's':
        case 'e':
        case 'd':
        case 'f':
        case 't':
        case 'g':
        case 'y':
        case 'h':
        case 'u':
        case 'j':
          song.setNote({ key: keys[event.key], volume: 64, length: 2 });
          break;

        case 'ArrowUp':
          currentTick.value = _clamp(0, currentTick.value - 1, TrackLength - 1);
          break;
        case 'ArrowDown':
          currentTick.value = _clamp(0, currentTick.value + 1, TrackLength - 1);
          break;
        case 'ArrowLeft':
          currentChannel.value--;
          break;
        case 'ArrowRight':
          currentChannel.value++;
          break;

        default:
          break;
      }
    }
  }

  function _clamp(min: number, max: number, value: number): number {
    return Math.max(min, Math.min(value, max));
  }

  return { currentPattern, currentChannel, currentTick, handleKey };
});
