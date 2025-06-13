<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { GetAllTopology, GetDeviceById, GetAllConnections } from "../../wailsjs/go/service/Service.js";

const devices = ref([])
const nodes = ref([])
const connections = ref([])
const updateInterval = ref(null)
const tooltip = ref({
    show: false,
    x: 0,
    y: 0,
    device: null
})

const containerWidth = ref(500)
const containerHeight = ref(400)
const scale = ref(1)

const draggingNode = ref(null)
const offset = ref({ x: 0, y: 0 })
const graphContainerRef = ref(null)

onMounted(async () => {
    if (graphContainerRef.value) {
        updateContainerSize()
    }

    await updateData()
    
    window.addEventListener('resize', onResize)
    // updateInterval.value = setInterval(updateData, 5000)
})

onUnmounted(() => {
    window.removeEventListener('resize', onResize)
    if (updateInterval.value) {
        clearInterval(updateInterval.value)
    }
})

function updateContainerSize() {
    const rect = graphContainerRef.value.getBoundingClientRect()
    containerWidth.value = rect.width
    containerHeight.value = rect.height
}

function onResize() {
    updateContainerSize()
    if (nodes.value.length > 0) {
        fitNodesToContainer()
    }
}

function fitNodesToContainer() {
    if (nodes.value.length === 0) return

    let minX = Infinity
    let minY = Infinity
    let maxX = -Infinity
    let maxY = -Infinity

    nodes.value.forEach(node => {
        minX = Math.min(minX, node.x)
        minY = Math.min(minY, node.y)
        maxX = Math.max(maxX, node.x)
        maxY = Math.max(maxY, node.y)
    })

    const padding = 100
    const contentWidth = maxX - minX + padding * 2
    const contentHeight = maxY - minY + padding * 2

    const scaleX = containerWidth.value / contentWidth
    const scaleY = containerHeight.value / contentHeight
    scale.value = Math.min(scaleX, scaleY, 1)

    const offsetX = (containerWidth.value - contentWidth * scale.value) / 2 - minX * scale.value + padding
    const offsetY = (containerHeight.value - contentHeight * scale.value) / 2 - minY * scale.value + padding

    nodes.value = nodes.value.map(node => ({
        ...node,
        x: node.x * scale.value + offsetX,
        y: node.y * scale.value + offsetY
    }))
}

function getNodeCoordinates(nodeIp) {
    const node = nodes.value.find(n => n.ip === nodeIp)
    if (!node) {
        console.warn(`Node with IP ${nodeIp} not found`)
        return { x: 0, y: 0 }
    }
    return {
        x: node.x + 30,
        y: node.y + 30
    }
}

function getConnectionLabelPosition(conn) {
    const fromPos = getNodeCoordinates(conn.from)
    const toPos = getNodeCoordinates(conn.to)
    
    const x = (fromPos.x + toPos.x) / 2
    const y = (fromPos.y + toPos.y) / 2 - 10
    
    return { x, y }
}

async function loadTopology() {
  try {
    devices.value = await GetAllTopology()

    devices.value.map(item => {
      console.log('Topology item:', item)
    })

    nodes.value = devices.value.map(item => ({
      id: item.ID,
      ip: item.device?.ip_address,
      x: item.x * containerWidth.value,
      y: item.y * containerHeight.value,
      model: item.device?.model || 'Без модели',
      customStyle: item.custom_style || ''
    }))

    fitNodesToContainer()

    console.log("Nodes:", nodes.value)
    console.log("Devices:", devices.value)
  } catch (e) {
    console.error('Error loading topology:', e)
  }
}

async function loadConnections() {
  try {
    const allConnections = await GetAllConnections()
    console.log('All connections:', allConnections)

    connections.value = allConnections.map(conn => ({
      id: conn.ID,
      from: conn.source_device_ip,
      to: conn.destination_device_ip,
      protocol: conn.protocol,
      status: conn.status,
      latency: conn.latency,
      packetLoss: conn.packet_loss
    }))

    console.log("Processed connections:", connections.value)
  } catch (e) {
    console.error('Error loading connections:', e)
  }
}

async function loadDeviceById() {
  try {
    const data = await GetDeviceById(1)
    console.log('DEVICE_BY_ID', data)
  } catch (e) {
    console.error(e)
  }
}

function parseCustomStyle(styleString) {
  if (!styleString) return {}
  return styleString
      .split(';')
      .filter(s => s.trim() !== '')
      .reduce((styleObj, item) => {
        const [key, value] = item.split(':')
        if (key && value) {
          styleObj[key.trim()] = value.trim()
        }
        return styleObj
      }, {})
}

function startDrag(node, event) {
  draggingNode.value = node
  offset.value = {
    x: event.clientX - node.x,
    y: event.clientY - node.y,
  }
  window.addEventListener('mousemove', onDrag)
  window.addEventListener('mouseup', stopDrag)
}

function onDrag(event) {
  if (!draggingNode.value) return
  draggingNode.value.x = event.clientX - offset.value.x
  draggingNode.value.y = event.clientY - offset.value.y
}

function stopDrag() {
  draggingNode.value = null
  window.removeEventListener('mousemove', onDrag)
  window.removeEventListener('mouseup', stopDrag)
}

function getConnectionColor(connection) {
  if (connection.status === 'Active') return '#4CAF50' // зеленый для активных
  if (connection.status === 'Disable') return '#F44336' // красный для отключенных
  return '#2196F3' // синий по умолчанию
}

function getConnectionStrokeWidth(connection) {
  // Толщина линии в зависимости от потери пакетов
  if (connection.packetLoss > 50) return 3
  if (connection.packetLoss > 20) return 2
  return 1
}

function getConnectionStrokeDasharray(connection) {
  if (connection.status === 'Disable') return '5,3'
  return '0'
}

function onWheel(event) {
  const delta = event.deltaY > 0 ? -0.1 : 0.1
  scale.value = Math.min(Math.max(scale.value + delta, 0.2), 3)
}

// Функции для работы с тултипом
function showTooltip(node, event) {
    const device = devices.value.find(d => d.device?.ip_address === node.ip)
    if (!device) return

    tooltip.value = {
        show: true,
        x: event.clientX,
        y: event.clientY,
        device: device.device
    }
}

function hideTooltip() {
    tooltip.value.show = false
}

function formatUptime(seconds) {
    const hours = Math.floor(seconds / 3600)
    const minutes = Math.floor((seconds % 3600) / 60)
    return `${hours}ч ${minutes}м`
}

// Объединяем функции обновления данных
async function updateData() {
    try {
        devices.value = await GetAllTopology()
        
        // Обновляем узлы
        nodes.value = devices.value.map(item => ({
            id: item.ID,
            ip: item.device?.ip_address,
            x: item.x * containerWidth.value,
            y: item.y * containerHeight.value,
            model: item.device?.model || 'Unknown',
            customStyle: item.custom_style || ''
        }))

        // Обновляем соединения
        const allConnections = await GetAllConnections()
        connections.value = allConnections.map(conn => ({
            id: conn.ID,
            from: conn.source_device_ip,
            to: conn.destination_device_ip,
            protocol: conn.protocol,
            status: conn.status,
            latency: conn.latency,
            packetLoss: conn.packet_loss
        }))

        // Подгоняем узлы под размер контейнера
        if (nodes.value.length > 0) {
            fitNodesToContainer()
        }

        console.log("Updated nodes:", nodes.value)
        console.log("Updated connections:", connections.value)
    } catch (e) {
        console.error('Error updating data:', e)
    }
}
</script>

<template>
  <div ref="graphContainerRef" class="graph-container" @wheel.prevent="onWheel">
    <!-- Добавляем тултип -->
    <div v-if="tooltip.show" 
         class="tooltip" 
         :style="{ 
             left: `${tooltip.x + 10}px`, 
             top: `${tooltip.y + 10}px` 
         }">
        <div class="tooltip-header">
            <strong>{{ tooltip.device?.hostname }}</strong>
        </div>
        <div class="tooltip-content">
            <p><strong>IP:</strong> {{ tooltip.device?.ip_address }}</p>
            <p><strong>Тип:</strong> {{ tooltip.device?.type }}</p>
            <p><strong>Модель:</strong> {{ tooltip.device?.model }}</p>
            <p><strong>Вендор:</strong> {{ tooltip.device?.vendor }}</p>
            <p><strong>OS:</strong> {{ tooltip.device?.os_version }}</p>
            <p><strong>Статус:</strong> 
                <span :class="['status-badge', tooltip.device?.status]">
                    {{ tooltip.device?.status }}
                </span>
            </p>
        </div>
    </div>

    <div
        class="graph-zoom-wrapper"
        :style="{
          width: `${containerWidth}px`,
          height: `${containerHeight}px`,
          transform: `scale(${scale})`,
          transformOrigin: '0 0',
          position: 'relative',
        }"
    >
      <!-- SVG соединения -->
      <svg
          class="connections-layer"
          :width="containerWidth"
          :height="containerHeight"
          :style="{
              position: 'absolute',
              top: 0,
              left: 0,
              pointerEvents: 'none',
              zIndex: 1,
              transform: `scale(${1/scale})`,
              transformOrigin: '0 0'
          }"
      >
        <defs>
          <marker
              id="arrowhead"
              markerWidth="10"
              markerHeight="7"
              refX="9"
              refY="3.5"
              orient="auto"
          >
            <polygon points="0 0, 10 3.5, 0 7" :fill="'#4f46e5'" />
          </marker>
        </defs>

        <!-- Линии -->
        <g v-for="conn in connections" :key="conn.id">
            <line
                :x1="getNodeCoordinates(conn.from).x * scale"
                :y1="getNodeCoordinates(conn.from).y * scale"
                :x2="getNodeCoordinates(conn.to).x * scale"
                :y2="getNodeCoordinates(conn.to).y * scale"
                :stroke="getConnectionColor(conn)"
                :stroke-width="getConnectionStrokeWidth(conn) * scale"
                :stroke-dasharray="getConnectionStrokeDasharray(conn)"
                marker-end="url(#arrowhead)"
            />
            
            <text
                :x="getConnectionLabelPosition(conn).x * scale"
                :y="getConnectionLabelPosition(conn).y * scale"
                class="connection-label"
                :style="{ fontSize: `${12 * scale}px` }"
            >
                {{ conn.protocol }} ({{ conn.latency }}ms)
            </text>
        </g>
      </svg>

      <!-- Узлы -->
      <div
          v-for="node in nodes"
          :key="node.id"
          class="graph-node"
          :style="{
          left: `${node.x}px`,
          top: `${node.y}px`,
          ...parseCustomStyle(node.customStyle)
        }"
          :title="`${node.model}`"
          @mousedown="startDrag(node, $event)"
          @mouseover="showTooltip(node, $event)"
          @mouseleave="hideTooltip"
      >
        {{ node.model }}
      </div>
    </div>
  </div>
</template>


<style scoped>
.graph-container {
  position: relative;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

.graph-zoom-wrapper {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  transform-origin: 0 0;
}

.connections-layer {
  position: absolute;
  top: 0;
  left: 0;
  pointer-events: none;
  z-index: 1;
  overflow: visible;
}

.connection-label {
  font-size: 12px;
  fill: #666;
  text-anchor: middle;
  dominant-baseline: middle;
  pointer-events: none;
  font-family: Arial, sans-serif;
}

.graph-node {
  position: absolute;
  width: 60px;
  height: 60px;
  background-color: #4f46e5;
  border-radius: 50%;
  color: white;
  font-size: 12px;
  font-weight: bold;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  user-select: none;
  cursor: grab;
  z-index: 10;
  transition: transform 0.2s, box-shadow 0.2s;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.graph-node:hover {
  transform: scale(1.05);
  box-shadow: 0 0 10px rgba(79, 70, 229, 0.7);
  z-index: 100;
}

.tooltip {
    position: fixed;
    z-index: 1000;
    background: white;
    border-radius: 8px;
    box-shadow: 0 2px 12px rgba(0, 0, 0, 0.15);
    padding: 12px;
    min-width: 200px;
    pointer-events: none;
}

.tooltip-header {
    border-bottom: 1px solid #eee;
    padding-bottom: 8px;
    margin-bottom: 8px;
    font-size: 14px;
}

.tooltip-content {
    font-size: 12px;
}

.tooltip-content p {
    margin: 4px 0;
}

.status-badge {
    padding: 2px 6px;
    border-radius: 4px;
    font-size: 11px;
    font-weight: bold;
}

.status-badge.up {
    background-color: #4CAF50;
    color: white;
}

.status-badge.down {
    background-color: #F44336;
    color: white;
}

.status-badge.warning {
    background-color: #FFC107;
    color: black;
}
</style>
