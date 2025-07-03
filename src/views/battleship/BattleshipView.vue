<script setup>
import {battleship_store} from "@/store/battleship_store.js";
import BattleshipGrid from "@/views/battleship/components/BattleshipGrid.vue";
import RegenerateBtn from "@/views/battleship/components/RegenerateBtn.vue";
import SettingsButton from "@/views/commons/SettingsButton.vue";
import MenuButton from "@/views/commons/MenuButton.vue";

battleship_store.randomizeList();
</script>

<template>
  <BattleshipGrid/>
  <RegenerateBtn/>

  <SettingsButton/>
  <MenuButton/>
</template>

<style scoped>
</style>