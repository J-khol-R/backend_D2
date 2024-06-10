<template>
    <div>
      <h1 class="font-sans font-medium text-4xl pb-14 pt-6 text-center">Crear Proyecto</h1>
      <div class="flex justify-center">
        <form class="w-96" @submit.prevent="handleAgregarProyecto">
          <div>
            <label class="block font-sans font-medium text-xl">Nombre</label>
            <input
              type="text"
              class="input input-bordered w-full my-2"
              v-model="nombreProyecto"
            />
          </div>
          <div>
            <label class="block font-sans font-medium text-xl">Descripción</label>
            <input
              type="text"
              class="input input-bordered w-full my-3"
              v-model="descripcionProyecto"
            />
          </div>
          <div>
            <label class="block font-sans font-medium text-xl">Fecha inicio</label>
            <input
              type="date"
              class="input input-bordered w-full my-2"
              v-model="fechaInicio"
            />
          </div>
          <div>
            <label class="block font-sans font-medium text-xl">Fecha limite</label>
            <input
              type="date"
              class="input input-bordered w-full my-2"
              v-model="fechaLimite"
            />
          </div>
          <div>
            <button
              type="submit"
              class="btn btn-outline w-full"
            >
              Agregar Proyecto
            </button>
          </div>
        </form>
      </div>
    </div>
  </template>
  
  <script>
    import ProjectService from '../services/projectService';

  export default {
    data() {
      return {
        nombreProyecto: '',
        fechaInicio: '',
        fechaLimite: '',
        descripcionProyecto: ''
      };
    },
    methods: {
      async handleAgregarProyecto() {
        try {
        // Objeto con los datos del proyecto
        const proyectoData = {
            nombre: this.nombreProyecto,
            fechaInicio: this.fechaInicio,
            fechaLimite: this.fechaLimite,
            descripcionProyecto: this.descripcionProyecto
        };
        
        // Envía los datos del proyecto al backend
        const response = await ProjectService.createProject(proyectoData);

        // Verifica si la solicitud fue exitosa
        if (response.status === 201) {
            // Limpiar los campos después de agregar el proyecto
            this.nombreProyecto = '';
            this.fechaInicio = '';
            this.fechaLimite = '';

            console.log('Proyecto agregado con éxito');
        } else {
            console.error('Error al agregar el proyecto:', response.data);
        }
        } catch (error) {
        console.error('Error al agregar el proyecto:', error);
        }
      }
    }
  };
  </script>
  
  <!-- <style scoped>
  /* Aquí puedes agregar tus estilos */
  </style> -->