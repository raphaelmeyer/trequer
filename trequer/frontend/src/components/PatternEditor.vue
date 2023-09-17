<script setup lang="ts">
import TrackRow from '@/components/TrackRow.vue';
import { useSongStore } from '@/stores/song';
import LabelBox from './LabelBox.vue';

const song = useSongStore();
const track = song.track(0, 0);
</script>

<template>
  <LabelBox label="Channel 0" class="w-32">
    <div class="flex flex-col">
      <TrackRow v-for="(note, index) in track" :key="index" :note="note" />
    </div>
  </LabelBox>
</template>
