import axios from "axios";

var konsole = console;

axios.interceptors.response.use(null, function(error) {
  if (error.response.status === 401) {
    location.reload();
  }
  return Promise.reject(error);
});

export default {
  get() {
    return axios.get("/api/habits").catch(error => konsole.log(error));
  },
  getByUser(user) {
    return axios.get("/api/habits/" + user).catch(error => konsole.log(error));
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
