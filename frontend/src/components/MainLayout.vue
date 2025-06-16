<template>
  <div class="min-h-screen bg-gray-900">
    <!-- Header -->
    <header class="bg-gray-800 shadow">
      <div class="max-w-7xl mx-auto py-4 px-4 sm:px-6 lg:px-8 flex justify-between items-center">
        <h1 class="text-2xl font-bold text-white">Network Monitor</h1>
        <button
          @click="logout"
          class="px-4 py-2 text-sm text-white bg-red-600 hover:bg-red-700 rounded-md"
        >
          Logout
        </button>
      </div>
    </header>

    <!-- Main Navigation -->
    <nav class="bg-gray-800 border-b border-gray-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex space-x-4">
          <button
            v-for="tab in tabs"
            :key="tab"
            @click="currentTab = tab"
            :class="[
              'px-3 py-2 text-sm font-medium',
              currentTab === tab
                ? 'text-white bg-gray-900'
                : 'text-gray-300 hover:bg-gray-700 hover:text-white'
            ]"
          >
            {{ tab }}
          </button>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
      <component
        :is="currentComponent"
        :devices="devices"
        :connections="connections"
        :alerts="alerts"
        @edit-device="handleEditDevice"
        @delete-device="handleDeleteDevice"
        @device-move="handleDeviceMove"
        @device-select="handleDeviceSelect"
        @update-device="handleDeviceUpdate"
      />
    </main>

    <!-- Device Modal -->
    <DeviceModal
      :is-open="isDeviceModalOpen"
      :device="editingDevice"
      @close="isDeviceModalOpen = false"
      @save="handleSaveDevice"
    />

    <!-- Confirm Delete Modal -->
    <ConfirmModal
      :is-open="isConfirmModalOpen"
      title="Delete Device"
      message="Are you sure you want to delete this device? This action cannot be undone."
      @confirm="confirmDelete"
      @cancel="isConfirmModalOpen = false"
    />
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';
import TopologyTab from './TopologyTab.vue';
import DevicesTab from './DevicesTab.vue';
import ConnectionsTab from './ConnectionsTab.vue';
import MonitoringTab from './MonitoringTab.vue';
import ReportsTab from './ReportsTab.vue';
import DeviceModal from './DeviceModal.vue';
import ConfirmModal from './ConfirmModal.vue';
import { ListDevices, ListConnections } from '../../wailsjs/go/main/App';

const router = useRouter();
const currentTab = ref('Topology');
const isDeviceModalOpen = ref(false);
const editingDevice = ref(null);
const isConfirmModalOpen = ref(false);
const deviceToDelete = ref(null);

// Alert types configuration
const alertTypes = [
  {
    type: 'High CPU',
    check: (device) => device.cpu > 80,
    severity: 'critical',
    title: 'High CPU Usage',
    message: (device) => `CPU usage on ${device.name} is at ${Math.round(device.cpu)}%`
  },
  {
    type: 'High Memory',
    check: (device) => device.memory > 80,
    severity: 'critical',
    title: 'High Memory Usage',
    message: (device) => `Memory usage on ${device.name} is at ${Math.round(device.memory)}%`
  },
  {
    type: 'Medium CPU',
    check: (device) => device.cpu > 60 && device.cpu <= 80,
    severity: 'warning',
    title: 'Elevated CPU Usage',
    message: (device) => `CPU usage on ${device.name} is at ${Math.round(device.cpu)}%`
  },
  {
    type: 'Medium Memory',
    check: (device) => device.memory > 60 && device.memory <= 80,
    severity: 'warning',
    title: 'Elevated Memory Usage',
    message: (device) => `Memory usage on ${device.name} is at ${Math.round(device.memory)}%`
  },
  {
    type: 'Port Error',
    check: (device) => Math.random() < 0.1, // 10% chance
    severity: 'warning',
    title: 'Port Errors Detected',
    message: (device) => `Multiple CRC errors detected on ${device.name}`
  },
  {
    type: 'Connection Lost',
    check: (device) => Math.random() < 0.05, // 5% chance
    severity: 'critical',
    title: 'Connection Issues',
    message: (device) => `Intermittent connection issues detected on ${device.name}`
  }
];

const devices = ref([
  {
    id: 1,
    name: 'Router-01',
    type: 'router',
    ip: '192.168.1.1',
    status: 'online',
    cpu: 45,
    memory: 62,
    uptime: '15d 4h 32m',
    x: 200,
    y: 150,
    vendor: 'Cisco',
    model: 'ISR4331'
  },
  {
    id: 2,
    name: 'Switch-01',
    type: 'switch',
    ip: '192.168.1.2',
    status: 'online',
    cpu: 30,
    memory: 45,
    uptime: '20d 8h 15m',
    x: 400,
    y: 150,
    vendor: 'Cisco',
    model: 'Catalyst 2960'
  },
  {
    id: 3,
    name: 'Server-01',
    type: 'server',
    ip: '192.168.1.10',
    status: 'online',
    cpu: 75,
    memory: 85,
    uptime: '10d 12h 45m',
    x: 300,
    y: 300,
    vendor: 'Dell',
    model: 'PowerEdge R740'
  }
]);

const connections = ref([
  {
    id: 1,
    source: 1,
    target: 2,
    status: 'active',
    bandwidth: '1Gbps',
    latency: 2.3,
    type: 'ethernet'
  },
  {
    id: 2,
    source: 2,
    target: 3,
    status: 'active',
    bandwidth: '1Gbps',
    latency: 1.5,
    type: 'ethernet'
  }
]);

const alerts = ref([
  {
    id: 1,
    device: 'Router-01',
    type: 'High CPU',
    severity: 'warning',
    time: '2 min ago',
    status: 'active',
    title: 'High CPU Usage',
    message: 'CPU usage above 75% for the last 5 minutes'
  },
  {
    id: 2,
    device: 'Switch-01',
    type: 'Port Error',
    severity: 'critical',
    time: '5 min ago',
    status: 'active',
    title: 'Port Errors Detected',
    message: 'Multiple CRC errors on port GigabitEthernet1/0/1'
  },
  {
    id: 3,
    device: 'Server-01',
    type: 'Memory',
    severity: 'warning',
    time: '10 min ago',
    status: 'acknowledged',
    title: 'High Memory Usage',
    message: 'Memory usage reached 85% threshold'
  }
]);

const tabs = ['Topology', 'Devices', 'Connections', 'Monitoring', 'Reports'];

const currentComponent = computed(() => {
  switch (currentTab.value) {
    case 'Topology':
      return TopologyTab;
    case 'Devices':
      return DevicesTab;
    case 'Connections':
      return ConnectionsTab;
    case 'Monitoring':
      return MonitoringTab;
    case 'Reports':
      return ReportsTab;
    default:
      return TopologyTab;
  }
});

async function loadData() {
  try {
    const [devicesData, connectionsData] = await Promise.all([
      ListDevices(),
      ListConnections()
    ]);
    devices.value = devicesData;
    connections.value = connectionsData;
  } catch (error) {
    console.error('Error loading data:', error);
  }
}

function logout() {
  localStorage.removeItem('token');
  router.push('/auth');
}

const handleEditDevice = (device) => {
  editingDevice.value = device;
  isDeviceModalOpen.value = true;
};

const handleSaveDevice = (device) => {
  const index = devices.value.findIndex(d => d.id === device.id);
  if (index !== -1) {
    devices.value[index] = { ...devices.value[index], ...device };
  } else {
    devices.value.push({
      ...device,
      id: Math.max(0, ...devices.value.map(d => d.id)) + 1,
      x: 300,
      y: 200
    });
  }
  isDeviceModalOpen.value = false;
  editingDevice.value = null;
};

const handleDeleteDevice = (deviceId) => {
  deviceToDelete.value = deviceId;
  isConfirmModalOpen.value = true;
};

const confirmDelete = async () => {
  try {
    // TODO: Call backend API
    devices.value = devices.value.filter(d => d.id !== deviceToDelete.value);
    console.log('Device deleted:', deviceToDelete.value);
  } catch (error) {
    console.error('Error deleting device:', error);
  } finally {
    isConfirmModalOpen.value = false;
    deviceToDelete.value = null;
  }
};

const handleDeviceMove = (deviceId, x, y) => {
  const device = devices.value.find(d => d.id === deviceId);
  if (device) {
    device.x = x;
    device.y = y;
    // TODO: Save position to backend
    console.log('Device moved:', deviceId, x, y);
  }
};

const handleDeviceSelect = (device) => {
  console.log('Device selected:', device);
};

// Generate alerts based on device metrics
const generateAlerts = () => {
  const newAlerts = [];
  let alertId = 1;

  devices.value
    .filter(device => device.status === 'online')
    .forEach(device => {
      alertTypes.forEach(alertType => {
        if (alertType.check(device)) {
          // Check if similar alert already exists
          const existingAlert = alerts.value.find(a => 
            a.device === device.name && 
            a.type === alertType.type &&
            a.status === 'active'
          );

          if (!existingAlert) {
            newAlerts.push({
              id: alertId++,
              device: device.name,
              type: alertType.type,
              severity: alertType.severity,
              time: '1 min ago',
              status: 'active',
              title: alertType.title,
              message: alertType.message(device)
            });
          }
        }
      });
    });

  // Update existing alerts time
  alerts.value = alerts.value.map(alert => ({
    ...alert,
    time: alert.time === '1 min ago' ? '2 min ago' : 
          alert.time === '2 min ago' ? '3 min ago' : 
          alert.time === '3 min ago' ? '5 min ago' : alert.time
  }));

  // Remove old alerts (older than 5 minutes or resolved)
  alerts.value = alerts.value.filter(alert => 
    alert.time !== '5 min ago' && 
    alert.status !== 'resolved'
  );

  // Add new alerts
  alerts.value = [...newAlerts, ...alerts.value];

  // Sort alerts by severity and time
  alerts.value.sort((a, b) => {
    if (a.severity === 'critical' && b.severity !== 'critical') return -1;
    if (a.severity !== 'critical' && b.severity === 'critical') return 1;
    return 0;
  });
};

// Interval for updating metrics
let metricsInterval = null;
let alertsInterval = null;

// Update device metrics and generate alerts
const updateMetrics = () => {
  devices.value = devices.value.map(device => {
    if (device.status !== 'online') {
      return {
        ...device,
        cpu: 0,
        memory: 0
      };
    }
    return {
      ...device,
      cpu: Math.min(100, Math.max(0, device.cpu + (Math.random() * 20 - 10))),
      memory: Math.min(100, Math.max(0, device.memory + (Math.random() * 20 - 10)))
    };
  });

  // Generate new alerts after metrics update
  generateAlerts();
};

// Start metrics and alerts update intervals
onMounted(() => {
  loadData();
  metricsInterval = setInterval(updateMetrics, 2000);
});

// Clean up intervals
onUnmounted(() => {
  if (metricsInterval) {
    clearInterval(metricsInterval);
  }
  if (alertsInterval) {
    clearInterval(alertsInterval);
  }
});

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

// Handle device status update
const handleDeviceUpdate = (updatedDevice) => {
  const index = devices.value.findIndex(d => d.id === updatedDevice.id);
  if (index !== -1) {
    // If device is going offline, reset metrics
    if (updatedDevice.status === 'offline') {
      updatedDevice.cpu = 0;
      updatedDevice.memory = 0;
    }
    devices.value[index] = updatedDevice;
  }
};
</script> 