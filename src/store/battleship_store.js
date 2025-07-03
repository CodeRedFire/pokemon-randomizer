import { reactive } from 'vue'
import {MaxValues} from "@/enums/max_values.js";
import {i18n} from "@/i18n/i18n.js";
import {DisplayedDataOptions} from "@/enums/displayed_data.js";
import PkmnTile from "@/views/battleship/components/PkmnTile.vue";
import {TileData} from "@/models/TileData.js";

export const battleship_store = reactive({
    nbrColumns: 6,
    nbrRows: 6,
    tileDatas: [],
    randomizeList() {
        let selectedPokemonIdsList = []
        let battleshipPkmnList = []

        for (let i = 0; i < this.nbrRows; i++) {
            let columnData =[]
            for (let j = 0; j < this.nbrColumns; j++) {
                let newValue = 0;
                while (newValue === 0 || selectedPokemonIdsList.includes(newValue)) {
                    newValue = Math.floor(Math.random() * MaxValues.Pokedex)
                }
                columnData.push(new TileData(newValue, "#829cbc", i, j))
                selectedPokemonIdsList.push(newValue)
            }
            battleshipPkmnList.push(columnData)
        }
        this.tileDatas = battleshipPkmnList;
    }
})