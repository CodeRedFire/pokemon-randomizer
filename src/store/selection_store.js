import { reactive } from 'vue'

export const selection_store = reactive({
    currentIndex: 0,
    selectedPkmnList: [0,0,0,0,0,0],
    selectedPkmn(pkmnNum) {
        this.selectedPkmnList[this.currentIndex] = pkmnNum;
    },
    next() {
        this.currentIndex = (this.currentIndex + 1) % 6;
    },
    selectIndex(index) {
        this.currentIndex = index;
    },
    reset() {
        this.currentIndex = 0;
        this.selectedPkmnList = [0,0,0,0,0,0];
    }
})