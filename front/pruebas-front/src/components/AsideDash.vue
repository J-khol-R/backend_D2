<template>
    <aside class="w-64 bg-gray-800 h-screen flex flex-col justify-between">
      <!-- Encabezado del aside -->
      <div class="p-4 text-white">
        <h2 class="text-xl font-semibold">Menú</h2>
      </div>
  
      <!-- Secciones -->
      <nav class="flex-1">
        <ul class="space-y-4">
          <!-- Sección de proyectos -->
          <li>
            <router-link to="/projects" class="flex items-center p-4 text-gray-300 hover:text-white">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path d="M10 12a2 2 0 100-4 2 2 0 000 4z" />
                <path fill-rule="evenodd" d="M3 5a2 2 0 012-2h10a2 2 0 012 2v1h2a1 1 0 110 2h-2v7a2 2 0 01-2 2H5a2 2 0 01-2-2V8H1a1 1 0 010-2h2V5zm2-1a1 1 0 00-1 1v10a1 1 0 001 1h10a1 1 0 001-1V5a1 1 0 00-1-1H5z" clip-rule="evenodd" />
              </svg>
              Proyectos
            </router-link>
          </li>
  
          <!-- Sección de crear proyecto -->
          <li>
            <router-link to="/create-project" class="flex items-center p-4 text-gray-300 hover:text-white">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M11 2.997A1 1 0 0112 2h6a1 1 0 011 1v6a1 1 0 01-1 1h-2.586a1 1 0 00-.707.293l-5.414 5.414a.999.999 0 01-1.414 0l-3.414-3.414A1 1 0 012 12.414V4a1 1 0 011-1h6a1 1 0 01.997.999zM14.707 4H12v2.707l4.707-4.707zM14 16H5a1 1 0 01-1-1v-9h2v7a1 1 0 001 1h7v2z" clip-rule="evenodd" />
              </svg>
              Crear Proyecto
            </router-link>
          </li>
  
          <!-- Sección de usuarios -->
          <li>
            <router-link to="/users" class="flex items-center p-4 text-gray-300 hover:text-white">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2" viewBox="0 0 20 20" fill="currentColor">
                <path d="M13 7a2 2 0 11-4 0 2 2 0 014 0zM9 9a3 3 0 013-3h1a1 1 0 011 1v8a1 1 0 01-1 1H8a1 1 0 01-1-1V7a1 1 0 011-1h1zm1 10a1 1 0 100-2 1 1 0 000 2zM7 5a3 3 0 11-6 0 3 3 0 016 0z" />
              </svg>
              Usuarios
            </router-link>
          </li>
        </ul>
      </nav>
  
      <!-- Pie del aside -->
      <div class="p-4 text-gray-300 mt-auto">
        <p>© 2024 </p>
      </div>
    </aside>
</template>
  

  
