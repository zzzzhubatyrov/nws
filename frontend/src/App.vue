<template>
  <div class="min-h-screen bg-gray-900 text-white">
    <AppHeader
      :tabs="tabs"
      :activeTab="activeTab"
      @update:activeTab="activeTab = $event"
    />

    <!-- Основной контент -->
    <main class="p-6">
      <TopologyTab
        v-if="activeTab === 'topology'"
        :devices="devices"
        :connections="connections"
        :alerts="alerts"
        :selected-device="selectedDevice"
        @device-select="setSelectedDevice"
        @device-move="handleDeviceMove"
        @edit-device="handleEditDevice"
        @delete-device="handleDeleteDevice"
      />

      <DevicesTab
        v-if="activeTab === 'devices'"
        :devices="devices"
        :selected-device="selectedDevice"
        @device-select="setSelectedDevice"
        @add-device="handleAddDevice"
        @auto-device="handleAutoDevice"
        @edit-device="handleEditDevice"
        @delete-device="handleDeleteDevice"
      />

      <ConnectionsTab
        v-if="activeTab === 'connections'"
        :connections="connections"
        :devices="devices"
        @add-connection="handleAddConnection"
        @edit-connection="handleEditConnection"
        @delete-connection="handleDeleteConnection"
      />

      <MonitoringTab v-if="activeTab === 'monitoring'" :devices="devices" />

      <ReportsTab v-if="activeTab === 'reports'" :devices="devices" />

      <!-- Alerts -->
      <div
        v-if="activeTab === 'alerts'"
        class="bg-gray-800 border border-gray-600 rounded-lg p-8 text-center"
      >
        <div class="text-gray-600 mb-4">
          <BellIcon  />
        </div>
        <h3 class="text-white text-lg font-medium mb-2">Alert Management</h3>
        <AlertsPanel :alerts="alerts" />
      </div>

    </main>

    <!-- Модальные окна -->
    <DeviceModal
        :is-open="isDeviceModalOpen"
        :device="editingDevice"
        @close="isDeviceModalOpen = false"
        @save="handleSaveDevice"
    />
    <ConnectionModal
        :is-open="isConnectionModalOpen"
        :connection="editingConnection"
        :devices="devices"
        @close="isConnectionModalOpen = false"
        @save="handleSaveConnection"
    />
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted, watch } from 'vue';
import {
  Monitor as MonitorIcon,
  Server as ServerIcon,
  Router as RouterIcon,
  Wifi as WifiIcon,
  Activity as ActivityIcon,
  AlertTriangle as AlertTriangleIcon,
  CheckCircle as CheckCircleIcon,
  XCircle as XCircleIcon,
  Clock as ClockIcon,
  Cpu as CpuIcon,
  HardDrive as HardDriveIcon,
  Network as NetworkIcon,
  Zap as ZapIcon,
  Settings as SettingsIcon,
  Plus as PlusIcon,
  Edit as EditIcon,
  Trash2 as Trash2Icon,
  Search as SearchIcon,
  Filter as FilterIcon,
  LineChart as LineChartIcon,
  PieChart as PieChartIcon,
  Bell as BellIcon,
  Shield as ShieldIcon,
  Database as DatabaseIcon,
  Cable as CableIcon,
  X as XIcon,
  BarChart3 as BarChart3Icon
} from 'lucide-vue-next';

// Импорт компонентов
import AppHeader from './components/AppHeader.vue';
// Wails backend functions
import { ListDevices, SaveDevice, DeleteDevice, ListConnections, SaveConnection, DeleteConnection, GenerateDevice, ListAlerts } from '../wailsjs/go/main/App';
import TopologyTab from './components/TopologyTab.vue';
import DevicesTab from './components/DevicesTab.vue';
import ConnectionsTab from './components/ConnectionsTab.vue';
import DeviceModal from './components/DeviceModal.vue';
import ConnectionModal from './components/ConnectionModal.vue';
import MonitoringTab from './components/MonitoringTab.vue';
import AlertsPanel from './components/AlertsPanel.vue';
import ReportsTab from './components/ReportsTab.vue';

// Моковые данные
const mockDevices = ref([
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
  // ... другие устройства
]);

const mockConnections = ref([
  {
    id: 1,
    source: 1,
    target: 2,
    status: 'active',
    bandwidth: '1Gbps',
    latency: 2.3,
    type: 'ethernet',
    name: 'R1-SW1',
    description: 'Primary link'
  },
  // ... другие соединения
]);

const mockAlerts = ref([
  {
    id: 1,
    device: 'Firewall-01',
    type: 'High CPU',
    severity: 'warning',
    time: '2 min ago',
    status: 'active',
    title: 'High CPU Usage',
    message: 'CPU usage above 75%'
  },
  // ... другие алерты
]);

// Состояние приложения
const devices = ref([]);
const connections = ref([]);
const alerts = ref([]);
const selectedDevice = ref(null);
const activeTab = ref('topology');
const isDeviceModalOpen = ref(false);
const isConnectionModalOpen = ref(false);
const editingDevice = ref(null);
const editingConnection = ref(null);
const lastSaveRef = ref(0);
const pendingMove = ref(null);

const tabs = [
  { id: 'topology', name: 'Network Topology', icon: NetworkIcon },
  { id: 'devices', name: 'Device Management', icon: ServerIcon },
  { id: 'connections', name: 'Connections', icon: CableIcon },
  { id: 'monitoring', name: 'Monitoring', icon: ActivityIcon },
  { id: 'alerts', name: 'Alerts', icon: BellIcon },
  { id: 'reports', name: 'Reports', icon: BarChart3Icon }
];

onMounted(async () => {
  devices.value = await ListDevices();
  connections.value = await ListConnections();
  alerts.value = await ListAlerts();
  const poll = setInterval(async () => {
    devices.value = await ListDevices();
    alerts.value = await ListAlerts();
  }, 15000);
  // clear if app reloaded/hot replaced
  if (import.meta.hot) {
    import.meta.hot.dispose(() => clearInterval(poll));
  }
});

// Методы

const setSelectedDevice = (device) => {
  selectedDevice.value = device;
};

const handleAddDevice = () => {
  editingDevice.value = null;
  isDeviceModalOpen.value = true;
};

const handleEditDevice = (device) => {
  editingDevice.value = device;
  isDeviceModalOpen.value = true;
};

const handleDeleteDevice = async (deviceId) => {
  if (!confirm('Are you sure you want to delete this device?')) return;
  await DeleteDevice(deviceId);
  devices.value = devices.value.filter(d => d.id !== deviceId);
  if (selectedDevice.value?.id === deviceId) selectedDevice.value = null;
};

const handleAutoDevice = async () => {
  const generated = await GenerateDevice();
  devices.value.push(generated);
};

const handleSaveDevice = async (device) => {
  const saved = await SaveDevice(device);
  const idx = devices.value.findIndex(d => d.id === saved.id);
  if (idx !== -1) devices.value[idx] = saved; else devices.value.push(saved);
  isDeviceModalOpen.value = false;
};

const handleAddConnection = () => {
  editingConnection.value = null;
  isConnectionModalOpen.value = true;
};

const handleEditConnection = (connection) => {
  editingConnection.value = connection;
  isConnectionModalOpen.value = true;
};

const handleDeleteConnection = async (id) => {
  if (!confirm('Delete this connection?')) return;
  await DeleteConnection(id);
  connections.value = connections.value.filter(c => c.id !== id);
};

const handleSaveConnection = async (connection) => {
  const saved = await SaveConnection(connection);
  const idx = connections.value.findIndex(c => c.id === saved.id);
  if (idx !== -1) connections.value[idx] = saved; else connections.value.push(saved);

  isConnectionModalOpen.value = false;
};

const flushSave = async () => {
  if (!pendingMove.value) return;
  const { id, x, y } = pendingMove.value;

  const device = devices.value.find(d => d.id === id);
  if (device) {
    device.x = x;
    device.y = y;
  }

  lastSaveRef.value = Date.now();
  pendingMove.value = null;
};

const handleDeviceMove = (id, x, y) => {
  const device = devices.value.find(d => d.id === id);
  if (device) {
    device.x = x;
    device.y = y;
  }

  if (Date.now() - lastSaveRef.value > 400) {
    flushSave();
  }

  pendingMove.value = { id, x, y };
  setTimeout(flushSave, 400);
};

// Обновление выбранного устройства при изменении списка устройств
watch(devices, (newDevices) => {
  if (selectedDevice.value) {
    const freshDevice = newDevices.find(d => d.id === selectedDevice.value.id);
    if (freshDevice) {
      selectedDevice.value = { ...freshDevice };
    }
  }
}, { deep: true });
</script>