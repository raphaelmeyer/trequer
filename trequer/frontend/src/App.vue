<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';

import { Quit } from '../wailsjs/runtime';
</script>

<template>
  <header>
    <div class="">
      <nav>
        <RouterLink to="/">Trequer</RouterLink>
        <RouterLink to="/settings">Settings</RouterLink>
        <div @click="Quit()">Exit</div>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

<style scoped></style>
