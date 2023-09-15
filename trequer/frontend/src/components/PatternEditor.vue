<script setup lang="ts">
import type { Note } from '@/models/song';
import TrackRow from '@/components/TrackRow.vue';

const notes: Note[] = [
  { tick: 0, key: 57, volume: 64, length: 2 },
  { tick: 4, key: 64, volume: 64, length: 2 },
  { tick: 8, key: 72, volume: 64, length: 2 },
  { tick: 12, key: 64, volume: 64, length: 2 },

  { tick: 16, key: 57, volume: 64, length: 2 },
  { tick: 20, key: 64, volume: 64, length: 2 },
  { tick: 24, key: 72, volume: 64, length: 2 },
  { tick: 28, key: 64, volume: 64, length: 2 },

  { tick: 32, key: 57, volume: 64, length: 2 },
  { tick: 36, key: 64, volume: 64, length: 2 },
  { tick: 40, key: 72, volume: 64, length: 2 },
  { tick: 44, key: 64, volume: 64, length: 2 },

  { tick: 48, key: 57, volume: 64, length: 2 },
  { tick: 52, key: 64, volume: 64, length: 2 },
  { tick: 56, key: 72, volume: 64, length: 2 },
  { tick: 60, key: 64, volume: 64, length: 2 },
];
</script>

<template>
  <div class="bg-sky-950 border-sky-600 border-4 rounded-xl w-32 p-2">
    <TrackRow v-for="(note, index) in notes" :key="index" :note="note" />
  </div>
</template>
