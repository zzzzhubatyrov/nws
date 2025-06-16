<template>
  <div class="bg-gray-800 border border-gray-600 rounded-lg p-4">
    <div class="flex justify-between items-center mb-4">
      <h2 class="text-white font-medium">Connections</h2>
      <button @click="$emit('add')" class="text-sm text-blue-400">+ Add</button>
    </div>
    <table class="w-full text-sm text-left">
      <thead class="text-gray-400">
        <tr>
          <th>Name</th>
          <th>Source</th>
          <th>Target</th>
          <th>Status</th>
          <th class="w-16"></th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="c in connections" :key="c.id" class="hover:bg-gray-700">
          <td class="py-1">{{ c.name }}</td>
          <td>{{ deviceName(c.source) }}</td>
          <td>{{ deviceName(c.target) }}</td>
          <td>{{ c.status }}</td>
          <td>
            <button @click.stop="$emit('edit', c)" class="text-xs text-blue-400 mr-2">Edit</button>
            <button @click.stop="$emit('delete', c.id)" class="text-xs text-red-400">Del</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
const props = defineProps({
  connections: Array,
  devices: Array,
})
const emit = defineEmits(['add', 'edit', 'delete'])

const deviceName = (id) => {
  const d = props.devices.find((x) => x.id === id)
  return d ? d.name : id
}
</script>
