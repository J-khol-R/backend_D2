<template>
    <div class="flex">
        <AsideDash/>
        <tr>
        <td>{{user.nombre }}</td>
        <td>{{ user.rol }}</td>
        <td>
            <!-- Botón para editar -->
            <button class="btn btn-outline btn-accent" @click="handleEdit">
            Editar
            </button>
            
            <!-- Botón para eliminar -->
            <button class="btn btn-outline btn-warning" @click="handleDelete">
            Eliminar
            </button>
        </td>
        </tr>
    </div>
</template>
  
<script>
    import { ref } from 'vue';
    import AsideDash from '../components/AsideDash.vue';
    import userService from '../services/userService';

  export default {
    setup(){
        const users = ref([]);
        const getUsers = async () => {
            try {
                const response = await userService.getAllUsers();
                console.log(response.data);
                users.value = response.data;
            } catch (error) {
                console.error(error);
            }
        };

        // const handleEdit = () {
        //     // Lógica para manejar la edición del elemento (a implementar)
        //     //console.log(`Editar elemento con ID: ${this.el.id}`);
        // };

        // const handleDelete = () {
        //     // Lógica para manejar la eliminación del elemento (a implementar)
        // // console.log(`Eliminar elemento con ID: ${this.el.id}`);
        // };
        return {
            users,
            handleEdit,
            handleDelete,
            getUsers
        };

    },
    components: {
      AsideDash,
    }
  };
</script>
  
<style scoped>
/* Estilos específicos para este componente */
/* Puedes añadir tus estilos personalizados aquí */
</style>