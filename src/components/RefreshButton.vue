<script setup>
import { config_store } from '@/store/config_store.js'
</script>

<template>
<button @click="config_store.randomizeList()">Refresh</button>
</template>

<style scoped>

</style>