<script setup>
import {ref} from "vue";
import {TileData} from "@/models/TileData.js";

const props = defineProps({
  tileData: TileData,
})

const bgColor = ref(props.tileData.color);

function toogleStatus() {
  if (bgColor.value === "#687FE5") {
    bgColor.value = "#829cbc";
  } else {
    bgColor.value = "#687FE5";
  }
}

function toogleStatusReverse(e) {
  if (bgColor.value === "#FF4F0F") {
    bgColor.value = "#829cbc";
  } else {
    bgColor.value = "#FF4F0F";
  }
  e.preventDefault();
}
</script>

<template>
  <div @click="toogleStatus()" @contextmenu="toogleStatusReverse">
    <img class="sprite"
         :src="`/pokemon-randomizer/assets/sprites/pkmn/sprite_${props.tileData.nbrPkmn}.png`"/>
  </div>
</template>

<style scoped>
div {
  margin: 8px;
  background-color: v-bind('bgColor');
  border-radius: 8px;
  box-shadow: 4px 4px 6px 0 rgba(0, 0, 0, 0.2);
}
div:hover {
  cursor: pointer;
}
img {
  width: 120px;
}
</style>