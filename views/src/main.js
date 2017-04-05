import Vue from 'vue';
import VueRouter from 'vue-router';
import VueResource from 'vue-resource';
import App from './App';
import Router from './routes';
import AuthService from './services/auth';
import store from './store';

Vue.use(VueRouter);
Vue.use(VueResource);
/* global google, document */
const router = new VueRouter({
  linkActiveClass: 'active',
  mode: 'history',
  routes: Router,
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.auth)) {
    // this route requires auth, check if logged in
    // if not, redirect to login page
    if (!AuthService.user.authenticated) {
      next({
        path: '/',
        query: { redirect: to.fullPath },
      });
    } else {
      next();
    }
  } else {
    next(); // make sure to always call next()!
  }
});

// export default router;

/* eslint-disable no-new */
new Vue({
  router,
  el: '#app',
  template: '<App/>',
  store,
  components: { App },
});
