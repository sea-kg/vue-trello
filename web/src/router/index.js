import Vue from 'vue';
import VueRouter from 'vue-router';

import Store from '@/store/index';

import PageBoards from '../views/PageBoards.vue';
import PageBoard from '../views/PageBoard.vue';
import PageHome from '../views/PageHome.vue';
import PageLogin from '../views/PageLogin.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'home',
    component: PageHome,
  },
  {
    path: '/login',
    name: 'login',
    component: PageLogin,
  },
  {
    path: '/boards',
    name: 'boards',
    component: PageBoards,
    meta: { requiresAuth: true },
  },
  {
    path: '/boards/:id',
    name: 'board',
    component: PageBoard,
    meta: { requiresAuth: true },
  },
  {
    path: '*',
    redirect: '/',
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth) && !Store.state.user.user.login) {
    next({ path: '/', query: { redirect: to.fullPath } });
  } else {
    next();
  }
});

export default router;