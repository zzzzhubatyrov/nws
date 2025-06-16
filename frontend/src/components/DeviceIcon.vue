<template>
  <component
    :is="iconComp"
    :class="[
      'w-6 h-6',
      typeColor,
    ]"
  />
</template>

<script setup>
import { computed } from 'vue'
import { Server as ServerIcon, Router as RouterIcon, Activity as ActivityIcon } from 'lucide-vue-next'

const props = defineProps({
  type: String,
  status: String,
})

const iconComp = computed(() => {
  if (props.type === 'router') return RouterIcon
  if (props.type === 'server') return ServerIcon
  return ActivityIcon
})

const typeColor = computed(() => {
  return props.status === 'online' ? 'text-green-400' : 'text-red-400'
})
</script>
