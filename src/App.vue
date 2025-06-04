<script setup>
import { MaxValues } from '@/enums/max_values.js';
import Pokemon from "@/components/Pokemon.vue";

let selectedPokemonIds = []

for (let i = 0; i < 6; i++) {
  selectedPokemonIds.push(Math.floor(Math.random() * MaxValues.Pokedex))
}

</script>

<template>
  <Pokemon v-for="i in selectedPokemonIds" :num-pokedex="i"/>
</template>

<style scoped>
</style>
