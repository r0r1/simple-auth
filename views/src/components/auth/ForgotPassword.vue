<template>
  <div class="login-box">
    <div class="login-logo">
      <a router-link to="/"><b>Simple Auth</b></a>
    </div>
    <!-- /.login-logo -->
    <div class="login-box-body">
      <p class="login-box-msg">Forgot Password</p>
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
          <input type="password" class="form-control" placeholder="New Password" v-model="creds.password" required="required">
          <span class="glyphicon glyphicon-envelope form-control-feedback"></span>
        </div>

        <div class="row">
          <!-- /.col -->
          <div class="col-xs-12">
            <button type="submit" class="btn btn-primary btn-block btn-flat">Submit</button>
          </div>
          <!-- /.col -->
        </div>
        
      </form>
      <div class="social-auth-links text-center">
        <p>- OR -</p>
      </div>
      <router-link to="/">
        I already have a membership
      </router-link>
    </div>
  </div>   
</template>

<script>
import { mapGetters } from 'vuex';
import authService from './../../services/auth';

export default {
  name: 'ForgotPassword',
  data() {
    return {
      errors: [],
      creds: {
        email: null,
        password: null,
      },
    };
  },
  methods: {
    login() {
      this.errors = [];
      authService.forgotPassword(this, this.creds, '/');
    },
  },
  computed: mapGetters([
    'successMessage',
  ]),
};
</script>