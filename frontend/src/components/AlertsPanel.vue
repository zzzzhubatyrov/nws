<template>
  <div class="h-full flex flex-col">
    <div class="p-4 border-b border-gray-700">
      <div class="flex items-center justify-between">
        <h3 class="text-lg font-medium text-white">Active Alerts</h3>
        <span class="px-3 py-1 text-sm rounded-full" :class="alertCountClass">
          {{ activeAlerts.length }} Active
        </span>
      </div>
    </div>

    <div class="flex-1 overflow-auto p-4">
      <div class="space-y-3">
        <div
          v-for="alert in alerts"
          :key="alert.id"
          class="bg-gray-700/50 rounded-lg p-3 flex items-start"
          :class="{
            'border-l-4 border-red-500': alert.severity === 'critical',
            'border-l-4 border-yellow-500': alert.severity === 'warning'
          }"
        >
          <!-- Alert Icon -->
          <div class="mr-3 mt-1">
            <span v-if="alert.severity === 'critical'" class="text-red-500 text-xl">⚠️</span>
            <span v-else-if="alert.severity === 'warning'" class="text-yellow-500 text-xl">⚠️</span>
            <span v-else class="text-blue-500 text-xl">ℹ️</span>
          </div>

          <!-- Alert Content -->
          <div class="flex-grow">
            <div class="flex justify-between items-start">
              <h4 class="text-white font-medium">{{ alert.title }}</h4>
              <span class="text-xs text-gray-400">{{ alert.time }}</span>
            </div>
            <p class="text-gray-300 text-sm mt-1">{{ alert.message }}</p>
            <div class="flex items-center justify-between mt-2">
              <div class="flex items-center text-sm">
                <span class="text-gray-400">Device:</span>
                <span class="text-gray-300 ml-1">{{ alert.device }}</span>
                <span
                  class="ml-2 px-2 py-0.5 rounded text-xs"
                  :class="getStatusClass(alert.status)"
                >
                  {{ alert.status }}
                </span>
              </div>
              
              <!-- Action Buttons -->
              <div class="flex gap-2" v-if="alert.status === 'active'">
                <button
                  @click="$emit('acknowledge', alert.id)"
                  class="text-xs px-2 py-1 bg-yellow-600/20 text-yellow-400 rounded hover:bg-yellow-600/30 transition-colors"
                >
                  Acknowledge
                </button>
                <button
                  @click="$emit('resolve', alert.id)"
                  class="text-xs px-2 py-1 bg-green-600/20 text-green-400 rounded hover:bg-green-600/30 transition-colors"
                >
                  Resolve
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Empty State -->
      <div
        v-if="alerts.length === 0"
        class="text-center py-8 text-gray-400"
      >
        <div class="text-4xl mb-2">🎉</div>
        <p>No active alerts</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';

const props = defineProps({
  alerts: {
    type: Array,
    required: true,
    default: () => []
  }
});

const emit = defineEmits(['acknowledge', 'resolve']);

const activeAlerts = computed(() => 
  props.alerts.filter(alert => alert.status === 'active')
);

const alertCountClass = computed(() => {
  const count = activeAlerts.value.length;
  if (count === 0) return 'bg-green-500/20 text-green-400';
  if (count > 5) return 'bg-red-500/20 text-red-400';
  return 'bg-yellow-500/20 text-yellow-400';
});

const getStatusClass = (status) => {
  switch (status.toLowerCase()) {
    case 'active':
      return 'bg-red-500/20 text-red-400';
    case 'acknowledged':
      return 'bg-yellow-500/20 text-yellow-400';
    case 'resolved':
      return 'bg-green-500/20 text-green-400';
    default:
      return 'bg-gray-500/20 text-gray-400';
  }
};
</script>
