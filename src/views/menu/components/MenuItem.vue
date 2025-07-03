<script setup>
const props = defineProps({
  link: String,
  title: String,
  img: String,
})
</script>

<template>
  <a :href="props.link">
    <div class="itemContainer">
      <img :src="props.img"/>
      <div class="itemTitle">{{props.title}}</div>
    </div>
  </a>
</template>

<style scoped>
a {
  text-decoration: none;
}

.itemContainer {
  padding: 20px 12px;
  margin: 20px;
  border-radius: 8px;
  text-decoration: none;
  color: #eeeeee;
  font-family: 'Roboto', sans-serif;
  letter-spacing: 0.5px;
  box-shadow: 2px 2px 4px rgba(0, 0, 0, 0.4);
  position: relative;
  width: 200px;
  height: 200px;
}

.itemTitle {
  position: absolute;
  left: 0px;
  bottom: 0px;
  width: 100%;
  padding: 20px 0px 12px 0px;
  background: linear-gradient(transparent 0%, #222222 40%);
  font-size: 14px;
  font-weight: 500;
  text-align: center;
}
img {
  position: absolute;
  top: 0px;
  left: 0px;
  width: 100%;
  height: 100%;
}

</style>