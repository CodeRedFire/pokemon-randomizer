<script setup>

import MenuItem from "@/views/menu/components/MenuItem.vue";
</script>

<template>
  <MenuItem :title="$t('menu.randomizer')" link="#/randomizer"/>
  <MenuItem :title="$t('menu.battleship')" link="#/battleship"/>
</template>

<style scoped>
a {
  text-decoration: none;
}

div {
  background-color: #1976D2;
  padding: 20px 12px;
  margin: 20px;
  border-radius: 8px;
  text-decoration: none;
  color: #eeeeee;
  font-family: 'Roboto', sans-serif;
  letter-spacing: 0.5px;
  box-shadow: 2px 2px 4px rgba(0, 0, 0, 0.4);
  position: relative;
}

div:hover {
  box-shadow: 1px 1px 2px rgba(0, 0, 0, 0.4);
  background-color: #104b85;
}

.material-symbols-rounded {
  position: absolute;
  right: 12px;
  top: 18px;
  font-size: 24px;
  opacity: 0.5;
}
</style>