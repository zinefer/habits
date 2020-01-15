<template>
  <v-container fluid :class="{ 'pt-10': pushContentDown, 'col-10': true }">
    <v-alert
      color="primary"
      dark
      icon="mdi-vuetify"
      border="top"
      prominent
      v-if="!loading && habits.length == 0 && isLoggedIn"
    >
      Welcome! To start tracking a habit click on the
      <v-avatar color="secondary" size="32">
        <v-icon dark>mdi-checkerboard-plus</v-icon>
      </v-avatar>
      button on the left.
    </v-alert>
    <v-alert type="error" v-if="error && user != nil">
      Error retrieving Habits for {{ user }}
    </v-alert>
    <HabitCard
      :habit="habit"
      :isMobile="isMobile"
      v-for="habit in habits"
      :key="habit.ID"
      class="mb-8"
    />
  </v-container>
</template>

<script>
import HabitsApi from "@/services/habits";

import HabitCard from "@/components/habit/card.vue";

import { EventBus } from "@/event_bus";

export default {
  name: "Habits",
  data() {
    return {
      error: false,
      loading: true,
      habits: []
    };
  },
  computed: {
    user: function() {
      return this.$route.params.user;
    },
    pushContentDown: function() {
      return this.isMobile && this.isLoggedIn;
    },
    isMobile: function() {
      return screen.width <= 960;
    },
    isLoggedIn: function() {
      return this.$store.getters.isLoggedIn;
    }
  },
  methods: {
    loadHabits() {
      this.loading = true;

      if (this.user != null) {
        HabitsApi.getByUser(this.user)
          .then(resp => {
            this.habits = resp.data;
          })
          .catch(() => {
            this.error = true;
          })
          .finally(() => {
            this.loading = false;
          });
      } else {
        HabitsApi.get()
          .then(resp => {
            this.habits = resp.data;
          })
          .finally(() => {
            this.loading = false;
          });
      }
    }
  },
  mounted() {
    this.loadHabits();

    EventBus.$on("reloadHabits", () => {
      this.loadHabits();
    });
  },
  components: {
    HabitCard
  }
};
</script>
