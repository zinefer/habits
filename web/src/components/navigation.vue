<template>
  <v-app-bar app color="primary" dark>
    <div class="d-flex align-center">
      <v-img
        alt="Habits Logo"
        class="shrink mr-2"
        contain
        :src="require('../assets/logo.svg')"
        transition="scale-transition"
        width="40"
      />

      <span class="display-1 font-weight-bold">
        Habits
      </span>
    </div>

    <v-spacer></v-spacer>

    <v-btn color="secondary" @click.stop="showLogin = true" v-if="!isLoggedIn">
      Sign in
    </v-btn>
    <Login :show="showLogin" @close="loginClose" />

    <v-menu offset-y v-if="isLoggedIn">
      <template v-slot:activator="{ on }">
        <v-btn v-on="on" outlined>
          <v-icon class="pr-2">mdi-account-circle</v-icon>
          Zinefer
          <v-icon>mdi-menu-down</v-icon>
        </v-btn>
      </template>
      <v-list>
        <v-list-item @click.stop="">
          <v-list-item-title>Sign out</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>
  </v-app-bar>
</template>

<script>
import Login from "@/components/login.vue";

export default {
  name: "Navigation",
  components: {
    Login
  },
  data: function() {
    return {
      showLogin: false
    };
  },
  computed: {
    isLoggedIn: function() {
      return this.$store.auth.isLoggedIn;
    }
  },
  methods: {
    loginClose() {
      this.showLogin = false;
    }
  }
};
</script>
