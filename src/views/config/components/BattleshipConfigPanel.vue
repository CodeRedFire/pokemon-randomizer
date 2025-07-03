<script setup>
import {config_store} from '@/store/config_store.js'
import {MaxValues} from "@/enums/max_values.js";
import {Languages} from "@/enums/languages.js";
import {DisplayedDataOptions} from "@/enums/displayed_data.js";
import {battleship_store} from "@/store/battleship_store.js";
</script>

<template>
  <h2>{{ $t("config.battleship.title") }}</h2>
  <div class="container">
    <div class="line">
      <div>{{ $t("config.battleship.column") }}</div>
      <input type="number" min="1" :max="10"
             v-model="battleship_store.nbrColumns">
    </div>
    <div class="line">
      <div>{{ $t("config.battleship.row") }}</div>
      <input type="number" min="1" :max="10"
             v-model="battleship_store.nbrRows">
    </div>
  </div>
</template>

<style scoped>
h2 {
  font-family: "Roboto", sans-serif;
}
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