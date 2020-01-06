<template>
  <v-container>
    <v-row dense>
      <v-col cols="12">
        <HabitCard
          v-for="habit in habits"
          :key="habit.ID"
          :habit="habit"
          v-on:showTooltip="showTooltip"
          v-on:hideTooltip="hideTooltip"
        />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import HabitsApi from "@/services/habits";

import HabitCard from "@/components/habit/card.vue";

import { EventBus } from "@/main";

export default {
  name: "Habits",
  data() {
    return {
      loading: true,
      habits: []
    };
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
