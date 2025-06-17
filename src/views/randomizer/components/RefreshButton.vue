<script setup>
import { config_store } from '@/store/config_store.js'
import {selection_store} from "@/store/selection_store.js";

function refresh() {
  config_store.randomizeList();
  selection_store.next();
}
</script>

<template>
<button class="refresh" @click="refresh()">{{ $t("randomizer.btn_label")}}</button>
</template>

<style scoped>
.refresh {
  background-color: #FFD152;
  border: none;
  padding: 8px 28px;
  border-radius: 8px;
  font-weight: bold;
  font-size: 20px;
  box-shadow: 2px 2px 4px rgba(0, 0, 0, 0.4);
  color: #321300;
  cursor: pointer;
}
</style>