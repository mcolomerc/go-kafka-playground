<template>
<p class="text-lg font-bold text-gray-800">Produce Messages</p> 
  <!-- Form request -->
  <div class="block p-6 rounded-lg shadow-lg bg-white flex flex-col md:flex-row pb-4 ">
    <div class="flex flex-col md:flex-row justify-between w-full">
       
        <div class="flex flex-col md:flex-row space-x-10 w-full">
          <div>
            <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
              Number of messages
            </label>
            <input
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="nmessages" type="number" placeholder="10000" v-model="nmessages" required>
          </div>
          <div >
            <label class="block text-gray-700 text-sm font-bold mb-2" for="email">
              Message size (bytes)
            </label>
            <input
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="size" type="number" placeholder="800" v-model="size" required>
          </div>
          <div >
            <label class="block text-gray-700 text-sm font-bold mb-2" for="email">
              Topics by Producer
            </label>
            <input
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="size" type="number" placeholder="1" v-model="topics" required>
          </div>
        </div>
        <div class="flex-1 self-center">
          <button class="btn disabled:opacity-50" @click="execute" :disabled='loading'>Execute</button>
        </div>
      
    </div>
  </div> 
</template>

<script setup>
import ApiService from "../services/api";
import { ref } from 'vue'  
 
const emit = defineEmits(['results']) 
const nmessages = ref(10000);
const size = ref(200);  
const topics = ref(1);  
const loading = ref(false);

const execute = () =>{  
  loading.value = true;
  var data = {
    messages: nmessages.value,
    size: size.value,
    topic_number: topics.value
  };
  ApiService.produce(data)
    .then(response => {  
      emit('results', response.data);
      loading.value = false;
    })
    .catch(e => {
      console.log(e);
      });
    
} 
</script>