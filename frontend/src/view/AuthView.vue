<script setup>
import {ref} from "vue";
import {useRouter} from "vue-router";
import {Login, Register} from "../../wailsjs/go/service/Service.js";


const router = useRouter();

const loginName = ref("")
const loginPassword = ref("")
const regName = ref("")
const regPassword = ref("")

const isLogin = ref(true)
const errorMsg = ref("")

async function login(e) {
  e.preventDefault()
  errorMsg.value = ""
  const user = {
    username: loginName.value,
    password: loginPassword.value
  }
  try {
    const res = await Login(user)
    if (res && res.token) {
      localStorage.setItem("authToken", res.token)
      await router.push("/index")
    } else {
      errorMsg.value = "Неверный email или пароль"
    }
  } catch (err) {
    errorMsg.value = "Ошибка авторизации"
  } finally {
  }
}

async function registration(e) {
  e.preventDefault()
  errorMsg.value = ""
  const user = {
    username: regName.value,
    password: regPassword.value
  }
  try {
    const res = await Register(user)
    if (res && res.id) {
      isLogin.value = true
      errorMsg.value = "Регистрация успешна! Войдите в аккаунт."
    } else {
      errorMsg.value = "Ошибка регистрации"
    }
  } catch (err) {
    errorMsg.value = "Ошибка регистрации"
  }
}
</script>

<template>
  <div class="auth-bg">
    <div class="auth-card">
      <div class="switcher">
        <button :class="['switch-btn', {active: isLogin}]" @click="isLogin = true">Вход</button>
        <button :class="['switch-btn', {active: !isLogin}]" @click="isLogin = false">Регистрация</button>
      </div>
      <transition name="fade-slide" mode="out-in">
        <form v-if="isLogin" key="login" class="form" @submit="login">
          <h2>Вход в NEXORA</h2>
          <input type="text" v-model="loginName" placeholder="Имя" required autocomplete="username">
          <input type="password" v-model="loginPassword" placeholder="Пароль" required autocomplete="current-password">
          <button class="btn">Войти</button>
        </form>
        <form v-else key="register" class="form" @submit="registration">
          <h2>Регистрация</h2>
          <input type="text" v-model="regName" placeholder="Имя" required autocomplete="username">
          <input type="password" v-model="regPassword" placeholder="Пароль" required autocomplete="new-password">
          <button class="btn">Зарегистрироваться</button>
        </form>
      </transition>
    </div>
    <p v-if="errorMsg" class="error-msg">{{ errorMsg }}</p>
  </div>
</template>

<style scoped>
.auth-bg {
  min-height: 100vh;
  min-width: 100vw;
  background: linear-gradient(120deg, #4285f4 0%, #4ecdc4 100%);
  display: flex;
  align-items: center;
  justify-content: center;
}
.auth-card {
  background: #fff;
  border-radius: 18px;
  box-shadow: 0 8px 32px rgba(60,60,120,0.18);
  padding: 40px 32px 32px 32px;
  min-width: 340px;
  max-width: 90vw;
  display: flex;
  flex-direction: column;
  align-items: center;
  position: relative;
}
.switcher {
  display: flex;
  gap: 0;
  width: 100%;
  margin-bottom: 24px;
}
.switch-btn {
  flex: 1;
  padding: 12px 0;
  border: none;
  background: #f0f0f0;
  color: #333;
  font-size: 1.1rem;
  font-weight: 500;
  border-radius: 8px 8px 0 0;
  cursor: pointer;
  transition: background 0.2s, color 0.2s;
}
.switch-btn.active {
  background: #4285f4;
  color: #fff;
}
.form {
  display: flex;
  flex-direction: column;
  gap: 18px;
  width: 260px;
  animation: fadeIn 0.5s;
}
.form h2 {
  margin-bottom: 8px;
  color: #4285f4;
  font-weight: 700;
  letter-spacing: 1px;
}
.form input {
  padding: 10px 12px;
  border: 1px solid #d0d0d0;
  border-radius: 6px;
  font-size: 1rem;
  outline: none;
  transition: border 0.2s;
}
.form input:focus {
  border: 1.5px solid #4ecdc4;
}
.btn {
  background: linear-gradient(90deg, #4285f4 0%, #4ecdc4 100%);
  color: #fff;
  border: none;
  border-radius: 6px;
  padding: 10px 0;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s;
}
.btn:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
.error-msg {
  color: #e53935;
  margin-top: 18px;
  text-align: center;
  font-size: 1rem;
}
.fade-slide-enter-active, .fade-slide-leave-active {
  transition: all 0.4s cubic-bezier(.4,0,.2,1);
}
.fade-slide-enter-from {
  opacity: 0;
  transform: translateY(30px);
}
.fade-slide-leave-to {
  opacity: 0;
  transform: translateY(-30px);
}
@keyframes fadeIn {
  from { opacity: 0; transform: scale(0.98); }
  to { opacity: 1; transform: scale(1); }
}
</style>