<template>
  <div class="min-h-screen bg-gray-900 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div>
        <h2 class="mt-6 text-center text-3xl font-extrabold text-white">
          {{ isLogin ? 'Sign in to your account' : 'Create new account' }}
        </h2>
      </div>
      <form class="mt-8 space-y-6" @submit.prevent="handleSubmit">
        <div class="rounded-md shadow-sm -space-y-px">
          <div v-if="!isLogin">
            <label for="name" class="sr-only">Name</label>
            <input
              id="name"
              v-model="form.name"
              type="text"
              required
              class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-700 bg-gray-800 text-white rounded-t-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
              placeholder="Name"
            />
          </div>
          <div>
            <label for="email-address" class="sr-only">Email address</label>
            <input
              id="email-address"
              v-model="form.email"
              type="email"
              required
              class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-700 bg-gray-800 text-white focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
              :class="{ 'rounded-t-md': isLogin }"
              placeholder="Email address"
            />
          </div>
          <div>
            <label for="password" class="sr-only">Password</label>
            <input
              id="password"
              v-model="form.password"
              type="password"
              required
              class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-700 bg-gray-800 text-white rounded-b-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
              placeholder="Password"
            />
          </div>
        </div>

        <div>
          <button
            type="submit"
            class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            {{ isLogin ? 'Sign in' : 'Sign up' }}
          </button>
        </div>

        <div class="text-center">
          <button
            type="button"
            class="text-sm text-blue-400 hover:text-blue-500"
            @click="isLogin = !isLogin"
          >
            {{ isLogin ? 'Need an account? Sign up' : 'Already have an account? Sign in' }}
          </button>
        </div>
      </form>

      <div v-if="error" class="mt-4 text-center text-red-500 text-sm">
        {{ error }}
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { Login, Register } from '../../wailsjs/go/main/App';

const router = useRouter();
const isLogin = ref(true);
const error = ref('');
const form = reactive({
  name: '',
  email: '',
  password: ''
});

async function handleSubmit() {
  try {
    error.value = '';
    
    if (isLogin.value) {
      const token = await Login(form.email, form.password);
      localStorage.setItem('token', token);
      router.push('/');
    } else {
      await Register(form.email, form.password, form.name);
      isLogin.value = true;
      form.password = '';
    }
  } catch (err) {
    error.value = err.message || 'An error occurred';
  }
}
</script> 