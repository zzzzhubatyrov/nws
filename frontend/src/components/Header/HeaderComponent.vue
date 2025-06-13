<script setup>
import { useRoute } from 'vue-router'

const route = useRoute()
console.log(route.path);

const navLeftLinks = [
  {to: '/', label: 'Dashboard'},
  {to: '/charts', label: 'Графики'},
  {to: '/', label: 'Web', activeFor: '/web'},
  {to: '/latest', label: 'Последние данные'},
]

const navCenterLink = [
  {to: '/', label: 'Мониторинг', activeFor: '/'},
  {to: '/requests', label: 'Заявки', activeFor: '/requests'},
  {to: '/config', label: 'Конфигурация', activeFor: '/config'},
]

</script>

<template>
  <header class="header">
    <div class="left hcontainer" :class="{ disable: route.path === '/inventory' || route.path === '/requests' || route.path === '/config' }">
      <RouterLink
          v-for="link in navLeftLinks"
          :key="link.label"
          :to="link.to"
          class="hcontainer-item"
          :class="{ active: (link.activeFor || link.to) === route.path }"
      >
        {{ link.label }}
      </RouterLink>
    </div>

    <div class="center hcontainer">
      <RouterLink
          v-for="link in navCenterLink"
          :key="link.label"
          :to="link.to"
          class="hcontainer-item"
          :class="{ active: (link.activeFor || link.to) === route.path }"
      >
        {{ link.label }}
      </RouterLink>
    </div>

    <div class="right hcontainer"></div>
  </header>
  <slot></slot>
</template>


<style scoped>
* {
    color: #000;
}
.header {
  display: flex;
  width: 100%;
  justify-content: space-between;
}
.hcontainer {
  display: flex;
  flex-wrap: wrap;
  background-color: #fff;
  border-radius: 10px;
  padding: 10px 20px;
  gap: 10px;
  font-size: 14px;
}
.hcontainer.disable {
  visibility: hidden;
}
.hcontainer:nth-child(3) {
  min-width: 420px;
  min-height: 100%;
}

/* .hcontainer-item:nth-child(1) {
    font-weight: 600;
} */
.hcontainer-item.active {
  font-weight: 600;
}
</style>