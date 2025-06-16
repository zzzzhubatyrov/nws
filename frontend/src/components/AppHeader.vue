<template>
  <header class="bg-gray-800 border-b border-gray-600">
    <div class="px-6 py-4">
      <div class="flex justify-between items-center">
        <div class="flex items-center gap-3">
          <div class="bg-blue-600 p-2 rounded-lg">
            <NetworkIcon class="w-6 h-6" />
          </div>
          <div>
            <h1 class="text-xl font-bold">Network Simulator</h1>
            <p class="text-gray-400 text-sm">Enterprise Network Management Platform</p>
          </div>
        </div>
        <div class="flex items-center gap-4">
          <div class="text-right">
            <div class="text-sm text-gray-400">Status: All Systems</div>
            <div class="text-green-500 font-medium">Operational</div>
          </div>
          <div class="w-2 h-2 bg-green-500 rounded-full animate-pulse"></div>
        </div>
      </div>
    </div>

    <!-- Navigation Tabs -->
    <nav class="px-6">
      <div class="flex space-x-8">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          @click="$emit('update:activeTab', tab.id)"
          :class="{
            'border-blue-500 text-blue-500': activeTab === tab.id,
            'border-transparent text-gray-400 hover:text-gray-300': activeTab !== tab.id
          }"
          class="py-3 px-1 border-b-2 font-medium text-sm flex items-center gap-2 transition-colors"
        >
          <component :is="tab.icon" class="w-4 h-4" />
          {{ tab.name }}
        </button>
      </div>
    </nav>
  </header>
</template>

<script setup>
import { Network as NetworkIcon } from 'lucide-vue-next';

const props = defineProps({
  tabs: {
    type: Array,
    required: true
  },
  activeTab: {
    type: String,
    required: true
  }
});
</script>
