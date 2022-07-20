import { App } from "../http-common";

class ApiService { 
  produce(data) {
    return App.post("/produce", data);
  } 
}
export default new ApiService();