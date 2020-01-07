import axios from "axios";

var konsole = console;

export default {
  get() {
    return axios.get("/api/habits").catch(error => konsole.log(error));
  },
  create(habit) {
    return axios.post("/api/habits", habit).catch(error => konsole.log(error));
  },
  update(habit) {
    return axios
      .patch("/api/habits/" + habit.ID, habit)
      .catch(error => konsole.log(error));
  },
  delete(habitID) {
    return axios
      .delete("/api/habits/" + habitID)
      .catch(error => konsole.log(error));
  }
};
