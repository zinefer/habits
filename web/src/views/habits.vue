<template>
  <v-container fluid :class="{ 'pt-10': pushContentDown, 'col-10': true }">
    <v-alert
      color="primary"
      dark
      icon="mdi-food-apple"
      border="top"
      prominent
      v-if="
        !loading && !error && habits.length == 0 && isLoggedIn && user == null
      "
    >
      Welcome! To start tracking a habit click on the
      <v-avatar color="secondary" size="32">
        <v-icon small dark>mdi-checkerboard-plus</v-icon>
      </v-avatar>
      button
    </v-alert>
    <v-alert
      type="warning"
      v-if="!loading && !error && habits.length == 0 && user != null"
    >
      {{ user }} is not tracking any habits
    </v-alert>
    <v-alert type="error" v-if="error && user != null">
      Error retrieving Habits for {{ user }}
    </v-alert>
    <HabitCard
      :habit="habit"
      :showActions="isLoggedIn && user == null"
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
