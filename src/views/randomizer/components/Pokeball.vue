<script setup>
import {ref} from "vue";
import {selection_store} from "@/store/selection_store.js";

const props = defineProps({
  pkmnNum: Number,
  index: Number,
})

</script>

<template>
    <div class="pokeball" :class="{'selected': index == selection_store.currentIndex}"
    @click="selection_store.selectIndex(index)">
      <div class="center-button"></div>
      <img class="sprite" v-if="props.pkmnNum != 0"
           :src="`/pokemon-randomizer/assets/sprites/pkmn/sprite_${props.pkmnNum}.png`"/>
    </div>

</template>

<style scoped>

.pokeball {
  position: relative;
  width: 60px;
  height: 60px;
  border: 2px solid black;
  border-radius: 50%;
  background: linear-gradient(to bottom, red 50%, white 50%);
  box-shadow: inset 0 0 0 2px black;
  margin: 0 4px;
  cursor: pointer;
}
.selected {
  border: 2px solid #299c6e !important;
  box-shadow: inset 0 0 0 2px #299c6e;
}

.sprite {
  position: absolute;
  top: 0px;
  left: 0px;
  width: 60px;
  height: 60px;
  z-index: 500;
}

.pokeball::before {
  content: '';
  position: absolute;
  top: 50%;
  left: 0;
  transform: translateY(-50%);
  width: 100%;
  height: 4px;
  background: black;
}

.center-button {
  position: absolute;
  top: 50%;
  left: 50%;
  width: 12px;
  height: 12px;
  background: white;
  border: 4px solid black;
  border-radius: 50%;
  transform: translate(-50%, -50%);
  z-index: 2;
}

.center-button::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  width: 4px;
  height: 4px;
  background: lightgray;
  border-radius: 50%;
  transform: translate(-50%, -50%);
}
</style>