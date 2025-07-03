<script setup>
import SettingsButton from "@/views/commons/SettingsButton.vue";
import {battleship_store} from "@/store/battleship_store.js";
import PkmnTile from "@/views/battleship/components/PkmnTile.vue";

battleship_store.randomizeList();
</script>

<template>
  <div class="grid-container">
  <div class="grid">
    <div>
      <div v-for="row in battleship_store.pkmnList" class="row">
        <PkmnTile v-for="pkmnNbr in row" :nbrPkmn="pkmnNbr"/>
      </div>
    </div>
  </div>
  </div>
</template>

<style scoped>
.row {
  display: flex;
}
.grid-container {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
}
.grid {
  background-color: rgba(29, 52, 97, 0.5);
  padding: 12px;
  border-radius: 12px;
}
</style>