<script setup>

import Menu from "@/views/menu/components/Menu.vue";
import SettingsButton from "@/views/commons/SettingsButton.vue";
</script>

<template>
  <h1>{{ $t("menu.title") }}</h1>
<Menu/>
  <SettingsButton/>
</template>

<style scoped>
h1 {
  text-align: center;
  font-family: "Roboto", sans-serif;
}
</style>