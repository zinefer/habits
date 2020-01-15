import Cookies from "js-cookie";

window.Cookies = Cookies;

export default {
  state: {
    status: "",
    currentUser: Cookies.get("current_user") || "",
    user: {}
  },
  getters: {
    isLoggedIn: state => !!state.currentUser,
    currentUser: state => state.currentUser,
    authStatus: state => state.status
  }
};
