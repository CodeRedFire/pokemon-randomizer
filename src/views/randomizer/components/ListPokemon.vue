<script setup>
import Pokemon from "@/views/randomizer/components/Pokemon.vue";
import {MaxValues} from "@/enums/max_values.js";
import {computed, onMounted} from "vue";
import { config_store } from '@/store/config_store.js'

onMounted(() => {config_store.randomizeList()});

const selectedPokemonIds = computed(() => {
  return config_store.pkmnList;
})


</script>

<template>
  <div class="main-container">
    <Pokemon v-for="i in selectedPokemonIds" :key="i" :num-pokedex="i"/>
  </div>
</template>

<style scoped>
.main-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}
</style>