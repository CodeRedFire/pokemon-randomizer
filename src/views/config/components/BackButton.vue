<script setup>

function goBack() {
  console.log("go back");
  window.history.back();
}
</script>

<template>
  <div @click="goBack()">
    <span class="material-symbols-rounded backButton">
arrow_back
</span>
  </div>
</template>

<style scoped>
.backButton {
  font-size: 32px;
  cursor: pointer;
  text-decoration: none;
  color: #000;
  opacity: 0.5;
}

.backButton:hover {
  opacity: 0.8;
}
</style>