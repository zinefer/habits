import axios from "axios";

var konsole = console;

export default {
  get() {
    return axios.get("/api/habits").catch(error => konsole.log(error));
  },
  create(habit) {
    return axios.post("/api/habits", habit).catch(error => konsole.log(error));
  }
};
