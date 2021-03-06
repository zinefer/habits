import axios from "axios";

var konsole = console;

export default {
  get(habit_id) {
    return axios
      .get("/api/habits/" + habit_id + "/activities?timezone=" + timezone())
      .catch(error => konsole.log(error));
  },
  streaks(habit_id) {
    return axios
      .get("/api/habits/" + habit_id + "/activities/streaks")
      .catch(error => konsole.log(error));
  },
  create(habit_id) {
    return axios
      .post("/api/habits/" + habit_id + "/activities", { TimeZone: timezone() })
      .catch(error => konsole.log(error));
  }
};

function timezone() {
  return Math.round((new Date().getTimezoneOffset() / 60) * -1);
}
