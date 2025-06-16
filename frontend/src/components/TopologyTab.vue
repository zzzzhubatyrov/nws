<template>
  <div class="topology-container bg-gray-900 flex flex-col gap-6 p-4">
    <div class="flex flex-grow gap-4 min-h-0">
      <!-- Network Graph -->
      <div ref="graphContainer" class="flex-grow relative bg-gray-800 rounded-lg shadow-lg overflow-hidden">
        <svg class="w-full h-full">
          <!-- Grid Background -->
          <pattern id="grid" width="50" height="50" patternUnits="userSpaceOnUse">
            <path d="M 50 0 L 0 0 0 50" fill="none" stroke="currentColor" stroke-width="0.5" class="text-gray-700/30" />
          </pattern>
          <rect width="100%" height="100%" fill="url(#grid)" />

          <!-- Connections -->
          <line
            v-for="conn in connections"
            :key="conn.id"
            :x1="getDeviceById(conn.source)?.x"
            :y1="getDeviceById(conn.source)?.y"
            :x2="getDeviceById(conn.target)?.x"
            :y2="getDeviceById(conn.target)?.y"
            :class="[
              'stroke-2',
              conn.status === 'active' ? 'stroke-green-500/50' : 'stroke-red-500/50'
            ]"
          />

          <!-- Devices -->
          <g
            v-for="device in devices"
            :key="device.id"
            :transform="'translate(' + device.x + ',' + device.y + ')'"
            @mousedown="startDrag(device, $event)"
            @click.stop="selectDevice(device)"
            class="cursor-move"
          >
            <!-- Device Icon Background -->
            <circle
              :class="[
                'transition-colors duration-200',
                device.id === selectedDevice?.id ? 'fill-blue-600 stroke-blue-400' : 'fill-gray-700 stroke-gray-600',
                'hover:fill-gray-600 stroke-2'
              ]"
              r="25"
            />

            <!-- Device Icon -->
            <text
              :class="[
                'text-2xl fill-current transition-colors duration-200',
                device.status === 'online' ? 'text-green-400' : 'text-red-400'
              ]"
              text-anchor="middle"
              dominant-baseline="middle"
            >
              {{ getDeviceIcon(device.type) }}
            </text>

            <!-- Device Label -->
            <text
              y="40"
              text-anchor="middle"
              class="fill-gray-300 text-sm font-medium"
            >
              {{ device.name }}
            </text>

            <!-- Status Indicator -->
            <circle
              cx="20"
              cy="-15"
              r="5"
              :class="[
                'transition-colors duration-200',
                device.status === 'online' ? 'fill-green-500' : 'fill-red-500'
              ]"
            />
          </g>
        </svg>
      </div>

      <!-- Device Details Sidebar -->
      <div v-if="selectedDevice" class="w-80 bg-gray-800 rounded-lg shadow-lg overflow-hidden">
        <div class="p-4 bg-gray-800 border-b border-gray-700">
          <div class="flex justify-between items-center">
            <h3 class="text-xl font-semibold text-white">{{ selectedDevice.name }}</h3>
            <button
              @click="selectedDevice = null"
              class="text-gray-400 hover:text-white transition-colors duration-200"
            >
              ✕
            </button>
          </div>
        </div>

        <div class="p-4 space-y-4">
          <div class="bg-gray-700/50 rounded-lg p-4">
            <div class="grid grid-cols-2 gap-3 text-sm">
              <!-- Status with toggle -->
              <div class="text-gray-400">Status</div>
              <div class="flex items-center gap-2">
                <div
                  class="w-3 h-3 rounded-full"
                  :class="selectedDevice.status === 'online' ? 'bg-green-500' : 'bg-red-500'"
                ></div>
                <button
                  @click="toggleDeviceStatus(selectedDevice)"
                  :class="selectedDevice.status === 'online' ? 'text-green-400' : 'text-red-400'"
                  class="hover:underline"
                >
                  {{ selectedDevice.status }}
                </button>
              </div>

              <div class="text-gray-400">IP Address</div>
              <div class="text-white">{{ selectedDevice.ip }}</div>
              <div class="text-gray-400">Type</div>
              <div class="text-white">{{ selectedDevice.type }}</div>
              <div class="text-gray-400">Vendor</div>
              <div class="text-white">{{ selectedDevice.vendor }}</div>
              <div class="text-gray-400">Model</div>
              <div class="text-white">{{ selectedDevice.model }}</div>
            </div>
          </div>

          <!-- Performance section only for online devices -->
          <div v-if="selectedDevice.status === 'online'" class="bg-gray-700/50 rounded-lg p-4">
            <h4 class="text-sm font-medium text-gray-400 mb-3">Performance</h4>
            <div class="space-y-3">
              <div>
                <div class="flex justify-between text-sm mb-1">
                  <span class="text-gray-400">CPU</span>
                  <span :class="selectedDevice.cpu > 80 ? 'text-red-400' : 'text-white'">
                    {{ Math.round(selectedDevice.cpu) }}%
                  </span>
                </div>
                <div class="h-2 bg-gray-600/50 rounded overflow-hidden">
                  <div
                    class="h-full rounded transition-all duration-300 ease-in-out"
                    :class="selectedDevice.cpu > 80 ? 'bg-red-500' : 'bg-blue-500'"
                    :style="{ width: selectedDevice.cpu + '%' }"
                  />
                </div>
              </div>
              <div>
                <div class="flex justify-between text-sm mb-1">
                  <span class="text-gray-400">Memory</span>
                  <span :class="selectedDevice.memory > 80 ? 'text-red-400' : 'text-white'">
                    {{ Math.round(selectedDevice.memory) }}%
                  </span>
                </div>
                <div class="h-2 bg-gray-600/50 rounded overflow-hidden">
                  <div
                    class="h-full rounded transition-all duration-300 ease-in-out"
                    :class="selectedDevice.memory > 80 ? 'bg-red-500' : 'bg-green-500'"
                    :style="{ width: selectedDevice.memory + '%' }"
                  />
                </div>
              </div>
            </div>
          </div>

          <!-- Offline message when device is not online -->
          <div v-else class="bg-gray-700/50 rounded-lg p-4">
            <div class="text-center text-gray-400">
              <div class="text-4xl mb-2">📡</div>
              <p>Device is offline</p>
              <p class="text-sm">No metrics available</p>
            </div>
          </div>

          <div class="flex gap-2">
            <button
              @click="$emit('edit-device', selectedDevice)"
              class="flex-1 bg-blue-600 hover:bg-blue-700 text-white rounded-lg px-4 py-2 transition-colors duration-200 font-medium"
            >
              Edit
            </button>
            <button
              @click="$emit('delete-device', selectedDevice.id)"
              class="flex-1 bg-red-600 hover:bg-red-700 text-white rounded-lg px-4 py-2 transition-colors duration-200 font-medium"
            >
              Delete
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Alerts Panel with fixed height -->
    <div class="bg-gray-800 rounded-lg shadow-lg overflow-hidden h-64">
      <AlertsPanel 
        :alerts="alerts" 
        @acknowledge="handleAlertAcknowledge"
        @resolve="handleAlertResolve"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, watch } from 'vue';
import AlertsPanel from './AlertsPanel.vue';

const props = defineProps({
  devices: {
    type: Array,
    required: true,
    default: () => []
  },
  connections: {
    type: Array,
    required: true,
    default: () => []
  },
  alerts: {
    type: Array,
    required: true,
    default: () => []
  }
});

const emit = defineEmits(['device-select', 'device-move', 'edit-device', 'delete-device', 'update-device', 'alertAcknowledge', 'alertResolve']);

const graphContainer = ref(null);
const selectedDevice = ref(null);
const isDragging = ref(false);
const draggedDevice = ref(null);
const dragOffset = ref({ x: 0, y: 0 });

// Device icons mapping
const deviceIcons = {
  router: '🌐',
  switch: '🔌',
  server: '💻',
  firewall: '🛡️',
  workstation: '🖥️',
  default: '📱'
};

// Get device icon based on type
const getDeviceIcon = (type) => {
  return deviceIcons[type] || deviceIcons.default;
};

// Get device by ID helper
const getDeviceById = (id) => {
  return props.devices.find(d => d.id === id);
};

// Toggle device status
const toggleDeviceStatus = (device) => {
  const newStatus = device.status === 'online' ? 'offline' : 'online';
  emit('update-device', { ...device, status: newStatus });
};

// Start dragging a device
const startDrag = (device, event) => {
  event.preventDefault();
  isDragging.value = true;
  draggedDevice.value = device;

  // Calculate offset from device center
  const rect = graphContainer.value.getBoundingClientRect();
  dragOffset.value = {
    x: device.x - (event.clientX - rect.left),
    y: device.y - (event.clientY - rect.top)
  };

  // Add event listeners
  window.addEventListener('mousemove', handleDrag);
  window.addEventListener('mouseup', stopDrag);
};

// Handle device dragging
const handleDrag = (event) => {
  if (!isDragging.value || !draggedDevice.value) return;

  const rect = graphContainer.value.getBoundingClientRect();
  const newX = event.clientX - rect.left + dragOffset.value.x;
  const newY = event.clientY - rect.top + dragOffset.value.y;

  // Update device position
  draggedDevice.value.x = Math.max(25, Math.min(rect.width - 25, newX));
  draggedDevice.value.y = Math.max(25, Math.min(rect.height - 25, newY));

  // Emit move event and save positions
  emit('device-move', draggedDevice.value.id, draggedDevice.value.x, draggedDevice.value.y);
  saveNodePositions();
};

// Stop dragging
const stopDrag = () => {
  isDragging.value = false;
  draggedDevice.value = null;
  window.removeEventListener('mousemove', handleDrag);
  window.removeEventListener('mouseup', stopDrag);
};

// Select a device
const selectDevice = (device) => {
  selectedDevice.value = device;
  emit('device-select', device);
};

// Update selected device when devices prop changes
watch(() => props.devices, (newDevices) => {
  if (selectedDevice.value) {
    const updatedDevice = newDevices.find(d => d.id === selectedDevice.value.id);
    if (updatedDevice) {
      selectedDevice.value = updatedDevice;
    }
  }
}, { deep: true });

// Save node positions to localStorage
const saveNodePositions = () => {
  const positions = {};
  props.devices.forEach(device => {
    positions[device.id] = { x: device.x, y: device.y };
  });
  localStorage.setItem('nodePositions', JSON.stringify(positions));
};

// Load node positions from localStorage
const loadNodePositions = () => {
  const savedPositions = localStorage.getItem('nodePositions');
  if (savedPositions) {
    const positions = JSON.parse(savedPositions);
    props.devices.forEach(device => {
      if (positions[device.id]) {
        device.x = positions[device.id].x;
        device.y = positions[device.id].y;
      }
    });
  }
};

onMounted(() => {
  loadNodePositions();
});

// Clean up event listeners
onUnmounted(() => {
  window.removeEventListener('mousemove', handleDrag);
  window.removeEventListener('mouseup', stopDrag);
});

// Add alerts handling
const alerts = ref([]);

// Handle alert acknowledgement
const handleAlertAcknowledge = (alertId) => {
  const alert = alerts.value.find(a => a.id === alertId);
  if (alert) {
    alert.status = 'acknowledged';
  }
};

// Handle alert resolution
const handleAlertResolve = (alertId) => {
  const alert = alerts.value.find(a => a.id === alertId);
  if (alert) {
    alert.status = 'resolved';
  }
};

// Watch for alerts from parent component
watch(() => props.alerts, (newAlerts) => {
  alerts.value = newAlerts;
}, { deep: true });
</script>

<style scoped>
.topology-container {
  height: calc(100vh - 12rem);
  min-height: 0;
  display: flex;
  flex-direction: column;
}

/* Prevent text selection during drag */
svg text {
  user-select: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
}

/* Make sure SVG lines are visible */
line {
  stroke-width: 2px;
  pointer-events: none;
}

/* Smooth transitions */
.transition-all {
  transition-property: all;
  transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
  transition-duration: 150ms;
}
</style>
