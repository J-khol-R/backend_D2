import { createRouter, createWebHistory } from 'vue-router';
import UserLogin from '../views/UserLogin.vue';
import UserRegister from '../views/UserRegister.vue';
import HomeView from '../views/HomeView.vue';
import CreateProject from '../views/CreateProject.vue';
import HomeProjects from '../views/HomeProjects.vue';
import UsersView from '../views/UsersView.vue';

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
  {
    path: '/projects',
    name: 'Projects',
    component: HomeProjects,
  },
  { 
    path: '/create-project',
    name: 'CreateProject',
    component: CreateProject,
  },
  {
    path: '/users',
    name: 'Users',
    component: UsersView,
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/login'
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;