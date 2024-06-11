<template>
    <div class="flex">
        <AsideDash/>
        <div class="container mx-auto py-8" v-if="project">
            <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
            <h1 class="text-2xl font-bold mb-4">{{ project.Name }}</h1>
            <p class="text-gray-700 mb-4">{{ project.Description }}</p>
            <p><span class="font-bold">Fecha de inicio:</span> {{ project.StartDate }}</p>
            <p><span class="font-bold">Fecha de finalizacion:</span> {{ project.FinishDate }}</p>
            <div class="mb-4">
                <h2 class="text-2xl font-bold mb-2">Tests</h2>
                <TestComponent class="m-2" v-for="test in project.Tests" :key="test.id" :test="test" />
            </div>
            <div class="flex justify-between items-center">
                <button @click="goCreateTest" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                    Crear Test</button>
                <div>
                <button @click="goEditProject" class="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded mr-2">Editar Proyecto</button>
                <button @click="deleteProject()" class="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">Borrar Proyecto</button>
                </div>
            </div>
            </div>
        </div>
    </div>
</template>
  
<script>
import { reactive, onMounted } from 'vue';
import projectService from '../services/projectService'; 
import { useRoute, useRouter } from 'vue-router';
import AsideDash from '../components/AsideDash.vue'
import TestComponent from '../components/TestComponent.vue';

export default {
    setup() {
        const project = reactive({});
        const route = useRoute();
        const router = useRouter();

        onMounted(async () => {
        try {
            const projectId = route.params.id;
            const response = await projectService.getProjectById(projectId);
            Object.assign(project, response.data);

            console.log(project.value)
        } catch (error) {
            console.error('Error al cargar el proyecto:', error);
        }
        });

        const goCreateTest = () => {
            router.push({ name: 'CreateTest', query: { id: route.params.id } });
        };
        
        const goEditProject = () => {
            router.push({ name: 'ProjectEdit', params: {
                Name: project.Name,
                StartDate: project.StartDate,
                FinishDate: project.FinishDate,
                Description: project.Description,
                Id: project.Id
                } });
        };

        // const goDeleteProyect = () => {
        //     router.push({ name: 'CreateTest', query: { id: route.params.id } });
        // };

        return {
        project,
        goCreateTest,
        goEditProject
        };
    },
    components: {
        AsideDash,
        TestComponent
    }
};
</script>
