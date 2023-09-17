<script setup lang="ts">
import type { Note } from '@/models/song';
import TrackRow from '@/components/TrackRow.vue';
import { useSongStore } from '@/stores/song';

const song = useSongStore();
const track = song.track(0, 0);
</script>

<template>
  <div class="bg-sky-950 border-sky-600 border-4 rounded-xl w-32 p-2">
    <TrackRow v-for="(note, index) in track" :key="index" :note="note" />
  </div>
</template>
