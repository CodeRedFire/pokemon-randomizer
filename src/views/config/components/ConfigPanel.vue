<script setup>
import {config_store} from '@/store/config_store.js'
import {MaxValues} from "@/enums/max_values.js";
import {Languages} from "@/enums/languages.js";
import {DisplayedDataOptions} from "@/enums/displayed_data.js";
</script>

<template>
  <div class="container">
    <div class="line">
      <div>{{ $t("config.language") }}</div>
      <select v-model="$i18n.locale">
        <option v-for="locale in $i18n.availableLocales" :key="`locale-${locale}`" :value="locale">{{ locale }}</option>
      </select>
    </div>

    <div class="line">
      <div>{{ $t("config.nbrPkmn") }}</div>
      <input type="number" min="1" :max="MaxValues.Pokedex"
             v-model="config_store.nbrPokemon">
    </div>

    <div class="line">
      <div>{{ $t("config.randomizerLanguage") }}</div>
      <select v-model="config_store.displayedLanguage">
        <option v-for="l in Languages" :value="l"> {{ $t("languages['" + l + "']") }}</option>
      </select>
    </div>

    <div class="line">
      <div>{{ $t("config.displayedInformations") }}</div>
      <select v-model="config_store.displayedInformations">
        <option v-for="d in DisplayedDataOptions" :value="d"> {{ $t("config['" + d + "']") }}</option>
      </select>
    </div>
  </div>
</template>

<style scoped>
.container {
  background-color: #D2D7DF;
  padding: 4px 8px;
  border-radius: 8px;
  box-shadow: 2px 2px 4px 0 rgba(0, 0, 0, 0.4);
  font-family: "Roboto", sans-serif;
  font-size: 14px;
}
.line {
  padding: 8px 12px;
}

.line div {
  width: 280px;
  display: inline-block;
  vertical-align: top;
}

.line input {
  width: 200px;
}

.line select {
  width: 218px;
}

input, select {
  border: solid rgba(52, 67, 94, 0.4) 1px;
  border-radius: 4px;
  padding: 4px 8px;
  font-size: 14px;
  color: #34435e;
}
</style>