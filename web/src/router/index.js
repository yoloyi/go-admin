import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '../views/Main.vue'
import AuthLogin from '../views/Auth/Login.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Main',
    component: Main
  },
  {
    path: '/auth/login',
    name: 'Login',
    component: AuthLogin
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
