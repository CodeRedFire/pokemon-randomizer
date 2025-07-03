<script setup lang="ts">

import {battleship_store} from "@/store/battleship_store.js";
import PkmnTile from "@/views/battleship/components/PkmnTile.vue";
</script>

<template>
  <div class="grid-container">
    <div class="grid">
      <div>
        <div v-for="row in battleship_store.tileDatas" class="row">
          <PkmnTile v-for="tileData in row" :tileData="tileData"/>
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