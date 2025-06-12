<script setup>
import {onMounted, ref} from "vue";
import {ColorType} from "@/enums/color_type.js";
import {config_store} from "@/store/config_store.js";
import {Languages} from "@/enums/languages.js";
import {convert} from "hangul-romanization";
import { transliterate as tr } from 'transliteration';
import {DisplayedDataOptions} from "@/enums/displayed_data.js";

const props = defineProps({
  numPokedex: Number,
})

const jsonData = ref(null);
const error = ref(null);
const imgUrl = ref(null);
const type1Color = ref("transparent");
const type2Color = ref("transparent");

const isAllShow = ref(false);
const isImgShow = ref(false);
const isNumberShow = ref(false);
const isTypeShow = ref(false);
const isNameShow = ref(false);
const isBestStatShow = ref(false);
const isWorstStatShow = ref(false);

onMounted(async () => {
  // Manage display of informations
  isAllShow.value = config_store.displayedInformations === DisplayedDataOptions.all;
  isImgShow.value = config_store.displayedInformations === DisplayedDataOptions.all;
  isNumberShow.value = config_store.displayedInformations === DisplayedDataOptions.all || config_store.displayedInformations === DisplayedDataOptions.pokedexNumber;
  isTypeShow.value = config_store.displayedInformations === DisplayedDataOptions.all || config_store.displayedInformations === DisplayedDataOptions.type;
  isNameShow.value = config_store.displayedInformations === DisplayedDataOptions.all || config_store.displayedInformations === DisplayedDataOptions.nameOnly;
  isBestStatShow.value = config_store.displayedInformations === DisplayedDataOptions.bestStat;
  isWorstStatShow.value = config_store.displayedInformations === DisplayedDataOptions.worstStat;

  try {
    imgUrl.value = `/assets/sprites/pkmn/sprite_${props.numPokedex}.png`;

    const response = await fetch(`/assets/data/pkmn/data_${props.numPokedex}.json`);

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    jsonData.value = await response.json();
    if (isTypeShow.value) {
      type1Color.value = ColorType[jsonData.value["types"][0]];
      if (jsonData.value["types"][1]) {
        type2Color.value = ColorType[jsonData.value["types"][1]];
      } else {
        type2Color.value = ColorType[jsonData.value["types"][0]];
      }
    }
  } catch (err) {
    error.value = err.message;
    console.error('Failed to load JSON:', err);
    imgUrl.value = null;
  }
});

function displayName(jsonData) {
  let lang = config_store.displayedLanguage

  // Select the right store language name for the pokemon
  if (lang == Languages.koreanLatin || lang == Languages.koreanHangul) {
    lang = "ko";
  }
  if (lang == Languages.chineseLatin) {
    lang = "zhHant";
  }

  let name = jsonData["names"].filter((n) => n["language"]["name"] == lang)[0]["name"];

  // Execute a convertion to latin alphabet
  if (config_store.displayedLanguage == Languages.koreanLatin) {
    return convert(name);
  }
  if (config_store.displayedLanguage == Languages.chineseLatin) {
    return tr(name);
  }

  return name;
}

function displayBestStat(jsonData) {
  let objetAvecPlusGrandeValeur = null;
  let plusGrandeValeur = -Infinity;

  for (const item of jsonData.stats) {
    if (item.value > plusGrandeValeur) {
      plusGrandeValeur = item.value;
      objetAvecPlusGrandeValeur = item;
    }
  }

  if (!objetAvecPlusGrandeValeur) {
    console.log("Le tableau est vide ou ne contient pas d'objets valides.");
  }

  let statName = {
    "hp": "HP",
    "attack": "Atk",
    "defense": "Def",
    "special-attack": "SpA",
    "special-defense": "SpD",
    "speed": "Spe",
  };

  return `${statName[objetAvecPlusGrandeValeur.name]} : ${objetAvecPlusGrandeValeur.value}`
}

function displayWorstStat(jsonData) {
  let smallestObject = null;
  let smallestValue = Infinity;

  for (const item of jsonData.stats) {
    if (item.value < smallestValue) {
      smallestValue = item.value;
      smallestObject = item;
    }
  }

  if (!smallestObject) {
    console.log("Le tableau est vide ou ne contient pas d'objets valides.");
  }

  let statName = {
    "hp": "HP",
    "attack": "Atk",
    "defense": "Def",
    "special-attack": "SpA",
    "special-defense": "SpD",
    "speed": "Spe",
  };

  return `${statName[smallestObject.name]} : ${smallestObject.value}`
}
</script>

<template>
  <div class="container">
    <div class="numero" v-if="isImgShow">#{{ numPokedex }}</div>
    <img class="sprite" v-if="imgUrl && isImgShow" :src="imgUrl"/>
    <div class="sprite centeredText" v-if="jsonData && isNumberShow && !isAllShow">
      #{{ numPokedex }}
    </div>
    <div class="sprite centeredText" v-if="jsonData && isBestStatShow">
      {{displayBestStat(jsonData)}}
    </div>
    <div class="sprite centeredText" v-if="jsonData && isWorstStatShow">
      {{displayWorstStat(jsonData)}}
    </div>
    <div class="name" v-if="jsonData && isNameShow">{{
        displayName(jsonData)
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
  margin: 4px;
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
  font-size: 28px;
  font-weight: bold;
  text-align: center;
  padding-bottom: 8px;
}

.centeredText {
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 40px;
}
</style>
