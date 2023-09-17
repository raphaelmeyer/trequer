<script setup lang="ts">
import TrackRow from '@/components/TrackRow.vue';
import { useSongStore } from '@/stores/song';
import LabelBox from './LabelBox.vue';

const song = useSongStore();
const track = song.track(0, 0);

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
        <TrackRow v-for="(note, noteIndex) in track" :key="noteIndex" :note="note" />
      </div>
    </LabelBox>
  </div>
</template>
