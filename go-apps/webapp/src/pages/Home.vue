<template> 
 <Info></Info>
 <RequestForm @results="onResults"></RequestForm>
 <div v-if="results.length > 0"> 
   <Response :results="results" ></Response>
</div>
</template>
<script setup>
 import Info from '../components/Info.vue'
 import RequestForm from '../components/RequestForm.vue'
 import Response from '../components/Response.vue'
 import { ref } from 'vue'

const results = ref([])  
 
            
 const onResults = (response) => { 
   results.value = response;   
 }
</script>