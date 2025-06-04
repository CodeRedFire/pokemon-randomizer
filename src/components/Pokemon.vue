<script setup>
import {onMounted, ref} from "vue";

const props = defineProps({
  numPokedex: Number,
})

const jsonData = ref(null);
const error = ref(null);
const imgUrl = ref(null);

onMounted(async () => {
  try {
    const response = await fetch(`/assets/data/pkmn/data_${props.numPokedex}.json`);

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    jsonData.value = await response.json();
    console.log('JSON Data:', jsonData.value);
    imgUrl.value = `/assets/sprites/pkmn/sprite_${props.numPokedex}.png`;
  } catch (err) {
    error.value = err.message;
    console.error('Failed to load JSON:', err);
    imgUrl.value = null;
  }
});
</script>

<template>
  <div>{{numPokedex}}</div>
  <div v-if="jsonData">
    <img v-if="imgUrl" :src="imgUrl"/>
    <pre>{{ jsonData["names"].filter((n) => n["language"]["name"] == "fr")[0]["name"] }}</pre>
  </div>
  <div v-if="error" style="color: red;">
    <p>Error loading data: {{ error }}</p>
  </div>
</template>

<style scoped>
</style>
