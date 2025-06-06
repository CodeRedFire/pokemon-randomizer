<script setup>
import ListPokemon from "@/views/randomizer/components/ListPokemon.vue";
import RefreshButton from "@/views/randomizer/components/RefreshButton.vue";
import SettingsButton from "@/views/randomizer/components/SettingsButton.vue";
</script>

<template>
  <SettingsButton/>
  <ListPokemon/>
  <div class="refreshContainer">
  <RefreshButton />
  </div>
</template>

<style scoped>
.refreshContainer {
  width: 100%;
  margin-top: 80px;
  text-align: center;
}
</style>
