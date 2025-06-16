<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center bg-black/60 backdrop-blur-sm">
    <div class="bg-gray-800 border border-gray-700 rounded-xl w-full max-w-xl p-6 shadow-xl">
      <h2 class="text-white text-lg font-semibold mb-6 flex items-center gap-2">
        <span>{{ form.id ? 'Edit Connection' : 'Add Connection' }}</span>
      </h2>

      <form @submit.prevent="submit" class="grid grid-cols-2 gap-5">
        <!-- Left column -->
        <div class="space-y-4">
          <div>
            <label class="label">Name<span class="text-red-400">*</span></label>
            <input v-model="form.name" class="input" required />
          </div>
          <div>
            <label class="label">Type</label>
            <input v-model="form.type" class="input" placeholder="ethernet" />
          </div>
          <div>
            <label class="label">Status</label>
            <select v-model="form.status" class="input">
              <option value="active">Active</option>
              <option value="inactive">Inactive</option>
            </select>
          </div>
        </div>

        <!-- Right column -->
        <div class="space-y-4">
          <div>
            <label class="label">Source Device<span class="text-red-400">*</span></label>
            <select v-model.number="form.source" class="input" required>
              <option :value="null" disabled>Select source</option>
              <option v-for="d in devices" :key="d.id" :value="d.id">{{ d.name }}</option>
            </select>
          </div>
          <div>
            <label class="label">Target Device<span class="text-red-400">*</span></label>
            <select v-model.number="form.target" class="input" required>
              <option :value="null" disabled>Select target</option>
              <option v-for="d in devices" :key="d.id" :value="d.id">{{ d.name }}</option>
            </select>
          </div>
          <p v-if="error" class="text-red-400 text-xs mt-2">{{ error }}</p>
        </div>

        <!-- Buttons full width below -->
        <div class="col-span-2 flex justify-end gap-3 pt-4">
          <button type="button" @click="$emit('close')" class="btn-secondary">Cancel</button>
          <button type="submit" class="btn-primary" :disabled="!!error">Save</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { reactive, watch, computed } from 'vue'

const props = defineProps({
  isOpen: Boolean,
  connection: Object,
  devices: Array,
})
const emit = defineEmits(['save', 'close'])

const blank = { name: '', source: null, target: null, type: 'ethernet', status: 'active' }
const form = reactive({ ...blank })

watch(
  () => props.connection,
  (c) => Object.assign(form, c ? { ...c } : { ...blank }),
  { immediate: true }
)

const error = computed(() => {
  if (form.source == null || form.target == null) return 'Please select both devices.'
  if (form.source === form.target) return 'Source and target must differ.'
  return ''
})

const submit = () => {
  if (error.value) return
  emit('save', { ...form })
}
</script>

<style scoped>
.input {
  @apply bg-gray-700 border border-gray-600 rounded px-3 py-1 w-full text-white text-sm focus:outline-none focus:ring-1 focus:ring-blue-500 placeholder-gray-400;
}
.label {
  @apply text-gray-300 text-xs mb-1 inline-block;
}
.btn-primary {
  @apply px-4 py-1.5 rounded bg-blue-600 text-white text-sm hover:bg-blue-500 disabled:opacity-40;
}
.btn-secondary {
  @apply px-4 py-1.5 rounded bg-gray-600 text-white text-sm hover:bg-gray-500;
}
select.input option {
  color: #000; /* keep dropdown readable */
}
</style>
