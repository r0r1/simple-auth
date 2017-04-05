import Login from './components/auth/Login';
import OAuth from './components/auth/OAuth';
import Register from './components/auth/Register';
import ForgotPassword from './components/auth/ForgotPassword';
import Profile from './components/users/Profile';
import EditProfile from './components/users/EditProfile';

export default [
  {
    path: '/',
    component: Login,
  },
  {
    path: '/register',
    component: Register,
  },
  {
    path: '/oauth',
    component: OAuth,
  },
  {
    path: '/forgot-password',
    component: ForgotPassword,
  },
  {
    path: '/profile',
    component: Profile,
    meta: { auth: true },
  },
  {
    path: '/profile/edit',
    component: EditProfile,
    meta: { auth: true },
  },
];
