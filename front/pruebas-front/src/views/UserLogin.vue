<template>
    <div class="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div class="max-w-md w-full space-y-8">
        <div>
          <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
            Iniciar Sesión
          </h2>
        </div>
        <form class="mt-8 space-y-6" @submit.prevent="submitForm">
          <div class="rounded-md shadow-sm -space-y-px">
            <div>
              <label for="username" class="sr-only">Nombre de Usuario</label>
              <input id="username" name="username" type="text" v-model="username" required class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Nombre de usuario">
            </div>
            <div>
              <label for="password" class="sr-only">Contraseña</label>
              <input id="password" name="password" type="password" v-model="password" required class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Contraseña">
            </div>
          </div>
          <div>
            <button type="submit" class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              Ingresar
            </button>
            <button @click="goToRegister" class="mt-1 group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
              Registrarse
            </button>
          </div>
        </form>
      </div>
    </div>
  </template>
  
  <script>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';
  import sessionService from '../services/sessionService'
  import { toast } from 'vue3-toastify';
  import 'vue3-toastify/dist/index.css';
  
  export default {
    setup() {
      const username = ref('');
      const password = ref('');
      const router = useRouter();
  
      const submitForm = async () => {
        const user = {
          username: username.value,
          password: password.value
        };
        
        try {
        const response = await sessionService.getSession(user);

        if (response.data) {
          const { id_user, token } = response.data;
          localStorage.setItem('user_id', id_user);
          localStorage.setItem('token', token);

          toast.success('Login successful!', {
            autoClose: 3000 
          });
          setTimeout(() => {
            router.push('/projects'); 
          }, 3000); 
          console.log(response)
        } else {
          toast.error('Login failed.', {
            autoClose: 5000 
          });
        }

        } catch (error) {
          toast.error('An error occurred during login.', {
            autoClose: 5000 
          });
        };
      };

      const goToRegister = () => {
      router.push('/register');
      };  
  
      return {
        username,
        password,
        submitForm,
        goToRegister
      };
    }
  };
  </script>