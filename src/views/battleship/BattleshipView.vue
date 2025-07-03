<script setup>
import {battleship_store} from "@/store/battleship_store.js";
import BattleshipGrid from "@/views/battleship/components/BattleshipGrid.vue";
import RegenerateBtn from "@/views/battleship/components/RegenerateBtn.vue";
import SettingsButton from "@/views/commons/SettingsButton.vue";

battleship_store.randomizeList();
</script>

<template>
  <BattleshipGrid/>
  <RegenerateBtn/>

  <SettingsButton/>
</template>

<style scoped>
</style>