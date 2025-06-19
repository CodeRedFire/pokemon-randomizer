<script setup>
import {battleship_store} from "@/store/battleship_store.js";
import {ref} from "vue";

const props = defineProps({
  nbrPkmn: Number,
})

const bgColor = ref("transparent");

function toogleStatus() {
  if (bgColor.value === "transparent") {
    bgColor.value = "#687FE5";
  } else if (bgColor.value === "#687FE5") {
    bgColor.value = "#FF4F0F";
  } else if (bgColor.value === "#FF4F0F") {
    bgColor.value = "transparent";
  }
}

function toogleStatusReverse(e) {
  if (bgColor.value === "#687FE5") {
    bgColor.value = "transparent";
  } else if (bgColor.value === "#FF4F0F") {
    bgColor.value = "#687FE5";
  }else if (bgColor.value === "transparent") {
    bgColor.value = "#FF4F0F";
  }
  e.preventDefault();
}
</script>

<template>
  <div @click="toogleStatus()" @contextmenu="toogleStatusReverse">
    <img class="sprite"
         :src="`/pokemon-randomizer/assets/sprites/pkmn/sprite_${props.nbrPkmn}.png`"/>
  </div>
</template>

<style scoped>
div {
  border: 1px solid rgba(0, 0, 0, 0.4);
  padding: 8px;
  background-color: v-bind('bgColor');
}
div:hover {
  cursor: pointer;
}
img {
  width: 120px;
}
</style>