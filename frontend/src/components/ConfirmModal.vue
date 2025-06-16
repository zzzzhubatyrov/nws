<template>
  <div v-if="isOpen" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-gray-800 rounded-lg shadow-xl w-full max-w-md p-6 transform transition-all">
      <!-- Icon -->
      <div class="mx-auto flex items-center justify-center h-12 w-12 rounded-full bg-red-100 mb-4">
        <span class="text-2xl text-red-600">⚠️</span>
      </div>

      <!-- Content -->
      <div class="text-center">
        <h3 class="text-lg font-medium text-white mb-2">
          {{ title }}
        </h3>
        <p class="text-sm text-gray-400">
          {{ message }}
        </p>
      </div>

      <!-- Actions -->
      <div class="mt-6 flex justify-end space-x-3">
        <button
          @click="$emit('cancel')"
          class="px-4 py-2 text-sm font-medium text-gray-400 hover:text-white transition-colors duration-200"
        >
          Cancel
        </button>
        <button
          @click="$emit('confirm')"
          class="px-4 py-2 text-sm font-medium text-white bg-red-600 hover:bg-red-700 rounded-lg transition-colors duration-200"
        >
          Delete
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  isOpen: {
    type: Boolean,
    required: true
  },
  title: {
    type: String,
    default: 'Confirm Action'
  },
  message: {
    type: String,
    default: 'Are you sure you want to perform this action?'
  }
});

defineEmits(['confirm', 'cancel']);
</script> 