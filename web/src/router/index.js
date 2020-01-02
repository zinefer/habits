import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Habits from "../views/habits.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: Home
  },
  {
    path: "/habits",
    name: "habits",
    component: Habits
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
