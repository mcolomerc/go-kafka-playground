<template>
  <p class="text-xl font-bold text-gray-800">Golang Kafka Libraries Benchmark</p>
  <div class="block p-6 rounded-lg shadow-lg bg-white mb-10">
    <div class="flex pb-4 justify-items-start space-x-20">
      <div class="flex justify-items-start space-x-10">
        <div class="flex self-center">
          <img :src="kafkaUrl" width="48">
        </div>
        <div class="flex flex-col items-start">
          <p class="text-base font-bold text-gray-800">Cluster: <span class="font-normal"> {{ cluster }}</span></p>
          <p class="text-base font-bold text-gray-800">Brokers: <span class="font-normal"> {{ brokers }} </span></p>
          <p class="text-base font-bold text-gray-800">Topics: <span class="font-normal"> {{ topics }} </span></p>
          <p class="text-base font-bold text-gray-800">Partitions: <span class="font-normal"> {{ partitions }}</span></p>
        </div>

      </div>
      <div class="flex space-x-10 justify-items-start">
        <div class="flex self-center">
          <img :src="gopherUrl" width="48">
        </div>
        <div class="flex flex-col justify-items-start">
          <a class="flex" href="https://github.com/confluentinc/confluent-kafka-go"> · Confluent Kafka Golang <span> - cgo based wrapper around librdkafka</span></a>
          <a class="flex" href="https://github.com/Shopify/sarama"> · Shopify Sarama <span> - MIT-licensed Go client library for Apache Kafka.</span> </a>
          <a class="flex" href="https://github.com/twmb/franz-go"> · Franz Go <span> - an all-encompassing Apache Kafka client </span></a>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'

import KafkaService from "../services/kafka";
import kafkaUrl from '../assets/kafka.png'
import gopherUrl from '../assets/gopher.svg'

const props = defineProps({
  update: { type: Boolean, required: false, default: false },
})

const cluster = ref('');
const brokers = ref(0);
const topics = ref(0);
const partitions = ref(0); 

const updateValues = () => { 
  KafkaService.status().then(response => {
    cluster.value = response.cluster;
    brokers.value = response.brokers;
    topics.value = response.topics;
    partitions.value = response.partitions; 
  }).catch(e => {
    console.log(e);
  });
}

onMounted(() => { 
  setInterval(() => {
		updateValues();
	}, 30000)
})

updateValues();

</script>