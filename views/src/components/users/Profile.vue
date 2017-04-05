<template>
  <div class="login-box">
    <!-- /.login-logo -->
    <div class="login-box-body">
      <div class="box-body box-profile">
        <img class="profile-user-img img-responsive img-circle" src="~assets/images/avatar04.png" alt="User profile picture">

        <h3 class="profile-username text-center">{{ currentUser.name }}</h3>

        <p class="text-muted text-center">{{ currentUser.title }}</p>

        <ul class="list-group list-group-unbordered">
          <li class="list-group-item">
            <b>Email</b> <a class="pull-right">{{ currentUser.email }}</a>
          </li>
          <li class="list-group-item">
            <b>Contact</b> <a class="pull-right">{{ currentUser.contact }}</a>
          </li>
          <li class="list-group-item">
            <b>Address</b> <a class="pull-right">{{ currentUser.address }}</a>
          </li>
        </ul>

        <router-link to="/profile/edit" class="btn btn-primary btn-block">
          <b>Edit</b>
        </router-link>
      </div>
    </div>
    <br/>
    <div class="text-center">
      <button type="button" @click="logout()" class="btn btn-primary btn-lg">
        <span class="glyphicon glyphicon-off"></span> Sign Out
      </button>
    </div>
  </div>   
</template>

<script>
import authService from './../../services/auth';

export default {
  name: 'Profile',
  data() {
    return {
      errors: [],
      currentUser: this.currentUser(),
    };
  },
  methods: {
    logout() {
      this.errors = [];
      authService.logout(this, '/');
    },
    currentUser() {
      authService.currentUser(this);
    },
  },
  computed: {
    currentUser() {
      return this.$store.state.currentUser;
    },
  },
};
</script>