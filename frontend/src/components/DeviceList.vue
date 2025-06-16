<template>
  <div class="bg-gray-800 border border-gray-600 rounded-lg p-4">
    <div class="flex justify-between items-center mb-4">
      <h2 class="text-white font-medium">Devices</h2>
      <div class="flex gap-3">
        <button @click="$emit('add-device')" class="text-sm text-blue-400 hover:text-blue-300">+ Add</button>
        <button @click="$emit('auto-device')" class="text-sm text-green-500 hover:text-green-400">Auto-Generate</button>
      </div>
    </div>
    <table class="w-full text-sm text-left">
      <thead class="text-gray-400">
        <tr>
          <th>Name</th>
          <th>IP</th>
          <th>Status</th>
          <th class="w-16"></th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="d in devices"
          :key="d.id"
          :class="selectedDevice?.id === d.id ? 'bg-gray-700' : ''"
          class="cursor-pointer hover:bg-gray-700"
          @click="$emit('device-select', d)"
        >
          <td class="py-1">{{ d.name }}</td>
          <td>{{ d.ip }}</td>
          <td>{{ d.status }}</td>
          <td>
            <button @click.stop="$emit('edit-device', d)" class="text-xs text-blue-500 hover:text-blue-400 mr-2">Edit</button>
            <button @click.stop="$emit('delete-device', d.id)" class="text-xs text-red-500 hover:text-red-400">Del</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
const props = defineProps({
  devices: Array,
  selectedDevice: Object,
})
const emit = defineEmits(['device-select', 'add-device', 'edit-device', 'delete-device', 'auto-device'])
</script>
