<template>
    <p class="text-lg font-bold text-gray-800 mt-10">Results</p>

    <div class="block p-6 rounded-lg shadow-lg bg-white flex flex-col md:flex-row pb-4 ">  
        <table class="table-fixed w-full">
            <thead>
                <tr class="bg-gray-100">
                    <th class="px-4 py-2">Producer</th>
                    <th class="whitespace-nowrap w-1/2 px-4 py-2">Topics</th>
                    <th class="px-4 py-2">Duration</th>
                    <th class="px-4 py-2">Messages</th>
                    <th class="px-4 py-2">Size (Bytes)</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="result in results">
                    <td class="text-left border px-4 py-2">{{ result.producer }}</td>
                    <td class="text-left border px-4 py-2">
                        <p class="max-w-md">{{ result.topics }}</p>
                    </td>
                    <td class="border px-4 py-2">{{ result.duration }}</td>
                    <td class="border px-4 py-2">{{ result.number_messages }} </td>
                    <td class="border px-4 py-2">{{ result.size_message }} </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<script setup>

const props = defineProps({
    results: Array
})
 
defineExpose({
    props
})

</script>