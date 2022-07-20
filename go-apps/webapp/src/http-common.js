import axios from "axios";

const App =  axios.create({
  baseURL: "/api",
  headers: {
    "Content-type": "application/json",
    'Access-Control-Allow-Origin': '*' 
  }
});

const Kafka =  axios.create({
    baseURL: "http://localhost:8090",
    headers: {
      "Content-type": "application/json"
    }
  });

  export {
    App,
    Kafka
  };