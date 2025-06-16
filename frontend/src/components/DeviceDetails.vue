<template>
  <div class="bg-gray-800 border border-gray-600 rounded-lg p-4" v-if="device">
    <h2 class="text-white font-medium mb-4">{{ device.name }}</h2>
    <ul class="text-sm text-gray-300 space-y-1">
      <li><span class="text-gray-400">Type:</span> {{ device.type }}</li>
      <li><span class="text-gray-400">IP:</span> {{ device.ip }}</li>
      <li><span class="text-gray-400">Status:</span> {{ device.status }}</li>
      <li><span class="text-gray-400">CPU:</span> {{ device.cpu }}%</li>
      <li><span class="text-gray-400">Memory:</span> {{ Math.min(device.memory, 100) }}%</li>
      <li><span class="text-gray-400">Uptime:</span> {{ device.uptime }}</li>
      <li><span class="text-gray-400">Vendor:</span> {{ device.vendor }}</li>
      <li><span class="text-gray-400">Model:</span> {{ device.model }}</li>
    </ul>
    <div class="flex gap-3 mt-4">
      <button @click="$emit('edit', device)" class="text-blue-400 text-sm">Edit</button>
      <button @click="$emit('delete', device.id)" class="text-red-400 text-sm">Delete</button>
    </div>
  </div>
  <div v-else class="bg-gray-800 border border-gray-600 rounded-lg p-4 text-gray-400 text-sm text-center">
    Select a device to view details
  </div>
</template>

<script setup>
const props = defineProps({
  device: Object,
})
const emit = defineEmits(['edit', 'delete'])
</script>
