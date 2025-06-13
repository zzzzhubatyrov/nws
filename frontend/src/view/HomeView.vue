<script setup>
import HeaderComponent from '../components/Header/HeaderComponent.vue'
import TreeNodeComponent from '../components/TreeNodeComponent.vue'
import { ref, onMounted, onBeforeUnmount } from 'vue'
import {
  CreateConnection,
  CreateDevice, DeleteConnection,
  GetAllConnections,
  GetAllDevices,
  GetRecentLogs
} from "../../wailsjs/go/service/Service.js";

const STORAGE_KEY_WIDTH = 'cgraph-width'
const STORAGE_KEY_HEIGHT = 'cgraph-height'

const defaultWidth = 50
const defaultHeight = 30

const cgraphWidth = ref(defaultWidth)
const clogHeight = ref(defaultHeight)

const isResizingHorizontal = ref(false)
const isResizingVertical = ref(false)

const isPanning = ref(false)
const start = ref({ x: 0, y: 0 })
const scrollStart = ref({ left: 0, top: 0 })
const cgraphRef = ref(null)

const devices = ref([])
const logs = ref([])
const logsUpdateInterval = ref(null)

const onMouseDown = (e) => {
  if (e.altKey && cgraphRef.value) {
    isPanning.value = true
    start.value = { x: e.clientX, y: e.clientY }
    scrollStart.value = {
      left: cgraphRef.value.scrollLeft,
      top: cgraphRef.value.scrollTop
    }
    document.body.style.cursor = 'grab'
  }
}

const onMouseMove = (e) => {
  if (isPanning.value && cgraphRef.value) {
    const dx = e.clientX - start.value.x
    const dy = e.clientY - start.value.y
    cgraphRef.value.scrollLeft = scrollStart.value.left - dx
    cgraphRef.value.scrollTop = scrollStart.value.top - dy
  }

  if (isResizingHorizontal.value) {
    const container = document.querySelector('.container')
    const rect = container.getBoundingClientRect()
    const percent = ((e.clientX - rect.left) / rect.width) * 100
    const clamped = Math.min(90, Math.max(10, percent))
    cgraphWidth.value = clamped
    localStorage.setItem(STORAGE_KEY_WIDTH, clamped.toFixed(2))
  }

  if (isResizingVertical.value) {
    const bodyHeight = window.innerHeight
    const percent = ((bodyHeight - e.clientY) / bodyHeight) * 100
    const clamped = Math.min(90, Math.max(10, percent))
    clogHeight.value = clamped
    localStorage.setItem(STORAGE_KEY_HEIGHT, clamped.toFixed(2))
  }
}

const onMouseUp = () => {
  isPanning.value = false
  isResizingHorizontal.value = false
  isResizingVertical.value = false
  document.body.style.cursor = ''
  document.body.style.userSelect = ''
}

const onMouseDownHorizontal = () => {
  isResizingHorizontal.value = true
  document.body.style.userSelect = 'none'
}

const onMouseDownVertical = () => {
  isResizingVertical.value = true
}

onMounted(async () => {
  const savedWidth = localStorage.getItem(STORAGE_KEY_WIDTH)
  const savedHeight = localStorage.getItem(STORAGE_KEY_HEIGHT)

  if (savedWidth) cgraphWidth.value = parseFloat(savedWidth)
  if (savedHeight) clogHeight.value = parseFloat(savedHeight)

  window.addEventListener('mousemove', onMouseMove)
  window.addEventListener('mouseup', onMouseUp)

  const modalConnection = document.getElementById('modalConnection');
  const btnConnection = document.querySelector('.add-connection');
  const closeBtn = modalConnection?.querySelector('.close');

  if (btnConnection && modalConnection && closeBtn) {
    btnConnection.addEventListener('click', () => {
      modalConnection.classList.remove('hidden');
    });

    closeBtn.addEventListener('click', () => {
      modalConnection.classList.add('hidden');
    });

    window.addEventListener('click', (e) => {
      if (e.target === modalConnection) {
        modalConnection.classList.add('hidden');
      }
    });
  }

  const modalAddDevice = document.getElementById('modalAddDevice');
  const btnAddDevice = document.querySelector('.add-device');

  if (btnAddDevice && modalAddDevice && closeBtn) {
    btnAddDevice.addEventListener('click', () => {
      modalAddDevice.classList.remove('hidden');
    });

    closeBtn.addEventListener('click', () => {
      modalAddDevice.classList.add('hidden');
    });

    window.addEventListener('click', (e) => {
      if (e.target === modalAddDevice) {
        modalAddDevice.classList.add('hidden');
      }
    });
  }

  try {
    devices.value = await GetAllDevices()

    // await updateLogs()
    // logsUpdateInterval.value = setInterval(updateLogs, 5000)
  } catch (e) {
    console.error(e)
  }

  await getAllConnections()
})

onBeforeUnmount(() => {
  window.removeEventListener('mousemove', onMouseMove)
  window.removeEventListener('mouseup', onMouseUp)
  
  // Очищаем интервал обновления логов
  // if (logsUpdateInterval.value) {
  //   clearInterval(logsUpdateInterval.value)
  // }
})

const device = {
  hostname: "",
  ipAddress: "",
  type: "",
  vendor: "",
  model: "",
}

// const connection = {
//   source_device_id: ref(0),
//   destination_device_id: ref(0),
// }

const connection = ref([])

async function createConnection(e) {
  e.preventDefault()
  try {
    const data = await CreateConnection(source_device_ip.value, destination_device_ip.value)
    console.log(data)
  } catch (e) {
    console.error(e)
  }
}

async function getAllConnections() {
  try {
    connection.value = await GetAllConnections()
    console.log(connection.value)
  } catch (e) {
    console.error(e)
  }
}

async function deleteConnection(id) {
  try {
    await DeleteConnection(id)
    // console.log(id)
  } catch (e) {
    console.error(e)
  }
}


const createDevice = async (e) => {
  e.preventDefault()
  try {
    const data = await CreateDevice(device)
    console.log(data)
  } catch (e) {
    console.error(e)
  }
}

const source_device_ip = ref("")
const destination_device_ip = ref("")



// Функция обновления логов
async function updateLogs() {
  try {
    // Получаем последние 50 логов
    logs.value = await GetRecentLogs(50)
  } catch (e) {
    console.error("Error fetching logs:", e)
  }
}

// Функция форматирования даты
function formatDate(date) {
  return new Date(date).toLocaleString()
}

// Функция определения цвета для действия
function getActionColor(action) {
  const colors = {
    'CONNECT': '#4CAF50',
    'DISCONNECT': '#F44336',
    'CONFIG_CHANGE': '#2196F3',
    'STATUS_CHECK': '#9C27B0',
    'INTERFACE_UP': '#4CAF50',
    'INTERFACE_DOWN': '#F44336',
    'SECURITY_ALERT': '#FF5722',
    'PERFORMANCE_WARNING': '#FFC107',
    'SYSTEM_REBOOT': '#795548',
    'FIRMWARE_UPDATE': '#607D8B'
  }
  return colors[action] || '#000000'
}

const isConnectionTab = ref(false)

</script>


<template>
    <HeaderComponent />
    <div class="container">
      <div
          class="cgraph"
          ref="cgraphRef"
      >
        <TreeNodeComponent />
      </div>
      <div class="chierarchy">
        <div class="chierarchy-header">
          <button class="btn add-device">add device +</button>
          <button class="btn add-connection">add connection +</button>
          <button class="btn connections">Соединения</button>
          <button class="btn devices">Устройства</button>
        </div>
        <div class="chierarchy-content" v-for="connect in connection" :key="connect.ID">
          <div class="connections-card">
            <div class="connection-from">source_device: {{ connect.source_device_ip }}</div>
            <div class="connection-to">destination_device: {{ connect.destination_device_ip || "null" }}</div>
            <div class="connection-status">status: {{ connect.status }}</div>
            <div class="connection-protocol">protocol: {{ connect.protocol }}</div>
            <div class="connection-port">port: {{ connect.port }}</div>
            <div class="connection-createdAt">created_at: {{ connect.CreatedAt }}</div>
            <button class="btn delete-connection" @click="deleteConnection(connect.ID)">delete connection</button>
          </div>
        </div>
      </div>
    </div>
    <div id="modalConnection" class="modalConnection hidden">
      <div class="modalConnection-content">
        <span class="close">&times;</span>
        <h2>Добавить соединение</h2>
        <form class="create-connection_form" @submit="createConnection">
          <select v-model="source_device_ip">
            <option disabled value="">Выберите источник</option>
            <option v-for="dev in devices" :key="dev.ip_address" :value="dev.ip_address">
              {{ dev.hostname }} — {{ dev.ip_address }}
            </option>
          </select>

          <select v-model="destination_device_ip">
            <option disabled value="">Выберите получателя</option>
            <option v-for="dev in devices" :key="dev.ip_address + '-dst'" :value="dev.ip_address">
              {{ dev.hostname }} — {{ dev.ip_address }}
            </option>
          </select>

          <button class="btn create-connection">create connection</button>
        </form>
      </div>
    </div>
    <div id="modalAddDevice" class="modalAddDevice hidden">
      <div class="modalAddDevice-content">
        <span class="close">&times;</span>
        <h2>Добавить устройство</h2>
        <form class="create-device_form" @submit="createDevice">
          <input v-model="device.hostname" type="text" name="hostname" placeholder="hostname">
          <input v-model="device.ipAddress" type="text" name="ipAddress" placeholder="ipAddress">
          <input v-model="device.type" type="text" name="type" placeholder="type">
          <input v-model="device.vendor" type="text" name="vendor" placeholder="vendor">
          <input v-model="device.model" type="text" name="model" placeholder="model">
          <button class="btn create-device">create device</button>
        </form>
      </div>
    </div>
    <div class="clog">
      <div class="clog-header">
        <button class="btn export-logs">export to csv</button>
      </div>
      <div class="clog-content">
        <div class="logs-container">
          <div v-for="log in logs" :key="log.ID" class="log-entry">
            <div class="log-time">{{ formatDate(log.timestamp) }}</div>
            <div class="log-action" :style="{ color: getActionColor(log.action) }">
              {{ log.action }}
            </div>
            <div class="log-message">{{ log.message }}</div>
          </div>
        </div>
      </div>
    </div>
</template>

<style scoped>
.container {
  display: flex;
  gap: 10px;
  width: 100%;
  height: 100%;
  transition: height .5ms;
}
.cgraph, .chierarchy, .clog {
  background-color: #fff;
  width: 100%;
  height: 100%;
  border-radius: 15px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.chierarchy {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: start;
  max-width: 500px;

  padding: 20px 20px;
  gap: 10px;
}
.chierarchy-header {
  display: flex;
  width: 100%;
  height: 40px;
  gap: 10px;
}
.chierarchy-content {
  width: 100%;
  height: 100%;
}
.clog {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: start;
  height: 300px;
  padding: 20px 20px;
  gap: 10px;
}
.clog-header {
  display: flex;
  width: 100%;
  align-items: center;
  justify-content: end;
}
.btn {
  border: solid #000 1px;
  padding: 5px 10px;
  border-radius: 10px;
  font-size: 16px;
}

.create-device_form, .create-connection_form {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
.resizer-horizontal {
  width: 5px;
  cursor: col-resize;
  background-color: transparent;
  transition: background-color 0.2s;
}
.container:hover .resizer-horizontal {
  background-color: #ccc;
  border-radius: 4px;
}
.resizer-vertical {
  height: 5px;
  cursor: row-resize;
  background-color: transparent;
  transition: background-color 0.2s;
  width: 100%;
}
.resizer-vertical:hover {
  background-color: #ccc;
  border-radius: 4px;
}

.clog-content {
  width: 100%;
  height: calc(100% - 40px);
  overflow-y: auto;
  padding: 10px;
}

.logs-container {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.log-entry {
  display: grid;
  grid-template-columns: 180px 150px 1fr;
  gap: 10px;
  padding: 8px;
  background-color: #f5f5f5;
  border-radius: 6px;
  font-size: 14px;
  align-items: center;
}

.log-time {
  color: #666;
  font-size: 12px;
}

.log-action {
  font-weight: 600;
  text-transform: uppercase;
  font-size: 12px;
}

.log-message {
  color: #333;
}

/* Стилизация скроллбара */
.clog-content::-webkit-scrollbar {
  width: 8px;
}

.clog-content::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

.clog-content::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 4px;
}

.clog-content::-webkit-scrollbar-thumb:hover {
  background: #555;
}
</style>