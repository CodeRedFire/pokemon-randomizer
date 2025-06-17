<script setup>
import ListPokemon from "@/views/randomizer/components/ListPokemon.vue";
import RefreshButton from "@/views/randomizer/components/RefreshButton.vue";
import SettingsButton from "@/views/randomizer/components/SettingsButton.vue";
import {selection_store} from "@/store/selection_store.js";
import Pokeball from "@/views/randomizer/components/Pokeball.vue";
</script>

<template>
  <SettingsButton/>
  <ListPokemon/>
  <div class="refreshContainer">
  <RefreshButton />
  </div>
  <div class="teamBuilderContainer">
    <Pokeball v-for="(pkmnNum, index) in selection_store.selectedPkmnList" :pkmnNum="pkmnNum"
    :index="index"></Pokeball>

    <span class="material-symbols-rounded resetBtn" @click="selection_store.reset()">
    refresh
    </span>
  </div>
</template>

<style scoped>
.refreshContainer {
  width: 100%;
  margin-top: 80px;
  text-align: center;
}

.teamBuilderContainer {
  display: flex;
  flex-wrap: wrap;
  position: sticky;
  bottom: 8px;
  z-index: 200;
  margin-top: 40px;
}
.resetBtn {
  cursor: pointer;
}
</style>
