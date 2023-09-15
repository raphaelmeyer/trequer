<script setup lang="ts">
import { keyName } from '@/models/midi';
import type { Note } from '@/models/song';

const props = defineProps<{
  note: Note;
}>();
</script>

<template>
  <div class="flex px-2 rounded-md [&:nth-child(4n+1)]:bg-sky-900">
    <div class="w-10">{{ keyName(props.note.key) }}</div>
    <div class="w-8 text-right">{{ props.note.volume }}</div>
    <div class="w-8 text-right">{{ props.note.length }}</div>
  </div>
</template>
