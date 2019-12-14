export default {
  state: {
    status: "",
    token: localStorage.getItem("token") || "",
    user: {}
  },
  getters: {
    isLoggedIn: state => !!state.token,
    authStatus: state => state.status
  }
};
