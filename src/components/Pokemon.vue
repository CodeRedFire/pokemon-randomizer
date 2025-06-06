<script setup>
import {onMounted, ref} from "vue";
import {ColorType} from "@/enums/color_type.js";

const props = defineProps({
  numPokedex: Number,
})

const jsonData = ref(null);
const error = ref(null);
const imgUrl = ref(null);
const type1Color = ref("transparent");
const type2Color = ref("transparent");

onMounted(async () => {
  try {
    const response = await fetch(`/assets/data/pkmn/data_${props.numPokedex}.json`);

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    jsonData.value = await response.json();
    console.log('JSON Data:', jsonData.value);
    type1Color.value = ColorType[jsonData.value["types"][0]];
    if (jsonData.value["types"][1]) {
      type2Color.value = ColorType[jsonData.value["types"][1]];
    } else {
      type2Color.value = ColorType[jsonData.value["types"][0]];
    }
    imgUrl.value = `/assets/sprites/pkmn/sprite_${props.numPokedex}.png`;
  } catch (err) {
    error.value = err.message;
    console.error('Failed to load JSON:', err);
    imgUrl.value = null;
  }
});
</script>

<template>
  <div class="container">
    <div class="numero">#{{ numPokedex }}</div>
    <img class="sprite" v-if="imgUrl" :src="imgUrl"/>
    <div class="name" v-if="jsonData">{{
        jsonData["names"].filter((n) => n["language"]["name"] == "fr")[0]["name"]
      }}
    </div>
  </div>
</template>

<style scoped>
.container {
  --largeur: 280px;
  --hauteur: 280px;
  width: var(--largeur);
  height: var(--hauteur);
  position: relative;
  background: linear-gradient(45deg, rgba(0, 0, 0, 0.8) 0%, rgba(0, 0, 0, 0.5) 10%, transparent 40%, transparent 60%, rgba(0, 0, 0, 0.5) 90%, rgba(0, 0, 0, 0.8) 100%),
  linear-gradient(135deg, v-bind('type1Color') 30%, v-bind('type2Color') 70%);
  font-family: "Roboto", sans-serif;
  text-transform: uppercase;
  color: white;
  text-shadow: 2px 2px 2px #000,
  2px 0px 2px #000,
  -2px 2px 2px #000,
  -2px 0px 2px #000,
  -2px -2px 2px #000,
  0px -2px 2px #000,
  2px -2px 2px #000,
  0px 2px 2px #000;
  letter-spacing: 1px;
  margin: 2px;
}

.numero {
  position: absolute;
  top: 8px;
  left: 8px;
  font-weight: bold;
  font-size: 32px;
}

.sprite {
  position: absolute;
  top: calc(var(--hauteur) / 10);
  left: calc(var(--largeur) / 10);
  width: calc(var(--largeur) / 5 * 4);
  height: calc(var(--hauteur) / 5 * 4);
  filter: drop-shadow(4px 4px 2px rgba(0, 0, 0, 0.5));
}

.name {
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  font-size: 32px;
  font-weight: bold;
  text-align: center;
  padding-bottom: 8px;
}
</style>
