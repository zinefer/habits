import Cookies from "js-cookie";
import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Habits from "../views/habits.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    meta: {
      redirectHabitsIfAuthed: true
    }
  },
  {
    path: "/habits/:user",
    name: "User Habits",
    component: Habits
  },
  {
    path: "/habits",
    name: "Habits",
    component: Habits,
    meta: {
      requiresAuth: true
    }
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

router.beforeEach((to, from, next) => {
  var loggedIn = Cookies.get("current_user") != null;

  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!loggedIn) {
      next({
        path: "/"
      });
    }
  }

  if (to.matched.some(record => record.meta.redirectHabitsIfAuthed)) {
    if (loggedIn) {
      next({
        path: "/habits"
      });
    }
  }

  next();
});

export default router;
