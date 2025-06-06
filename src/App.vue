<script setup>
import { ref, computed } from 'vue'
import RandomizerView from '@/views/randomizer/RandomizerView.vue'
import Config from '@/views/config/Config.vue'

const routes = {
  '/': RandomizerView,
  '/config': Config
}

const currentPath = ref(window.location.hash)

window.addEventListener('hashchange', () => {
  currentPath.value = window.location.hash
})

const currentView = computed(() => {
  return routes[currentPath.value.slice(1) || '/'] || RandomizerView
})
</script>

<template>
  <component :is="currentView" />
</template>

<style scoped>

</style>
