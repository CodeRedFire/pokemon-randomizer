<script setup>
import SettingsButton from "@/views/commons/SettingsButton.vue";
import {battleship_store} from "@/store/battleship_store.js";
import PkmnTile from "@/views/battleship/components/PkmnTile.vue";

battleship_store.randomizeList();
</script>

<template>
  <div class="grid">
    <div style="border: 1px solid rgba(0, 0, 0, 0.4);">
  <div v-for="row in battleship_store.pkmnList" class="row">
    <PkmnTile v-for="pkmnNbr in row" :nbrPkmn="pkmnNbr"/>
  </div>
  </div>
  </div>
  <SettingsButton/>
</template>

<style scoped>
.row {
  display: flex;
}
.grid {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
}
</style>