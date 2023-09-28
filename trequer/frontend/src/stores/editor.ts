import { ref } from 'vue';
import { defineStore } from 'pinia';

import type { ChannelId, PatternId, Tick } from '@/models/song';

export const useEditorStore = defineStore('editor', () => {
  const currentPattern = ref<PatternId>(0);
  const currentChannel = ref<ChannelId>(0);
  const currentTick = ref<Tick>(0);

  return { currentPattern, currentChannel, currentTick };
});
