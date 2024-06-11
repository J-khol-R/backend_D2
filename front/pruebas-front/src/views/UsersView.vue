<template>
    <div class="flex">
        <AsideDash/>
        <div class=" ml-16 w-full justify-center items-center">
            <p class="font-sans font-medium text-4xl pb-14 pt-6 text-center">Usuarios</p>
            <div class="flex justify-center items-center">
                <div v-for="user in users" :key="user.name" class="max-w-sm rounded overflow-hidden shadow-lg bg-white mx-4 my-4 transition-shadow duration-300">
                    <div class="px-6 py-4">
                        <h2 class="font-bold text-xl mb-2 text-gray-800">{{ user.name }}</h2>
                        <p class="text-gray-700 text-base">Rol: {{ user.role }}</p>
                        <select v-model="user.newRole" class="mt-4" @change="handleChange(user)">
                            <option v-for="role in roles" :key="role" :value="role">{{ role }}</option>
                        </select>
                    </div>
                </div>
            </div>
        </div>
    
    </div>
</template>
  
<script>
    import { ref } from 'vue';
    import AsideDash from '../components/AsideDash.vue';
    import userService from '../services/userService';

  export default {
    setup(){
        const roles = ['Admin', 'Developer', 'Tester'];
        const users = [
            {
                name: 'Juan',
                role: 'Admin'
            },
            {
                name: 'Pedro',
                role: 'User'
            },
            {
                name: 'Maria',
                role: 'User'
            }
        ];//ref([]);
        const getUsers = async () => {
            try {
                const response = await userService.getAllUsers();
                console.log(response.data);
                users.value = response.data;
            } catch (error) {
                console.error(error);
            }
        };
        const handleChange = async (user) => {
            try {
                // Enviar la petición al backend para cambiar el rol
                console.log('Cambiando rol:', user);
                await userService.updateUser(user.id, user.newRole);
                user.role = user.newRole;
                user.newRole = '';
            } catch (error) {
                console.error(error);
            }
        };
        return {
            users,
            roles,
            // handleEdit,
            // handleDelete,
            getUsers,
            handleChange,
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