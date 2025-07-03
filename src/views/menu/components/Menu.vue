<script setup>

import MenuItem from "@/views/menu/components/MenuItem.vue";
</script>

<template>
  <div>
    <MenuItem :title="$t('menu.randomizer')" link="#/randomizer" :img="`/pokemon-randomizer/assets/imgs/randomizer.png`"/>
    <MenuItem :title="$t('menu.battleship')" link="#/battleship" :img="`/pokemon-randomizer/assets/imgs/battleship.png`"/>
  </div>
</template>

<style scoped>
div {
  display: flex;
  flex-direction: row;
  gap: 10px;
}
</style>