<script setup>
import HeaderComponent from '../components/Header/HeaderComponent.vue'
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { Line } from 'vue-chartjs'
import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend } from 'chart.js'
import { GetAllDevices, GetMetricsByDeviceIdAndInterval, GetLatestMetricsByDeviceId } from "../../wailsjs/go/service/Service.js";

// Регистрируем компоненты Chart.js
ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend)

// Состояние компонента
const devices = ref([])
const selectedDevice = ref(null)
const selectedDeviceId = ref(null)
const selectedInterval = ref('1m')
const selectedTimeRange = ref('1h')
const metrics = ref([])
const updateInterval = ref(null)

const intervals = [
    { value: '1s', label: '1 секунда' },
    { value: '5s', label: '5 секунд' },
    { value: '10s', label: '10 секунд' },
    { value: '30s', label: '30 секунд' },
    { value: '1m', label: '1 минута' },
    { value: '5m', label: '5 минут' }
]

const timeRanges = [
    { value: '5m', label: '5 минут' },
    { value: '15m', label: '15 минут' },
    { value: '30m', label: '30 минут' },
    { value: '1h', label: '1 час' },
    { value: '3h', label: '3 часа' },
    { value: '6h', label: '6 часов' },
    { value: '12h', label: '12 часов' },
    { value: '24h', label: '24 часа' }
]

// Опции для графиков
const chartOptions = {
    responsive: true,
    maintainAspectRatio: false,
    animation: {
        duration: 300
    },
    scales: {
        y: {
            beginAtZero: true,
            max: 100
        }
    },
    plugins: {
        legend: {
            position: 'top'
        }
    }
}

// Данные для графиков
const chartData = ref({
    labels: [],
    datasets: [
        {
            label: 'CPU (%)',
            data: [],
            borderColor: '#FF6384',
            tension: 0.2
        },
        {
            label: 'Memory (%)',
            data: [],
            borderColor: '#36A2EB',
            tension: 0.2
        },
        {
            label: 'Network (%)',
            data: [],
            borderColor: '#4BC0C0',
            tension: 0.2
        }
    ]
})

// Загрузка устройств
async function loadDevices() {
    try {
        devices.value = await GetAllDevices()
        if (devices.value.length > 0) {
            handleDeviceChange(devices.value[0].ID.toString())
        }
    } catch (e) {
        console.error('Error loading devices:', e)
    }
}

// Обновление метрик
async function updateMetrics() {
    if (!selectedDeviceId.value) return

    try {
        // Сначала получим последние метрики для быстрого отображения
        const latestMetrics = await GetLatestMetricsByDeviceId(parseInt(selectedDeviceId.value), 1)
        if (latestMetrics && latestMetrics.length > 0) {
            updateMetricsSummary(latestMetrics[0])
        }

        // Затем получим исторические данные для графика
        const now = new Date()
        const timeRangeInMinutes = parseTimeRange(selectedTimeRange.value)
        const from = new Date(now.getTime() - timeRangeInMinutes * 60 * 1000)

        const newMetrics = await GetMetricsByDeviceIdAndInterval(
            parseInt(selectedDeviceId.value),
            selectedInterval.value,
            from.toISOString(),
            now.toISOString()
        )

        if (newMetrics && newMetrics.length > 0) {
            metrics.value = newMetrics
            updateChartData()
        }
    } catch (e) {
        console.error('Error updating metrics:', e)
    }
}

// Обновление сводки метрик
function updateMetricsSummary(metric) {
    if (!metric) return

    const cpuCard = document.querySelector('.metric-card:nth-child(1) .metric-value')
    const memoryCard = document.querySelector('.metric-card:nth-child(2) .metric-value')
    const networkCard = document.querySelector('.metric-card:nth-child(3) .metric-value')
    const statusCard = document.querySelector('.metric-card:nth-child(4) .metric-value')

    if (cpuCard) {
        cpuCard.textContent = `${metric.cpu}%`
        cpuCard.style.color = getMetricColor(metric.cpu)
    }
    if (memoryCard) {
        memoryCard.textContent = `${metric.memory}%`
        memoryCard.style.color = getMetricColor(metric.memory)
    }
    if (networkCard) {
        networkCard.textContent = `${metric.network}%`
        networkCard.style.color = getMetricColor(metric.network)
    }
    if (statusCard) {
        statusCard.textContent = metric.status
        statusCard.className = `metric-value ${metric.status.toLowerCase()}`
    }
}

// Обновление данных графика
function updateChartData() {
    const sortedMetrics = [...metrics.value].sort((a, b) => 
        new Date(a.CreatedAt).getTime() - new Date(b.CreatedAt).getTime()
    )

    chartData.value.labels = sortedMetrics.map(m => formatDate(m.CreatedAt))
    chartData.value.datasets[0].data = sortedMetrics.map(m => m.cpu)
    chartData.value.datasets[1].data = sortedMetrics.map(m => m.memory)
    chartData.value.datasets[2].data = sortedMetrics.map(m => m.network)
}

// Форматирование даты
function formatDate(date) {
    const d = new Date(date)
    return d.toLocaleTimeString()
}

// Парсинг временного диапазона
function parseTimeRange(range) {
    const value = parseInt(range)
    const unit = range.slice(-1)
    switch (unit) {
        case 'm': return value
        case 'h': return value * 60
        default: return 60 // по умолчанию 1 час
    }
}

// Обработчики изменения выбора
async function handleDeviceChange(deviceId) {
    selectedDeviceId.value = deviceId
    selectedDevice.value = devices.value.find(d => d.ID === parseInt(deviceId))
    await updateMetrics()
}

async function handleIntervalChange(interval) {
    selectedInterval.value = interval
    await updateMetrics()
}

async function handleTimeRangeChange(range) {
    selectedTimeRange.value = range
    await updateMetrics()
}

// Хуки жизненного цикла
onMounted(async () => {
    await loadDevices()
    // Запускаем автоматическое обновление каждые 5 секунд
    updateInterval.value = setInterval(updateMetrics, 5000)
})

onBeforeUnmount(() => {
    // Очищаем интервал при размонтировании компонента
    if (updateInterval.value) {
        clearInterval(updateInterval.value)
    }
})

// Определение цвета метрики в зависимости от значения
function getMetricColor(value) {
    if (value >= 90) return '#FF0000' // Красный для критических значений
    if (value >= 75) return '#FFA500' // Оранжевый для предупреждений
    return '#4BC0C0' // Зеленый для нормальных значений
}
</script>

<template>
    <HeaderComponent />
    <div class="container">
        <div class="controls">
            <div class="control-group">
                <label>Устройство:</label>
                <select :value="selectedDeviceId" @change="handleDeviceChange($event.target.value)">
                    <option value="">Выберите устройство</option>
                    <option v-for="device in devices" :key="device.ID" :value="device.ID">
                        {{ device.hostname }} ({{ device.ip_address }})
                    </option>
                </select>
            </div>

            <div class="control-group">
                <label>Интервал:</label>
                <select :value="selectedInterval" @change="handleIntervalChange($event.target.value)">
                    <option v-for="interval in intervals" :key="interval.value" :value="interval.value">
                        {{ interval.label }}
                    </option>
                </select>
            </div>

            <div class="control-group">
                <label>Период:</label>
                <select :value="selectedTimeRange" @change="handleTimeRangeChange($event.target.value)">
                    <option v-for="range in timeRanges" :key="range.value" :value="range.value">
                        {{ range.label }}
                    </option>
                </select>
            </div>
        </div>

        <div class="charts-container">
            <div class="chart-wrapper">
                <Line :data="chartData" :options="chartOptions" />
            </div>

            <div class="metrics-summary" v-if="metrics.length > 0">
                <div class="metric-card">
                    <h3>CPU</h3>
                    <div class="metric-value" :style="{ color: getMetricColor(metrics[metrics.length-1].cpu) }">
                        {{ metrics[metrics.length-1].cpu }}%
                    </div>
                </div>

                <div class="metric-card">
                    <h3>Memory</h3>
                    <div class="metric-value" :style="{ color: getMetricColor(metrics[metrics.length-1].memory) }">
                        {{ metrics[metrics.length-1].memory }}%
                    </div>
                </div>

                <div class="metric-card">
                    <h3>Network</h3>
                    <div class="metric-value" :style="{ color: getMetricColor(metrics[metrics.length-1].network) }">
                        {{ metrics[metrics.length-1].network }}%
                    </div>
                </div>

                <div class="metric-card">
                    <h3>Status</h3>
                    <div class="metric-value" :class="metrics[metrics.length-1].status.toLowerCase()">
                        {{ metrics[metrics.length-1].status }}
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.container {
    background: #fff;
    border-radius: 15px;
    width: 100%;
    height: 100%;
    padding: 20px;
    display: flex;
    flex-direction: column;
    gap: 20px;
}

.controls {
    display: flex;
    gap: 20px;
    padding: 10px;
    background: #f5f5f5;
    border-radius: 8px;
}

.control-group {
    display: flex;
    flex-direction: column;
    gap: 5px;
}

.control-group label {
    font-size: 14px;
    color: #666;
}

.control-group select {
    padding: 8px;
    border-radius: 4px;
    border: 1px solid #ddd;
    min-width: 200px;
}

.charts-container {
    flex: 1;
    display: flex;
    flex-direction: column;
    gap: 20px;
}

.chart-wrapper {
    flex: 1;
    background: #fff;
    border-radius: 8px;
    padding: 15px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.metrics-summary {
    display: flex;
    gap: 20px;
    padding: 15px;
}

.metric-card {
    flex: 1;
    background: #f8f9fa;
    border-radius: 8px;
    padding: 15px;
    text-align: center;
    box-shadow: 0 2px 4px rgba(0,0,0,0.05);
}

.metric-card h3 {
    margin: 0;
    color: #666;
    font-size: 14px;
    margin-bottom: 10px;
}

.metric-value {
    font-size: 24px;
    font-weight: bold;
}

.metric-value.normal {
    color: #4CAF50;
}

.metric-value.warning {
    color: #FFC107;
}

.metric-value.critical {
    color: #F44336;
}
</style>