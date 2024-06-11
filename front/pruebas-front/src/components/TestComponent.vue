<template>
    <div class="bg-gray-100 rounded-lg p-4" style="width: 1100px;"> <!-- Ancho específico de 400px -->
      <div @click="toggleExpanded" class="cursor-pointer flex justify-between items-center">
        <div>
            <h1 class="text-xl font-bold">{{ test.Title }}</h1>
        </div>
        <div :class="{ 'transform rotate-180': expanded }">▼</div>
      </div>
      <div v-if="expanded" class="mt-2 overflow-auto" style="max-width: 100%; word-wrap: break-word;">
        <p><span class="font-bold">Fecha de publicación:</span> {{ test.PublicationDate }}</p>
        <p><span class="font-bold">Fecha límite</span>: {{ test.LimitDate }}</p>
        <p><span class="font-bold">Prioridad:</span> {{ test.Priority }}</p>
        <p><span class="font-bold">Fase:</span> {{ test.Fase }}</p>
        <p><span class="font-bold">Evidencia:</span> {{ test.Evidence }}</p>
        <p><span class="font-bold">Descripción:</span> {{ test.Description }}</p>
        <p><span class="font-bold">Expectativas:</span> {{ test.Expectations }}</p>
        <p><span class="font-bold">Pasos:</span> {{ test.Steps }}</p>
        <p><span class="font-bold">Reporte:</span> {{ test.Report }}</p>
        <p><span class="font-bold">Responsable:</span> {{ test.Responsable }}</p>
        <p><span class="font-bold">Estado:</span> {{ test.State }}</p>
      </div>
    </div>
  </template>
  
  <script>
  import { ref } from 'vue';
  
  export default {
    props: {
      test: Object 
    },
    setup(props) {
      const expanded = ref(false);
  
      function toggleExpanded() {
        expanded.value = !expanded.value;
      }
  
      return {
        expanded,
        toggleExpanded
      };
    }
  };
  </script>