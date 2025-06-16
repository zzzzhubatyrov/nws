<template>
  <div class="space-y-6">
    <!-- Inner tabs -->
    <div class="flex gap-4 border-b border-gray-700 mb-4">
      <button
        v-for="t in innerTabs"
        :key="t"
        @click="innerTab = t"
        :class="[
          'pb-2 px-2 text-sm',
          innerTab === t ? 'border-b-2 border-blue-500 text-white' : 'text-gray-400'
        ]"
      >
        {{ t }}
      </button>
    </div>

    <!-- Tickets list -->
    <div v-if="innerTab === 'Tickets'" class="space-y-3">
      <p class="text-gray-400">No tickets yet – coming soon…</p>
    </div>

    <!-- Charts -->
    <div v-else>
      <!-- Filters -->
      <div class="bg-gray-800 border border-gray-700 rounded-lg p-6 mb-6">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
          <!-- Device Selection -->
          <div>
            <label class="block text-sm font-medium text-gray-400 mb-2">Select Devices</label>
            <div class="space-y-2">
              <label v-for="device in devices" :key="device.id" class="flex items-center">
                <input
                  type="checkbox"
                  v-model="selectedDevices"
                  :value="device.id"
                  class="form-checkbox h-4 w-4 text-blue-500 bg-gray-700 border-gray-600 rounded"
                >
                <span class="ml-2 text-sm text-gray-300">{{ device.name }}</span>
              </label>
            </div>
          </div>

          <!-- Time Range Selection -->
          <div>
            <label class="block text-sm font-medium text-gray-400 mb-2">Time Range</label>
            <select
              v-model="selectedTimeRange"
              class="w-full bg-gray-700 border border-gray-600 rounded-md py-2 px-3 text-white focus:outline-none focus:ring-2 focus:ring-blue-500"
            >
              <option v-for="range in timeRanges" :key="range.value" :value="range.value">
                {{ range.label }}
              </option>
            </select>
          </div>
        </div>

        <!-- Apply Filters Button -->
        <div class="mt-4 flex justify-end">
          <button
            @click="applyFilters"
            class="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-md text-sm"
          >
            Apply Filters
          </button>
        </div>
      </div>

      <!-- Chart -->
      <div class="bg-gray-800 border border-gray-700 rounded-lg p-6">
        <Line v-if="hasData" :data="chartData" :options="chartOptions" />
        <div v-else class="text-center text-gray-400 py-8">
          Please select at least one device to display the chart
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch, onMounted } from 'vue';
import { defineComponent } from 'vue';
import { Line } from 'vue-chartjs';
import {
  Chart as ChartJS,
  LineElement,
  PointElement,
  LinearScale,
  Title,
  CategoryScale,
  Tooltip,
  Legend,
  TimeScale
} from 'chart.js';
import 'chartjs-adapter-date-fns';
import { ListDeviceMetrics } from '../../wailsjs/go/main/App';

ChartJS.register(
  LineElement,
  PointElement,
  LinearScale,
  Title,
  CategoryScale,
  Tooltip,
  Legend,
  TimeScale
);

const props = defineProps({
  devices: { type: Array, default: () => [] }
});

const innerTabs = ['Charts', 'Tickets'];
const innerTab = ref('Charts');

// Filter states
const selectedDevices = ref([]);
const selectedTimeRange = ref('1h');
const metrics = ref([]);

const timeRanges = [
  { label: 'Last Hour', value: '1h' },
  { label: 'Last 6 Hours', value: '6h' },
  { label: 'Last 24 Hours', value: '24h' },
  { label: 'Last 7 Days', value: '7d' },
  { label: 'Last 30 Days', value: '30d' }
];

// Initialize with all devices selected
watch(() => props.devices, (newDevices) => {
  if (newDevices.length && selectedDevices.value.length === 0) {
    selectedDevices.value = newDevices.map(d => d.id);
  }
}, { immediate: true });

const hasData = computed(() => selectedDevices.value.length > 0 && metrics.value.length > 0);

// Fetch metrics for selected devices
async function fetchMetrics() {
  try {
    const promises = selectedDevices.value.map(deviceId => ListDeviceMetrics(deviceId));
    const results = await Promise.all(promises);
    metrics.value = results.flat().map(metric => ({
      ...metric,
      timestamp: new Date(metric.timestamp).getTime()
    }));
  } catch (error) {
    console.error('Failed to fetch metrics:', error);
  }
}

// Apply filters and fetch new data
async function applyFilters() {
  await fetchMetrics();
}

// Process metrics based on time range
const filteredMetrics = computed(() => {
  const now = Date.now();
  const ranges = {
    '1h': 60 * 60 * 1000,
    '6h': 6 * 60 * 60 * 1000,
    '24h': 24 * 60 * 60 * 1000,
    '7d': 7 * 24 * 60 * 60 * 1000,
    '30d': 30 * 24 * 60 * 60 * 1000
  };
  
  const timeLimit = now - ranges[selectedTimeRange.value];
  return metrics.value.filter(m => m.timestamp > timeLimit);
});

// Build chart data based on filtered metrics
const chartData = computed(() => {
  if (!hasData.value) return { datasets: [] };

  const deviceNames = props.devices
    .filter(d => selectedDevices.value.includes(d.id))
    .reduce((acc, d) => ({ ...acc, [d.id]: d.name }), {});

  const datasets = [];
  
  // Group metrics by device and type
  const groupedMetrics = {};
  filteredMetrics.value.forEach(m => {
    const key = `${m.device_id}-${m.name}`;
    if (!groupedMetrics[key]) {
      groupedMetrics[key] = [];
    }
    groupedMetrics[key].push(m);
  });

  // Create datasets for each device and metric type
  Object.entries(groupedMetrics).forEach(([key, metrics]) => {
    const [deviceId, metricName] = key.split('-');
    const color = metricName === 'cpu' ? '#3b82f6' : '#10b981';
    const alpha = metricName === 'cpu' ? 'rgba(59,130,246,0.3)' : 'rgba(16,185,129,0.3)';
    
    // Sort metrics by timestamp
    const sortedMetrics = [...metrics].sort((a, b) => a.timestamp - b.timestamp);
    
    datasets.push({
      label: `${deviceNames[deviceId]} - ${metricName.toUpperCase()}`,
      data: sortedMetrics.map(m => ({
        x: m.timestamp,
        y: m.value
      })),
      borderColor: color,
      backgroundColor: alpha,
      tension: 0.4,
      pointRadius: 2
    });
  });

  return {
    datasets
  };
});

const chartOptions = {
  responsive: true,
  maintainAspectRatio: false,
  interaction: {
    mode: 'index',
    intersect: false
  },
  scales: {
    y: {
      beginAtZero: true,
      max: 100,
      grid: {
        color: 'rgba(255,255,255,0.1)'
      },
      ticks: {
        color: 'rgba(255,255,255,0.7)'
      }
    },
    x: {
      type: 'time',
      time: {
        unit: 'minute',
        displayFormats: {
          minute: 'HH:mm',
          hour: 'HH:mm',
          day: 'MMM d'
        }
      },
      grid: {
        color: 'rgba(255,255,255,0.1)'
      },
      ticks: {
        color: 'rgba(255,255,255,0.7)',
        maxRotation: 0
      }
    }
  },
  plugins: {
    legend: {
      position: 'top',
      labels: {
        color: 'rgba(255,255,255,0.7)',
        padding: 20,
        usePointStyle: true,
        pointStyle: 'circle'
      }
    },
    tooltip: {
      mode: 'index',
      intersect: false,
      backgroundColor: 'rgba(0,0,0,0.8)',
      titleColor: 'rgba(255,255,255,0.9)',
      bodyColor: 'rgba(255,255,255,0.9)',
      borderColor: 'rgba(255,255,255,0.2)',
      borderWidth: 1
    }
  }
};

// Initial data fetch
onMounted(() => {
  if (selectedDevices.value.length > 0) {
    fetchMetrics();
  }
});
</script>

<style scoped>
.chart-container {
  height: 400px;
}
</style>
