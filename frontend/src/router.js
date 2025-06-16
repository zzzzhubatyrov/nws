import { createRouter, createWebHistory } from 'vue-router';
import { ValidateToken } from '../wailsjs/go/main/App';
import AuthPage from './components/AuthPage.vue';
import MainLayout from './components/MainLayout.vue';

const routes = [
  {
    path: '/auth',
    name: 'Auth',
    component: AuthPage,
    meta: { public: true }
  },
  {
    path: '/',
    name: 'Home',
    component: MainLayout
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach(async (to, from, next) => {
  const token = localStorage.getItem('token');
  const isPublicRoute = to.matched.some(record => record.meta.public);

  if (!token && !isPublicRoute) {
    next('/auth');
    return;
  }

  if (token && !isPublicRoute) {
    try {
      await ValidateToken(token);
      next();
    } catch (err) {
      localStorage.removeItem('token');
      next('/auth');
    }
    return;
  }

  if (token && to.path === '/auth') {
    next('/');
    return;
  }

  next();
});

export default router; 