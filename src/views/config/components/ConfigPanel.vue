<script setup>
import { config_store } from '@/store/config_store.js'
import {MaxValues} from "@/enums/max_values.js";
</script>

<template>
<input type="number" min="1" :max="MaxValues.Pokedex"
 v-model="config_store.nbrPokemon">
</template>

<style scoped>

</style>