<template>
    <div class="flex h-screen overflow-hidden">
        <AsideDash/>
        <div class="w-full mx-auto p-6 bg-white shadow-md rounded-lg overflow-y-auto">
            <h2 class="text-2xl font-bold mb-6">Crear Test</h2>
            <form @submit.prevent="submitForm" class="space-y-4">
                <div>
                <label class="block text-gray-700 text-sm font-bold mb-2" for="title">Título</label>
                <input
                    id="title"
                    type="text"
                    v-model="title"
                    class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    placeholder="Título del test"
                />
                </div>
                <div>
                <label class="block text-gray-700 text-sm font-bold mb-2" for="limitDate">Fecha límite</label>
                <input
                    id="limitDate"
                    type="date"
                    v-model="limitDate"
                    class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                />
                </div>
                <div>
                <label class="block text-gray-700 text-sm font-bold mb-2" for="priority">Prioridad</label>
                <select
                    id="priority"
                    v-model="priority"
                    class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                >
                    <option value="alta">Alta</option>
                    <option value="media">Media</option>
                    <option value="baja">Baja</option>
                </select>
                </div>
                <div>
                <label class="block text-gray-700 text-sm font-bold mb-2" for="fase">Fase</label>
                <select
                    id="fase"
                    v-model="fase"
                    class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                >
                    <option value="análisis">Análisis</option>
                    <option value="diseño">Diseño</option>
                    <option value="implementación">Implementación</option>
                    <option value="pruebas">Pruebas</option>
                    <option value="mantenimiento">Mantenimiento</option>
                </select>
                </div>
                <div>
                <label class="block text-gray-700 text-sm font-bold mb-2" for="evidence">Evidencia</label>
                <input
                    id="evidence"
                    type="url"
                    v-model="evidence"
                    class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    placeholder="URL de evidencia"
                />
                </div>
                <div>
                <label class="block text-gray-700 text-sm font-bold mb-2" for="description">Descripción</label>
                <textarea
                    id="description"
                    v-model="description"
                    class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    placeholder="Descripción del test"
                ></textarea>
                </div>
                <div>
                <label class="block text-gray-700 text-sm font-bold mb-2" for="expectations">Expectativas</label>
                <textarea
                    id="expectations"
                    v-model="expectations"
                    class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    placeholder="Expectativas del test"
                ></textarea>
                </div>
                <div>
                <label class="block text-gray-700 text-sm font-bold mb-2" for="steps">Pasos</label>
                <textarea
                    id="steps"
                    v-model="steps"
                    class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    placeholder="Pasos del test"
                ></textarea>
                </div>
                <div>
                <label class="block text-gray-700 text-sm font-bold mb-2" for="report">Reporte</label>
                <textarea
                    id="report"
                    v-model="report"
                    class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    placeholder="Reporte del test"
                ></textarea>
                </div>
                <div>
                <label class="block text-gray-700 text-sm font-bold mb-2" for="responsable">Responsable</label>
                <input
                    id="responsable"
                    type="text"
                    v-model="responsable"
                    class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                    placeholder="Responsable del test"
                />
                </div>
                <div>
                <label class="block text-gray-700 text-sm font-bold mb-2" for="state">Estado</label>
                <select
                    id="state"
                    v-model="state"
                    class="bg-gray-200 input input-bordered w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                >
                    <option value="sin empezar">Sin empezar</option>
                    <option value="en progreso">En progreso</option>
                    <option value="terminado">Terminado</option>
                </select>
                </div>
                <div class="flex justify-center">
                <button
                    type="submit"
                    class="btn btn-primary group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                >
                    Crear Test
                </button>
                </div>
            </form>
            </div>
    </div>
    
  </template>
  
  <script>
  import { ref } from 'vue';
  import testService from '../services/testService'; 
  import { toast } from 'vue3-toastify';
  import 'vue3-toastify/dist/index.css';
  import { useRoute } from 'vue-router';
  import AsideDash from '../components/AsideDash.vue'
  import { useRouter } from 'vue-router';
  
  export default {
    setup() {
      const title = ref('');
      const limitDate = ref('');
      const priority = ref('');
      const fase = ref('');
      const evidence = ref('');
      const description = ref('');
      const expectations = ref('');
      const steps = ref('');
      const report = ref('');
      const responsable = ref('');
      const state = ref('');
      const route = useRoute();
      const id = route.query.id;
      const router = useRouter();

      const submitForm = async () => {
        try {
          const testData = {
            title: title.value,
            limitDate: limitDate.value,
            priority: priority.value,
            fase: fase.value,
            evidence: evidence.value,
            description: description.value,
            expectations: expectations.value,
            steps: steps.value,
            report: report.value,
            responsable: responsable.value,
            state: state.value,
            idProject: id
          };

          console.log(testData)
  
          const response = await testService.createTest(testData);
  
          if (response.status === 201) {
            title.value = '';
            limitDate.value = '';
            priority.value = '';
            fase.value = '';
            evidence.value = '';
            description.value = '';
            expectations.value = '';
            steps.value = '';
            report.value = '';
            responsable.value = '';
            state.value = '';
  
            toast.success('create test successful!', {
            autoClose: 3000  
          });

          router.push('/project/'+ id);

          } else {
            toast.error('create test failed.', {
            autoClose: 5000 
          });
          }
        } catch (error) {
          console.error('Error al crear el test:', error);
        }
      };
  
      return {
        title,
        limitDate,
        priority,
        fase,
        evidence,
        description,
        expectations,
        steps,
        report,
        responsable,
        state,
        submitForm
      };
    },
    components: {
        AsideDash
    }
  };
  </script>
  