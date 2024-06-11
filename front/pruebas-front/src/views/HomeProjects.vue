<template>
    <div class="flex">
      <AsideDash />
      <div class="ml-16 w-full">
        <p class="font-sans font-medium text-4xl pb-14 pt-6 text-center">Proyectos</p>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 px-6">
          <Project
            v-for="proyecto in proyectos"
            :key="proyecto.codigo"
            :codigo="proyecto.codigo"
            :nombre="proyecto.nombre"
            :descripcion="proyecto.descripcion"
            :fechaLimite="proyecto.fechaLimite"
            :pruebas="proyecto.pruebas"
          />
        </div>
      </div>
    </div>
  </template>

<script>
    import Project from "../components/Project.vue";
    import AsideDash from '../components/AsideDash.vue'

    export default {
        setup() {
            const proyectos = [
            {
                codigo: "1",
                nombre: "Proyecto 1",
                descripcion: "Este es el proyecto 1",
                fechaInicio: "2021-10-10",
                fechaLimite: "2021-10-10",
                pruebas: [
                {
                    id_prueba: "1",
                    titulo: "Prueba 1",
                    tipo_prueba: "Prueba de caja negra",
                    descripcion: "Esta es la prueba 1",
                    pasos_a_seguir: "1. Ingresar al sistema con el usuario y contraseña del administrador. 2. Seleccionar la opción de proyectos. 3. Seleccionar la opción de crear proyecto. 4. Ingresar los datos del proyecto. 5. Guardar los datos del proyecto. 6. Verificar que el proyecto se haya creado correctamente. 7. Seleccionar la opción de proyectos. 8. Seleccionar la opción de editar proyecto. 9. Ingresar los datos del proyecto. 10. Guardar los datos del proyecto. 11. Verificar que el proyecto se haya editado correctamente. 12. Seleccionar la opción de proyectos. 13. Seleccionar la opción de eliminar proyecto. 14. Verificar que el proyecto se haya eliminado correctamente.",
                    datos_de_prueba: "url couldinary",
                    expectativas: "Se espera que el proyecto se cree, edite y elimine correctamente.",
                    fase: "Terminada",
                    id_responsable: "a1",
                    fecha_inicio: "2021-10-10",
                    fecha_limete: "2021-11-10",
                    prioridad: "Alta prioridad",
                    resultados: "En proceso",
                    evidencia: "url couldinary",
                    informe: "en proceso"
                },
                // Aquí continúa con las pruebas restantes del proyecto 1
                ]
            },
            {
                codigo: "2",
                nombre: "Proyecto 2",
                descripcion: "Este es el proyecto 2",
                fechaInicio: "2021-10-10",
                fechaLimite: "2021-10-10",
                pruebas: 
                    [
                        {
                            id_prueba:"1",
                            titulo: "Prueba 1",
                            tipo_prueba: "Prueba de caja negra",
                            descripcion: "Esta es la prueba 1",
                            pasos_a_seguir: "1. Ingresar al sistema con el usuario y contraseña del administrador. 2. Seleccionar la opción de proyectos. 3. Seleccionar la opción de crear proyecto. 4. Ingresar los datos del proyecto. 5. Guardar los datos del proyecto. 6. Verificar que el proyecto se haya creado correctamente. 7. Seleccionar la opción de proyectos. 8. Seleccionar la opción de editar proyecto. 9. Ingresar los datos del proyecto. 10. Guardar los datos del proyecto. 11. Verificar que el proyecto se haya editado correctamente. 12. Seleccionar la opción de proyectos. 13. Seleccionar la opción de eliminar proyecto. 14. Verificar que el proyecto se haya eliminado correctamente.",
                            datos_de_prueba: "url couldinary",
                            expectativas: "Se espera que el proyecto se cree, edite y elimine correctamente.",
                            fase:   "Terminada",
                            id_responsable: "a1",
                            fecha_inicio: "2021-10-10",
                            fecha_limete: "2021-11-10",
                            prioridad: "Alta prioridad",
                            resultados:  "En proceso",
                            evidencia: "url couldinary",
                            informe: "en proceso"

                        },
                        {
                            id_prueba:"2",
                            titulo: "Prueba 2",
                            tipo_prueba: "Prueba de caja negra",
                            descripcion: "Esta es la prueba 2",
                            pasos_a_seguir: "1. Ingresar al sistema con el usuario y contraseña del administrador. 2. Seleccionar la opción de proyectos. 3. Seleccionar la opción de crear proyecto. 4. Ingresar los datos del proyecto. 5. Guardar los datos del proyecto. 6. Verificar que el proyecto se haya creado correctamente. 7. Seleccionar la opción de proyectos. 8. Seleccionar la opción de editar proyecto. 9. Ingresar los datos del proyecto. 10. Guardar los datos del proyecto. 11. Verificar que el proyecto se haya editado correctamente. 12. Seleccionar la opción de proyectos. 13. Seleccionar la opción de eliminar proyecto. 14. Verificar que el proyecto se haya eliminado correctamente.",
                            datos_de_prueba: "url couldinary",
                            expectativas: "Se espera que el proyecto se cree, edite y elimine correctamente.",
                            fase:   "Terminada",
                            id_responsable: "a1",
                            fecha_inicio: "2021-10-10",
                            fecha_limete: "2021-11-10",
                            prioridad: "Alta prioridad",
                            resultados:  "En proceso",
                            evidencia: "url couldinary",
                            informe: "en proceso"
                        },
                        {
                            id_prueba:"3",
                            titulo: "Prueba 3",
                            tipo_prueba: "Prueba de caja negra",
                            descripcion: "Esta es la prueba 3",
                            pasos_a_seguir: "1. Ingresar al sistema con el usuario y contraseña del administrador. 2. Seleccionar la opción de proyectos. 3. Seleccionar la opción de crear proyecto. 4. Ingresar los datos del proyecto. 5. Guardar los datos del proyecto. 6. Verificar que el proyecto se haya creado correctamente. 7. Seleccionar la opción de proyectos. 8. Seleccionar la opción de editar proyecto. 9. Ingresar los datos del proyecto. 10. Guardar los datos del proyecto. 11. Verificar que el proyecto se haya editado correctamente. 12. Seleccionar la opción de proyectos. 13. Seleccionar la opción de eliminar proyecto. 14. Verificar que el proyecto se haya eliminado correctamente.",
                            datos_de_prueba: "url couldinary",
                            expectativas: "Se espera que el proyecto se cree, edite y elimine correctamente.",
                            fase:   "Terminada",
                            id_responsable: "a1",
                            fecha_inicio: "2021-10-10",
                            fecha_limete: "2021-11-10",
                            prioridad: "Alta prioridad",
                            resultados:  "En proceso",
                            evidencia: "url couldinary",
                            informe: "en proceso"
                        }  
                    ]    
            },
            {
                codigo: "3",
                nombre: "Proyecto 3",
                descripcion: "Este es el proyecto 3",
                fechaLimite: "2021-10-10",
                pruebas: 
                [
                    {
                        id_prueba:"1",
                        titulo: "Prueba 1",
                        tipo_prueba: "Prueba de caja negra",
                        descripcion: "Esta es la prueba 1",
                        pasos_a_seguir: "1. Ingresar al sistema con el usuario y contraseña del administrador. 2. Seleccionar la opción de proyectos. 3. Seleccionar la opción de crear proyecto. 4. Ingresar los datos del proyecto. 5. Guardar los datos del proyecto. 6. Verificar que el proyecto se haya creado correctamente. 7. Seleccionar la opción de proyectos. 8. Seleccionar la opción de editar proyecto. 9. Ingresar los datos del proyecto. 10. Guardar los datos del proyecto. 11. Verificar que el proyecto se haya editado correctamente. 12. Seleccionar la opción de proyectos. 13. Seleccionar la opción de eliminar proyecto. 14. Verificar que el proyecto se haya eliminado correctamente.",
                        datos_de_prueba: "url couldinary",
                        expectativas: "Se espera que el proyecto se cree, edite y elimine correctamente.",
                        fase:   "Terminada",
                        id_responsable: "a1",
                        fecha_inicio: "2021-10-10",
                        fecha_limete: "2021-11-10",
                        prioridad: "Alta prioridad",
                        resultados:  "En proceso",
                        evidencia: "url couldinary",
                        informe: "en proceso"

                    },
                    {
                        id_prueba:"2",
                        titulo: "Prueba 2",
                        tipo_prueba: "Prueba de caja negra",
                        descripcion: "Esta es la prueba 2",
                        pasos_a_seguir: "1. Ingresar al sistema con el usuario y contraseña del administrador. 2. Seleccionar la opción de proyectos. 3. Seleccionar la opción de crear proyecto. 4. Ingresar los datos del proyecto. 5. Guardar los datos del proyecto. 6. Verificar que el proyecto se haya creado correctamente. 7. Seleccionar la opción de proyectos. 8. Seleccionar la opción de editar proyecto. 9. Ingresar los datos del proyecto. 10. Guardar los datos del proyecto. 11. Verificar que el proyecto se haya editado correctamente. 12. Seleccionar la opción de proyectos. 13. Seleccionar la opción de eliminar proyecto. 14. Verificar que el proyecto se haya eliminado correctamente.",
                        datos_de_prueba: "url couldinary",
                        expectativas: "Se espera que el proyecto se cree, edite y elimine correctamente.",
                        fase:   "Terminada",
                        id_responsable: "a1",
                        fecha_inicio: "2021-10-10",
                        fecha_limete: "2021-11-10",
                        prioridad: "Alta prioridad",
                        resultados:  "En proceso",
                        evidencia: "url couldinary",
                        informe: "en proceso"
                    },
                    {
                        id_prueba:"3",
                        titulo: "Prueba 3",
                        tipo_prueba: "Prueba de caja negra",
                        descripcion: "Esta es la prueba 3",
                        pasos_a_seguir: "1. Ingresar al sistema con el usuario y contraseña del administrador. 2. Seleccionar la opción de proyectos. 3. Seleccionar la opción de crear proyecto. 4. Ingresar los datos del proyecto. 5. Guardar los datos del proyecto. 6. Verificar que el proyecto se haya creado correctamente. 7. Seleccionar la opción de proyectos. 8. Seleccionar la opción de editar proyecto. 9. Ingresar los datos del proyecto. 10. Guardar los datos del proyecto. 11. Verificar que el proyecto se haya editado correctamente. 12. Seleccionar la opción de proyectos. 13. Seleccionar la opción de eliminar proyecto. 14. Verificar que el proyecto se haya eliminado correctamente.",
                        datos_de_prueba: "url couldinary",
                        expectativas: "Se espera que el proyecto se cree, edite y elimine correctamente.",
                        fase:   "Terminada",
                        id_responsable: "a1",
                        fecha_inicio: "2021-10-10",
                        fecha_limete: "2021-11-10",
                        prioridad: "Alta prioridad",
                        resultados:  "En proceso",
                        evidencia: "url couldinary",
                        informe: "en proceso"
                    }  
                ]
            },
            {
                codigo: "4",
                nombre: "Proyecto 4",
                descripcion: "Este es el proyecto 4",
                fechaLimite: "2021-10-10",
                pruebas: 
                [
                    {
                        id_prueba:"1",
                        titulo: "Prueba 1",
                        tipo_prueba: "Prueba de caja negra",
                        descripcion: "Esta es la prueba 1",
                        pasos_a_seguir: "1. Ingresar al sistema con el usuario y contraseña del administrador. 2. Seleccionar la opción de proyectos. 3. Seleccionar la opción de crear proyecto. 4. Ingresar los datos del proyecto. 5. Guardar los datos del proyecto. 6. Verificar que el proyecto se haya creado correctamente. 7. Seleccionar la opción de proyectos. 8. Seleccionar la opción de editar proyecto. 9. Ingresar los datos del proyecto. 10. Guardar los datos del proyecto. 11. Verificar que el proyecto se haya editado correctamente. 12. Seleccionar la opción de proyectos. 13. Seleccionar la opción de eliminar proyecto. 14. Verificar que el proyecto se haya eliminado correctamente.",
                        datos_de_prueba: "url couldinary",
                        expectativas: "Se espera que el proyecto se cree, edite y elimine correctamente.",
                        fase:   "Terminada",
                        id_responsable: "a1",
                        fecha_inicio: "2021-10-10",
                        fecha_limete: "2021-11-10",
                        prioridad: "Alta prioridad",
                        resultados:  "En proceso",
                        evidencia: "url couldinary",
                        informe: "en proceso"

                    },
                    {
                        id_prueba:"2",
                        titulo: "Prueba 2",
                        tipo_prueba: "Prueba de caja negra",
                        descripcion: "Esta es la prueba 2",
                        pasos_a_seguir: "1. Ingresar al sistema con el usuario y contraseña del administrador. 2. Seleccionar la opción de proyectos. 3. Seleccionar la opción de crear proyecto. 4. Ingresar los datos del proyecto. 5. Guardar los datos del proyecto. 6. Verificar que el proyecto se haya creado correctamente. 7. Seleccionar la opción de proyectos. 8. Seleccionar la opción de editar proyecto. 9. Ingresar los datos del proyecto. 10. Guardar los datos del proyecto. 11. Verificar que el proyecto se haya editado correctamente. 12. Seleccionar la opción de proyectos. 13. Seleccionar la opción de eliminar proyecto. 14. Verificar que el proyecto se haya eliminado correctamente.",
                        datos_de_prueba: "url couldinary",
                        expectativas: "Se espera que el proyecto se cree, edite y elimine correctamente.",
                        fase:   "Terminada",
                        id_responsable: "a1",
                        fecha_inicio: "2021-10-10",
                        fecha_limete: "2021-11-10",
                        prioridad: "Alta prioridad",
                        resultados:  "En proceso",
                        evidencia: "url couldinary",
                        informe: "en proceso"
                    },
                    {
                        id_prueba:"3",
                        titulo: "Prueba 3",
                        tipo_prueba: "Prueba de caja negra",
                        descripcion: "Esta es la prueba 3",
                        pasos_a_seguir: "1. Ingresar al sistema con el usuario y contraseña del administrador. 2. Seleccionar la opción de proyectos. 3. Seleccionar la opción de crear proyecto. 4. Ingresar los datos del proyecto. 5. Guardar los datos del proyecto. 6. Verificar que el proyecto se haya creado correctamente. 7. Seleccionar la opción de proyectos. 8. Seleccionar la opción de editar proyecto. 9. Ingresar los datos del proyecto. 10. Guardar los datos del proyecto. 11. Verificar que el proyecto se haya editado correctamente. 12. Seleccionar la opción de proyectos. 13. Seleccionar la opción de eliminar proyecto. 14. Verificar que el proyecto se haya eliminado correctamente.",
                        datos_de_prueba: "url couldinary",
                        expectativas: "Se espera que el proyecto se cree, edite y elimine correctamente.",
                        fase:   "Terminada",
                        id_responsable: "a1",
                        fecha_inicio: "2021-10-10",
                        fecha_limete: "2021-11-10",
                        prioridad: "Alta prioridad",
                        resultados:  "En proceso",
                        evidencia: "url couldinary",
                        informe: "en proceso"
                    }  
                ]
            }
            // Aquí continúan los demás proyectos
            ];
            return {
                proyectos
            }
        },
        components: {
            AsideDash,
            Project,
        }
    }
</script>

<style scoped>
    /* Estilos específicos para este componente */
    /* Puedes añadir tus estilos personalizados aquí */
</style>
  