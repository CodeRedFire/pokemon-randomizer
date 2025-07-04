import { reactive } from 'vue'
import {MaxValues} from "@/enums/max_values.js";
import {i18n} from "@/i18n/i18n.js";
import {DisplayedDataOptions} from "@/enums/displayed_data.js";

export const config_store = reactive({
    nbrPokemon: 3,
    pkmnList: [],
    randomizeList() {
        let selectedPokemonIdsList = []

        for (let i = 0; i < this.nbrPokemon; i++) {
            let newValue = 0;
            while (newValue === 0 || selectedPokemonIdsList.includes(newValue)) {
                newValue = Math.floor(Math.random() * MaxValues.Pokedex)
            }
            selectedPokemonIdsList.push(newValue)
        }
        this.pkmnList = selectedPokemonIdsList;
    },
    displayedLanguage: i18n.global.locale,
    displayedInformations: DisplayedDataOptions.all,
})