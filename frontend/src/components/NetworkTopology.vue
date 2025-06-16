<template>
  <div ref="container" class="relative w-full h-[500px] border border-gray-700 rounded-lg overflow-hidden select-none">
    <!-- Connections -->
    <svg class="absolute top-0 left-0 w-full h-full pointer-events-none">
      <line
        v-for="c in connections"
        :key="c.id"
        :x1="deviceById(c.source)?.x + iconHalf"
        :y1="deviceById(c.source)?.y + iconHalf"
        :x2="deviceById(c.target)?.x + iconHalf"
        :y2="deviceById(c.target)?.y + iconHalf"
        stroke="gray"
        stroke-width="2"
        :stroke-dasharray="c.status === 'inactive' ? '6 4' : '0'"
      />
    </svg>

    <!-- Devices -->
    <div
      v-for="d in devices"
      :key="d.id"
      class="absolute cursor-pointer flex flex-col items-center"
      :style="{ left: d.x + 'px', top: d.y + 'px' }"
      @mousedown="startDrag($event, d)"
      @click.stop="emit('device-select', d)"
    >
      <DeviceIcon :type="d.type" :status="d.status" />
      <span :class="['text-xs mt-1', selectedDevice?.id === d.id ? 'text-blue-400' : 'text-gray-300']">{{ d.name }}</span>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import DeviceIcon from './DeviceIcon.vue'

const props = defineProps({
  devices: Array,
  connections: Array,
  selectedDevice: Object,
})
const emit = defineEmits(['device-select', 'device-move'])

const container = ref(null)
const iconHalf = 12 // half of 24px icon
let dragging = null

const deviceById = (id) => props.devices.find((d) => d.id === id)

const startDrag = (e, device) => {
  dragging = {
    id: device.id,
    offsetX: e.offsetX,
    offsetY: e.offsetY,
  }
  window.addEventListener('mousemove', onDrag)
  window.addEventListener('mouseup', endDrag)
}

const onDrag = (e) => {
  if (!dragging) return
  const rect = container.value.getBoundingClientRect()
  const x = e.clientX - rect.left - dragging.offsetX
  const y = e.clientY - rect.top - dragging.offsetY
  emit('device-move', dragging.id, x, y)
}

const endDrag = () => {
  dragging = null
  window.removeEventListener('mousemove', onDrag)
  window.removeEventListener('mouseup', endDrag)
}

const STORAGE_KEY = 'network_topology_positions';

// Load saved positions from localStorage
function loadSavedPositions() {
  const saved = localStorage.getItem(STORAGE_KEY);
  return saved ? JSON.parse(saved) : {};
}

// Save positions to localStorage
function savePositions(positions) {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(positions));
}
</script>
