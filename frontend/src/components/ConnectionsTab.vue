<template>
  <ConnectionList
    :connections="connections"
    :devices="devices"
    @add="() => emit('add-connection')"
    @edit="(c) => emit('edit-connection', c)"
    @delete="(id) => emit('delete-connection', id)"
  />
</template>

<script setup>
import ConnectionList from './ConnectionList.vue'
const props = defineProps({
  connections: Array,
  devices: Array
});

const emit = defineEmits(['add-connection', 'edit-connection', 'delete-connection']);
</script>
