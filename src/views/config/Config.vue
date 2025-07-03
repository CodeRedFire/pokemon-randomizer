<script setup>

import ConfigPanel from "@/views/config/components/ConfigPanel.vue";
import BackButton from "@/views/config/components/BackButton.vue";
import RandomizerConfigPanel from "@/views/config/components/RandomizerConfigPanel.vue";
import BattleshipConfigPanel from "@/views/config/components/BattleshipConfigPanel.vue";
</script>

<template>
  <div>
    <BackButton/>
  </div>
  <ConfigPanel/>
  <RandomizerConfigPanel/>
  <BattleshipConfigPanel/>
</template>

<style scoped>

</style>