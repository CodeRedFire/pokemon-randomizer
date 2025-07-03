<script setup>

</script>

<template>
  <a href="#/menu">
    <span class="material-symbols-rounded settingsButton">
      menu
    </span>
  </a>
</template>

<style scoped>
.settingsButton {
  position: absolute;
  top: 10px;
  left: 10px;
  font-size: 32px;
  cursor: pointer;
  text-decoration: none;
  color: #000;
  opacity: 0.5;
  z-index: 200;
}
.settingsButton:hover {
  opacity: 0.8;
}
</style>