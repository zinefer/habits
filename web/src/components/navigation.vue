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

      <span class="display-1 font-weight-bold" href="/">
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
          {{ currentUser }}
          <v-icon>mdi-menu-down</v-icon>
        </v-btn>
      </template>
      <v-list>
        <v-list-item href="/api/logout">
          <v-list-item-title>Sign out</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-menu>

    <template v-slot:extension v-if="isLoggedIn && showAddHabitButton">
      <v-fab-transition>
        <v-btn
          ref="add_habit"
          @click.stop="showHabitDialog = true"
          color="secondary"
          fab
          absolute
          bottom
        >
          <v-icon>mdi-checkerboard-plus</v-icon>
        </v-btn>
      </v-fab-transition>
    </template>

    <HabitDialog
      :show="showHabitDialog"
      :habit="habit"
      @close="habitDialogClose"
    />
  </v-app-bar>
</template>

<script>
import Login from "@/components/dialogs/login.vue";
import HabitDialog from "@/components/dialogs/habit.vue";

import { EventBus } from "@/event_bus";

export default {
  name: "Navigation",
  components: {
    Login,
    HabitDialog
  },
  data: function() {
    return {
      showLogin: false,
      showHabitDialog: false,
      habit: undefined
    };
  },
  mounted() {
    return EventBus.$on("editHabit", event => {
      this.habit = event.habit;
      this.showHabitDialog = true;
    });
  },
  computed: {
    showAddHabitButton: function() {
      return this.$route && this.$route.name == "Habits";
    },
    isLoggedIn: function() {
      return this.$store.getters.isLoggedIn;
    },
    currentUser: function() {
      return this.$store.getters.currentUser;
    }
  },
  methods: {
    loginClose() {
      this.showLogin = false;
    },
    habitDialogClose() {
      this.habit = undefined;
      this.showHabitDialog = false;
    }
  }
};
</script>
