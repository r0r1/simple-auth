<template>
  <transition>
    <router-view></router-view>
  </transition>
</template>

<script>
import { mapGetters } from 'vuex';

export default {
  name: 'App',
  data() {
    return {
      title: 'Simple Auth',
    };
  },
  computed: mapGetters([
    'login',
    'logout',
  ]),
};
</script>

<style src="../static/adminLte/css/skins/_all-skins.min.css"></style>
<style src="../static/adminLte/css/AdminLTE.min.css"></style>
