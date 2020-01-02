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
        <v-list-item href="/logout">
          <v-list-item-title>Sign out</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>

    <template v-slot:extension v-if="isLoggedIn">
      <v-fab-transition>
        <v-btn v-show="!hidden" color="secondary" fab absolute bottom>
          <v-icon>mdi-checkerboard-plus</v-icon>
        </v-btn>
      </v-fab-transition>
    </template>
    <NewHabit :show="showNewHabit" @close="newHabitClose" />
  </v-app-bar>
</template>

<script>
import Login from "@/components/dialogs/login.vue";
import NewHabit from "@/components/dialogs/login.vue";

export default {
  name: "Navigation",
  components: {
    Login,
    NewHabit
  },
  data: function() {
    return {
      showLogin: false,
      showNewHabit: false
    };
  },
  computed: {
    isLoggedIn: function() {
      return this.$store.getters.isLoggedIn;
    }
  },
  methods: {
    loginClose() {
      this.showLogin = false;
    },
    newHabitClose() {
      this.showNewHabit = false;
    }
  }
};
</script>
