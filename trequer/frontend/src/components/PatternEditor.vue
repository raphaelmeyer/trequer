<script setup lang="ts">
import { computed } from 'vue';

import { useEditorStore } from '@/stores/editor';
import { useSongStore } from '@/stores/song';
import LabelBox from './LabelBox.vue';
import TrackRow from '@/components/TrackRow.vue';

const song = useSongStore();
const track = song.track(0, 0);

const editor = useEditorStore();
const current = computed(() => {
  return { channel: editor.currentChannel, tick: editor.currentTick };
});

const channels = [0, 1, 2, 3, 4];
function channelName(channel: number): string {
  return `Channel ${channel}`;
}
</script>

<template>
  <div class="flex gap-x-2">
    <LabelBox
      v-for="(channel, index) in channels"
      :key="index"
      :label="channelName(channel)"
      class="w-32"
    >
      <div class="flex flex-col">
        <TrackRow
          v-for="(note, noteIndex) in track"
          :key="noteIndex"
          :note="note"
          :selected="channel === current.channel && noteIndex === current.tick"
        />
      </div>
    </LabelBox>
  </div>
</template>
