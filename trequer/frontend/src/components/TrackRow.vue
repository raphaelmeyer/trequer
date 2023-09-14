<script setup lang="ts">
import { keyName } from '@/models/midi';
import type { Note } from '@/models/song';

const props = defineProps<{
  note: Note;
}>();
</script>

<template>
  <div class="flex bg-sky-950 w-32 px-2">
    <div class="w-10">{{ keyName(props.note.key) }}</div>
    <div class="w-8">{{ props.note.volume }}</div>
    <div class="w-8">{{ props.note.length }}</div>
  </div>
</template>
