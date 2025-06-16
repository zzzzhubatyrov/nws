<template>
  <div v-if="isOpen" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
    <div class="bg-gray-800 rounded-lg shadow-xl w-full max-w-lg">
      <!-- Header -->
      <div class="px-6 py-4 border-b border-gray-700">
        <h3 class="text-lg font-medium text-white">
          {{ device ? 'Edit Device' : 'Add New Device' }}
        </h3>
      </div>

      <!-- Form -->
      <form @submit.prevent="handleSubmit" class="p-6">
        <div class="space-y-4">
          <!-- Name -->
          <div>
            <label class="block text-sm font-medium text-gray-400 mb-1">
              Device Name
            </label>
            <input
              v-model="form.name"
              type="text"
              class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-blue-500"
              required
            />
          </div>

          <!-- Type -->
          <div>
            <label class="block text-sm font-medium text-gray-400 mb-1">
              Type
            </label>
            <select
              v-model="form.type"
              class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-blue-500"
              required
            >
              <option value="router">Router</option>
              <option value="switch">Switch</option>
              <option value="server">Server</option>
              <option value="firewall">Firewall</option>
              <option value="workstation">Workstation</option>
            </select>
          </div>

          <!-- IP Address -->
          <div>
            <label class="block text-sm font-medium text-gray-400 mb-1">
              IP Address
            </label>
            <input
              v-model="form.ip"
              type="text"
              pattern="^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$"
              class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-blue-500"
              required
            />
          </div>

          <!-- Vendor -->
          <div>
            <label class="block text-sm font-medium text-gray-400 mb-1">
              Vendor
            </label>
            <input
              v-model="form.vendor"
              type="text"
              class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-blue-500"
              required
            />
          </div>

          <!-- Model -->
          <div>
            <label class="block text-sm font-medium text-gray-400 mb-1">
              Model
            </label>
            <input
              v-model="form.model"
              type="text"
              class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-blue-500"
              required
            />
          </div>

          <!-- Status -->
          <div>
            <label class="block text-sm font-medium text-gray-400 mb-1">
              Status
            </label>
            <select
              v-model="form.status"
              class="w-full bg-gray-700 border border-gray-600 rounded-lg px-3 py-2 text-white focus:outline-none focus:border-blue-500"
              required
            >
              <option value="online">Online</option>
              <option value="offline">Offline</option>
            </select>
          </div>
        </div>

        <!-- Actions -->
        <div class="mt-6 flex justify-end space-x-3">
          <button
            type="button"
            @click="$emit('close')"
            class="px-4 py-2 text-sm font-medium text-gray-400 hover:text-white transition-colors duration-200"
          >
            Cancel
          </button>
          <button
            type="submit"
            class="px-4 py-2 text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 rounded-lg transition-colors duration-200"
          >
            Save
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue';

const props = defineProps({
  isOpen: {
    type: Boolean,
    required: true
  },
  device: {
    type: Object,
    default: null
  }
});

const emit = defineEmits(['close', 'save']);

const form = ref({
  name: '',
  type: 'router',
  ip: '',
  vendor: '',
  model: '',
  status: 'online'
});

// Update form when device changes
watch(() => props.device, (newDevice) => {
  if (newDevice) {
    form.value = { ...newDevice };
  } else {
    form.value = {
      name: '',
      type: 'router',
      ip: '',
      vendor: '',
      model: '',
      status: 'online'
    };
  }
}, { immediate: true });

const handleSubmit = () => {
  emit('save', {
    ...form.value,
    id: props.device?.id
  });
};
</script>

<style scoped>
.input {
  @apply bg-gray-700 border border-gray-600 rounded px-2 py-1 w-full text-white text-sm;
}
.label {
  @apply text-gray-400 text-xs mb-1 inline-block;
}
.btn-primary {
  @apply px-3 py-1 rounded bg-blue-600 text-white text-sm hover:bg-blue-500;
}
.btn-secondary {
  @apply px-3 py-1 rounded bg-gray-600 text-white text-sm hover:bg-gray-500;
}
</style>
