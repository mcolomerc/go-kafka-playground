<script setup>
 
</script>

<template>
  <div class="min-h-full">
  <nav class="bg-gray-800">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <div class="flex items-center">
          <div class="flex-shrink-0"> 
             <img class="w-8" :src="logoUrl" alt="Gopher"> 
          </div>
          <div class="hidden md:block">
            <div class="ml-10 flex items-baseline space-x-4"> 
               <router-link class="bg-gray-900 text-white px-3 py-2 rounded-md text-sm font-medium" to="/">Home</router-link>&nbsp;
               <router-link class="bg-gray-900 text-white px-3 py-2 rounded-md text-sm font-medium" to="/about">About</router-link>
            </div>
          </div>
        </div>
        <div class="hidden md:block">
          <div class="ml-4 flex items-center md:ml-6"> 
          </div>
        </div>
        <div class="-mr-2 flex md:hidden"> 
        </div>
      </div>
    </div> 
    
  </nav> 
  <main>
    <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      <!-- Replace with your content -->
      <router-view />
      <!-- /End replace -->
    </div>
  </main>
</div>
 
</template>
<script setup> 
import logoUrl from './assets/gokafkalogo.png'
</script>
<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
