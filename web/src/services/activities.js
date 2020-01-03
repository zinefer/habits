import axios from "axios";

var konsole = console;

export default {
  get(habit_id) {
    return axios
      .get("/api/habits/" + habit_id + "/activities")
      .catch(error => konsole.log(error));
  },
  create(habit_id) {
    return axios
      .post("/api/habits" + habit_id + "/activities")
      .catch(error => konsole.log(error));
  }
};
