import { SIGNUP_URL, LOGIN_URL, CURRENT_USER_URL, FORGOT_PASSWORD_URL, LINKEDIN_URL } from './main';
// URL and endpoint constants


/* global localStorage: false, console: false, $: false */
/* eslint no-param-reassign: [2, { "props": false }]*/

export default {

  // User object will let us check authentication status
  user: {
    authenticated: false,
  },

  // Send a request to the login URL and save the returned JWT
  login(context, creds, redirect) {
    context.$http.post(LOGIN_URL, creds)
      .then((res) => {
        if (res.data.data) {
          localStorage.setItem('token', res.data.data);
          this.user.authenticated = true;
          context.$store.dispatch('login');
          context.$router.replace(redirect);
        }
      }, (err) => {
        context.errors.push(err.data.error);
      });
  },

  linkedinUrl(context) {
    context.$http.get(LINKEDIN_URL)
      .then((res) => {
        console.log(res.data.data);
        context.linkedinUrl = res.data.data;
      }, (err) => {
        console.log('error linkedin', err.data);
        context.errors.push(err.data.error);
      });
  },

  signup(context, creds, redirect) {
    context.$http.post(SIGNUP_URL, creds)
      .then((res) => {
        if (res.data.ID && res.data.ID !== 0) {
          context.$router.replace(redirect);
          context.$store.dispatch('addSuccessMessage', 'Register has been successful, please login.');
        }
      }, (err) => {
        context.errors.push(err.data.error);
      });
  },

  forgotPassword(context, creds, redirect) {
    context.$http.post(FORGOT_PASSWORD_URL, creds)
      .then((res) => {
        if (res.data.data) {
          context.$router.replace(redirect);
          context.$store.dispatch('addSuccessMessage', 'Forgot password has been successful, please login.');
        }
      }, (err) => {
        context.errors.push(err.data.error);
      });
  },

  currentUser(context) {
    context.$http.get(CURRENT_USER_URL, this.getAuthHeader())
      .then((res) => {
        localStorage.getItem('user_id', res.data.ID);
        context.profile = res.data;
        context.$store.dispatch('currentUser', res.data);
      }, (err) => {
        const errors = err.data.error;
        context.$store.dispatch('addErrorMessages', errors);
        context.errors.push(errors);
      });
  },

  // To log out, we just need to remove the token
  logout(context, redirect) {
    localStorage.removeItem('token');
    this.user.authenticated = false;
    context.$store.dispatch('logout');
    context.$router.replace(redirect);
  },

  checkAuth() {
    const jwt = localStorage.getItem('token');
    if (jwt) {
      this.user.authenticated = true;
    } else {
      this.user.authenticated = false;
    }
  },

  // The object to be passed as a header for authenticated requests
  getAuthHeader(jwtToken = null) {
    let token = null;
    if (jwtToken != null) {
      token = jwtToken;
    } else {
      token = localStorage.getItem('token');
    }
    return {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    };
  },
};
