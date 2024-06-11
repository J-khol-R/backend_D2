import { createRouter, createWebHistory } from 'vue-router';
import UserLogin from '../views/UserLogin.vue';
import UserRegister from '../views/UserRegister.vue';
import HomeView from '../views/HomeView.vue';
import ProjectView from '../views/ProjectView.vue';
import CreateProject from '../views/CreateProject.vue';
import CreateTest from '../views/CreateTest.vue';

const routes = [
  {
    path: '/',
    redirect: '/login'  
  },
  {
    path: '/login',
    name: 'Login',
    component: UserLogin
  },
  {
    path: '/register',
    name: 'Register',
    component: UserRegister,
  },
  // {
  //   path: '/projects',
  //   name: 'Projects',
  //   component: HomeView,
  // },
  {
    path: '/create-project',
    name: 'CreateProject',
    component: CreateProject,
  },
  {
    path: '/create-test',
    name: 'CreateTest',
    component: CreateTest,
  },
  {
    path: '/project/:id',
    name: 'ProjectView',
    component: ProjectView,
  },
  // {
  //   path: '/users',
  //   name: 'Users',
  //   component: HomeView,
  // },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;