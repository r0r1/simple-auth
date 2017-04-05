<template>
  <div class="login-box">
    <div class="login-logo">
      <a router-link to="/"><b>Simple Auth</b></a>
    </div>
    <!-- /.login-logo -->
    <div class="login-box-body">
      <p class="login-box-msg">Sign in to start your session</p>
      <div v-if="successMessage" class="alert alert-success">
        <a href="#" class="close" data-dismiss="alert">&times;</a>
        {{ successMessage }}
      </div>

      <form v-on:submit.prevent="login()">
        <div v-if="errors.length > 0">
          <div  class="alert alert-danger">
            <a href="#" class="close" data-dismiss="alert">&times;</a>
            <ul v-for="msg in errors">
                <li>{{ msg }}</li>
            </ul>
          </div>
        </div>

        <div class="form-group has-feedback">
          <input type="email" class="form-control" placeholder="Email" v-model="creds.email" required="required">
          <span class="glyphicon glyphicon-envelope form-control-feedback"></span>
        </div>

        <div class="form-group has-feedback">
          <input type="password" class="form-control" placeholder="Password" v-model="creds.password" required="required">
          <span class="glyphicon glyphicon-lock form-control-feedback"></span>
        </div>

        <div class="row">
          <!-- /.col -->
          <div class="col-xs-12">
            <button type="submit" class="btn btn-primary btn-block btn-flat">Sign In</button>
          </div>
          <!-- /.col -->
        </div>
        
      </form>
      <div class="social-auth-links text-center">
        <p>- OR -</p>
        <a :href="linkedinUrl" class="btn btn-block btn-social btn-linkedin btn-flat">
          <i class="fa fa-linkedin"></i>
          Sign In using Linkedin
        </a>
      </div>
      <div class="pull-right">
        <router-link to="/forgot-password">
          Forgot Password
        </router-link>
      </div>
      <router-link to="/register" class="text-left">
        Register a new membership
      </router-link>
    </div>
  </div>   
</template>

<script>
import { mapGetters } from 'vuex';
import authService from './../../services/auth';

export default {
  name: 'Login',
  data() {
    return {
      errors: [],
      creds: {
        email: null,
        password: null,
      },
      linkedinUrl: this.getLinkedinUrl(),
    };
  },
  methods: {
    login() {
      this.errors = [];
      authService.login(this, this.creds, '/profile');
    },
    getLinkedinUrl() {
      authService.linkedinUrl(this);
    },
  },
  computed: mapGetters([
    'successMessage',
  ]),
};
</script>