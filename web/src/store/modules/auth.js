import Cookies from "js-cookie";

window.Cookies = Cookies;

export default {
  state: {
    status: "",
    token: Cookies.get("current_user") || "",
    user: {}
  },
  getters: {
    isLoggedIn: state => !!state.token,
    authStatus: state => state.status
  }
};
