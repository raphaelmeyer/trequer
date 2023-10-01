import { computed, ref } from 'vue';
import { defineStore } from 'pinia';

import type { ChannelId, MidiChannel, Note, Pattern, PatternId } from '@/models/song';

export const useSongStore = defineStore('song', () => {
  const channels = ref<Record<ChannelId, MidiChannel>>({});
  const patterns = ref<Record<PatternId, Pattern>>({});
  const sequence = ref<PatternId[]>([]);

  const track = computed(() => (pattern: PatternId, channel: ChannelId): (Note | undefined)[] => {
    const track = patterns.value[pattern].tracks[channel];

    return Array.from(Array(track.length).keys()).map((i) => track.notes[i]);
  });

  function setNote(note: Note): void {
    const track = patterns.value[0].tracks[0];
    track.notes[0] = note;
  }

  return { channels, patterns, sequence, track, setNote };
});
