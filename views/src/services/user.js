import { USER_URL } from './main';
import authService from './auth';

/* global localStorage: false, console: false, $: false */
/* eslint no-param-reassign: [2, { "props": false }]*/

export default {
  update(context, id, data, redirect) {
    context.$http.patch(`${USER_URL}/${id}`, data, authService.getAuthHeader())
      .then((res) => {
        if (res.data.ID) {
          context.$store.dispatch('addSuccessMessage', 'Update profile has been successful.');
          context.$router.replace(redirect);
        }
      }, (err) => {
        console.log(err);
        context.errors.push(err.data.error);
      });
  },
};
