<template>
  <div class="flex min-h-screen">
    <AsideDash />
    <div class="flex-grow bg-gray-100 flex flex-col justify-center items-center">
      <h1 class="font-sans font-medium text-4xl pb-14 pt-6 text-center">Crear Proyecto</h1>
      <div class="w-full max-w-xs">
        <form @submit.prevent="submitForm" class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="nombreProyecto">Nombre</label>
            <input
              id="nombreProyecto"
              type="text"
              v-model="nombreProyecto"
              class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              placeholder="Nombre del proyecto"
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="descripcionProyecto">Descripción</label>
            <textarea
              id="descripcionProyecto"
              v-model="descripcionProyecto"
              class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              placeholder="Descripción del proyecto"
            ></textarea>
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="fechaInicio">Fecha inicio</label>
            <input
              id="fechaInicio"
              type="date"
              v-model="fechaInicio"
              class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            />
          </div>
          <div class="mb-4">
            <label class="block text-gray-700 text-sm font-bold mb-2" for="fechaLimite">Fecha límite</label>
            <input
              id="fechaLimite"
              type="date"
              v-model="fechaLimite"
              class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            />
          </div>
          <div class="flex items-center justify-center">
            <button
              type="submit"
              class="btn btn-primary group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              Agregar Proyecto
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>
  
  <script>
  import { ref } from 'vue';
  import projectService from '../services/projectService';
  import { toast } from 'vue3-toastify';
  import 'vue3-toastify/dist/index.css';
  import AsideDash from '../components/AsideDash.vue'
  
  export default {
    setup() {
      const nombreProyecto = ref('');
      const fechaInicio = ref('');
      const fechaLimite = ref('');
      const descripcionProyecto = ref('');
  
      const submitForm = async () => {
        const project = {
            name: nombreProyecto.value,
            startDate: fechaInicio.value,
            finishDate: fechaLimite.value,
            description: descripcionProyecto.value
        };

        const response = await projectService.createProject(project);
        if (response.status === 201) {
            toast.success('project created successfully!', {
            autoClose: 3000 // 3 segundos
          });
          // setTimeout(() => {
          //   router.push('/login');
          // }, 3000); // Espera 3 segundos antes de redirigir
        } else {
            toast.error('Failed to create project.', {
            autoClose: 5000 // 5 segundos
          });
        }
        console.log(response)

        };

        return {
        nombreProyecto,
        fechaInicio,
        fechaLimite,
        descripcionProyecto,
        submitForm
      };
    },
    components: {
      AsideDash
    }
  };
  </script>
