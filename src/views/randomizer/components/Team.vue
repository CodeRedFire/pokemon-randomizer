<script setup>

import {selection_store} from "@/store/selection_store.js";
import Pokeball from "@/views/randomizer/components/Pokeball.vue";
</script>

<template>
  <div class="teamBuilderContainer">
    <Pokeball v-for="(pkmnNum, index) in selection_store.selectedPkmnList" :pkmnNum="pkmnNum"
              :index="index"></Pokeball>

    <span class="material-symbols-rounded resetBtn" @click="selection_store.reset()" :title="$t('randomizer.team.reset_tooltip')">
    refresh
    </span>
  </div>
</template>

<style scoped>
.teamBuilderContainer {
  display: flex;
  flex-wrap: wrap;
  position: absolute;
  bottom: 0px;
  z-index: 200;
  left: calc( 100vw / 2 - 240px);
  background-color: #752d2d;
  padding: 4px 12px;
  border-radius: 8px 8px 0px 0px;
  border: solid 4px #5c2424;
}
.resetBtn {
  cursor: pointer;
}
</style>