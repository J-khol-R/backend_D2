<template>
    <div class="max-w-sm rounded overflow-hidden shadow-lg bg-white mx-auto hover:shadow-2xl transition-shadow duration-300">
        <div class="px-6 py-4">
            <h2 class="font-bold text-xl mb-2 text-gray-800">{{ nombre }}</h2>
            <p class="text-gray-700 text-base">{{ descripcion }}</p>
            <p class="mt-4 inline-block bg-blue-100 text-blue-800 text-xs px-2 py-1 rounded-full uppercase font-semibold tracking-wide">
            Fecha inicio: {{ fechaInicio }}
            </p>
            <p class="mt-4 inline-block bg-blue-100 text-blue-800 text-xs px-2 py-1 rounded-full uppercase font-semibold tracking-wide">
            Fecha limite: {{ fechaLimite }}
            </p>
        </div>
        <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-full m-4" @click="verProyecto">
            Ver Proyecto
        </button>
    </div>
  </template>
  
  <script>
  import { ref } from 'vue';
  import { useRouter } from 'vue-router';
  //import PruebaProyecto from '../pruebas/PruebaProyecto.vue';
  
  export default {
    setup(){

        const router = useRouter();
        const verProyecto = () => {
            router.push({path: `/projects/${codigo}`});
        };
        return {verProyecto};
    },
    props: {
        nombre: String,
        descripcion: String,
        fechaInicio: String,
        fechaLimite: String,
        codigo: String
    }
  }


    
  
  </script>
  
  <style scoped>
  /* Estilos específicos para este componente */
  /* Puedes añadir tus estilos personalizados aquí */
  </style>
  