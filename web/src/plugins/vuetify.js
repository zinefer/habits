import Vue from "vue";
import Vuetify from "vuetify/lib";

Vue.use(Vuetify);

const opts = {
  theme: {
    themes: {
      light: {
        primary: "#546e7a",
        secondary: "#5c6bc0",
        accent: "#819ca9",
        error: "#f44336",
        warning: "#ffeb3b",
        info: "#03a9f4",
        success: "#4caf50"
      }
    }
  }
};

export default new Vuetify(opts);
