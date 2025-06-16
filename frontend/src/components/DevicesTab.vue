<template>
  <div class="space-y-6">
    <DeviceList
      :devices="devices"
      :selected-device="selectedDevice"
      @device-select="onDeviceSelect"
      @add-device="onAddDevice"
      @auto-device="onAutoDevice"
      @edit-device="onEditDevice"
      @delete-device="onDeleteDevice"
    />
  </div>
</template>

<script setup>
import DeviceList from './DeviceList.vue'

const props = defineProps({
  devices: Array,
  selectedDevice: Object
});

const emit = defineEmits([
  'device-select',
  'add-device',
  'edit-device',
  'delete-device',
  'auto-device'
]);

const onDeviceSelect = (d) => emit('device-select', d);
const onAddDevice = () => emit('add-device');
const onEditDevice = (d) => emit('edit-device', d);
const onDeleteDevice = (id) => emit('delete-device', id);
const onAutoDevice = () => emit('auto-device');
</script>
