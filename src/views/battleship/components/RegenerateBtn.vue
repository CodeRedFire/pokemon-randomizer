<script setup lang="ts">
import { battleship_store } from '@/store/battleship_store';


</script>

<template>
  <div class="regenerateBtn">
    <button @click="battleship_store.randomizeList()">{{ $t("battleship.regenerate") }}</button>
  </div>
</template>

<style scoped>
.regenerateBtn {
  width: 200px;
  text-align: center;
  position: absolute;
  bottom: 0;
  left: calc(50% - 100px);
}

button {
  background-color: #FFD152;
  border: none;
  padding: 8px 28px;
  border-radius: 8px;
  font-weight: bold;
  font-size: 12px;
  box-shadow: 2px 2px 4px rgba(0, 0, 0, 0.4);
  color: #321300;
  cursor: pointer;
}
</style>