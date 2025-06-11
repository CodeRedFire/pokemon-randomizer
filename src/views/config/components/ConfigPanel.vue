<script setup>
import { config_store } from '@/store/config_store.js'
import {MaxValues} from "@/enums/max_values.js";
import {Languages} from "@/enums/languages.js";
</script>

<template>
  <div>
    <select v-model="$i18n.locale">
      <option v-for="locale in $i18n.availableLocales" :key="`locale-${locale}`" :value="locale">{{ locale }}</option>
    </select>
  </div>
  <div>

    <input type="number" min="1" :max="MaxValues.Pokedex"
           v-model="config_store.nbrPokemon">
  </div>

  <div>
    <select v-model="config_store.displayedLanguage">
      <option v-for="l in Languages" :value="l"> {{$t("languages['" + l + "']")}}</option>
    </select>
  </div>
</template>

<style scoped>

</style>