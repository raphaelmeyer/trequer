import { ref } from 'vue';
import { defineStore } from 'pinia';

import type { ChannelId, MidiChannel, Pattern, PatternId } from '@/models/song';

export const useSongStore = defineStore('song', () => {
  const channels = ref<Record<ChannelId, MidiChannel>>({});
  const patterns = ref<Record<PatternId, Pattern>>({});
  const sequence = ref<PatternId[]>([]);

  return { channels, patterns, sequence };
});
