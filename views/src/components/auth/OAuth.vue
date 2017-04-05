<template>
    <div >
        OAuth
    </div>
</template>

<script>
import authService from './../../services/auth';

/* global localStorage: false, console: false, $: false */
/* eslint no-param-reassign: [2, { "props": false }]*/
/* global $route, $router */
export default {
  name: 'OAuth',
  data() {
    return {
    };
  },
  ready() {
    console.log('OAUTH');
  },
  beforeRouteEnter(to, from, next) {
    localStorage.setItem('token', to.query.token);
    authService.user.authenticated = true;
    if (to.query.type === 'login') {
      next('/profile');
    } else {
      next('/profile/edit');
    }
    next();
  },
};
</script>