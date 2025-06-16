<template>
  <div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 gap-6">
    <div
      v-for="device in devices"
      :key="device.id"
      class="bg-gray-800 border border-gray-600 rounded-lg p-4"
    >
      <div class="flex items-center gap-3 mb-4">
        <DeviceIcon :type="device.type" :status="device.status" />
        <div>
          <h3 class="text-white font-medium">{{ device.name }}</h3>
          <p class="text-gray-400 text-sm">{{ device.ip }}</p>
        </div>
      </div>
      <div class="space-y-3">
        <div>
          <div class="flex justify-between text-sm mb-1">
            <span class="text-gray-400">CPU</span>
            <span class="text-white">{{ device.cpu }}%</span>
          </div>
          <div class="w-full bg-gray-600 rounded-full h-2">
            <div
              class="bg-blue-500 h-2 rounded-full transition-all duration-300"
              :style="{ width: `${device.cpu}%` }"
            />
          </div>
        </div>
        <div>
          <div class="flex justify-between text-sm mb-1">
            <span class="text-gray-400">Memory</span>
            <span class="text-white">{{ Math.min(device.memory, 100) }}%</span>
          </div>
          <div class="w-full bg-gray-600 rounded-full h-2">
            <div
              class="bg-green-500 h-2 rounded-full transition-all duration-300"
              :style="{ width: `${Math.min(device.memory, 100)}%` }"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  devices: Array
});
</script>
