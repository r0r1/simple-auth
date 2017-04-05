<template>
  <div class="register-box">
    <div class="register-logo">
      <a router-link to="/">
        <b>Simple Auth</b>
      </a>
    </div>

    <div class="register-box-body">
      <p class="login-box-msg">Register a new membership</p>

      <form v-on:submit.prevent="register()">
        <div v-if="errors.length > 0">
          <div  class="alert alert-danger">
            <a href="#" class="close" data-dismiss="alert">&times;</a>
            <ul v-for="msg in errors">
                <li>{{ msg }}</li>
            </ul>
          </div>
        </div>
        <div class="form-group has-feedback">
          <input type="email" class="form-control" placeholder="Email" v-model="signup.email" required="required">
          <span class="glyphicon glyphicon-envelope form-control-feedback"></span>
        </div>
        <div class="form-group has-feedback">
          <input type="password" class="form-control" placeholder="Password" v-model="signup.password" required="required">
          <span class="glyphicon glyphicon-lock form-control-feedback"></span>
        </div>
        <div class="row">
          <!-- /.col -->
          <div class="col-xs-12">
            <button type="submit" class="btn btn-primary btn-block btn-flat">Register</button>
          </div>
          <!-- /.col -->
        </div>
      </form>

      <div class="social-auth-links text-center">
        <p>- OR -</p>
        <a :href="linkedinUrl" class="btn btn-block btn-social btn-linkedin btn-flat">
          <i class="fa fa-linkedin"></i>
          Sign Up using Linkedin
        </a>
      </div>

      <router-link to="/" class="text-center">
        I already have a membership
      </router-link>
    </div>
    <!-- /.form-box -->
  </div>
</template>

<script>
import authService from './../../services/auth';

export default {
  name: 'Register',
  data() {
    return {
      errors: [],
      signup: {
        name: null,
        email: null,
        password: null,
      },
      linkedinUrl: this.getLinkedinUrl(),
    };
  },
  methods: {
    register() {
      this.errors = [];
      authService.signup(this, this.signup, '/');
    },
    getLinkedinUrl() {
      authService.linkedinUrl(this);
    },
  },
};
</script>