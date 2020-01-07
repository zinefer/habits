<template>
  <v-container fluid :class="{ 'pt-10': isMobile }">
    <v-row dense v-for="habit in habits" :key="habit.ID" class="mb-6">
      <v-spacer />
      <v-col cols="10">
        <HabitCard :habit="habit" :isMobile="isMobile" />
      </v-col>
      <v-spacer />
    </v-row>
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
      loading: true,
      habits: []
    };
  },
  computed: {
    isMobile: function() {
      return screen.width <= 960;
    }
  },
  methods: {
    loadHabits() {
      this.loading = true;
      HabitsApi.get()
        .then(resp => {
          this.habits = resp.data;
        })
        .finally(() => {
          this.loading = false;
        });
    }
  },
  created() {
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
