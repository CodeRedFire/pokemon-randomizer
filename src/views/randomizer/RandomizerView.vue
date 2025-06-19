<script setup>
import ListPokemon from "@/views/randomizer/components/ListPokemon.vue";
import RefreshButton from "@/views/randomizer/components/RefreshButton.vue";
import SettingsButton from "@/views/commons/SettingsButton.vue";
import Team from "@/views/randomizer/components/Team.vue";
</script>

<template>
  <SettingsButton/>
  <div class="randomizerContainer">
    <ListPokemon/>
  </div>
  <div class="refreshContainer">
    <RefreshButton />
  </div>
  <Team></Team>
</template>

<style scoped>
.randomizerContainer {
  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  overflow: auto;
}
.refreshContainer {
  width: 100%;
  margin-top: 80px;
  margin-bottom: 80px;
  text-align: center;
  position: absolute;
  bottom: 0;
  left: 0;
}
</style>
