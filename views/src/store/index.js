import Vue from 'vue';
import Vuex from 'vuex';
import * as actions from './actions';
import * as getters from './getters';
import * as mutations from './mutations';
import authService from './../services/auth';

Vue.use(Vuex);

authService.checkAuth();

const state = {
  checkAuth: authService.user.authenticated,
  currentUser: {},
  success_message: null,
  error_messages: [],
};

const store = new Vuex.Store({
  state,
  getters,
  actions,
  mutations,
});

export default store;
