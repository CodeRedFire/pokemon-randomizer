<script setup>
import { MaxValues } from '@/enums/max_values.js';
import Pokemon from "@/components/Pokemon.vue";

let selectedPokemonIds = []

for (let i = 0; i < 6; i++) {
  selectedPokemonIds.push(Math.floor(Math.random() * MaxValues.Pokedex))
}

</script>

<template>
  <div class="main-container">
  <Pokemon v-for="i in selectedPokemonIds" :num-pokedex="i"/>
  </div>
</template>

<style scoped>
.main-container {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}
</style>
