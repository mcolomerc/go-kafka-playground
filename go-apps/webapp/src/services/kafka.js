import { Kafka } from "../http-common";
 
class KafkaService {
    async status() {
        const response = await Kafka.get("/kafka/v3/clusters");
        const cluster = response.data.data[0].cluster_id; 
        const brokersResp = await Kafka.get(`/kafka/v3/clusters/${cluster}/brokers`);
        const brokers = brokersResp.data.data.length;
        const topicsResponse  = await Kafka.get(`/kafka/v3/clusters/${cluster}/topics`)
        const topics = topicsResponse.data.data.length; 
        const partitions_count = topicsResponse.data.data.reduce((a, b) => {
                return a + b.partitions_count;
              }, 0); 
        const replication_factor = topicsResponse.data.data.reduce((a, b) => {
                return a + b.replication_factor;
              }, 0); 
        return {
          cluster: cluster,
          brokers: brokers,
          topics: topics,
          partitions: partitions_count,
          replication_factor: replication_factor,
        }          
    }
}
export default new KafkaService();

