<template>
  <v-container>
    <v-row dense>
      <v-col cols="12" v-for="habit in habits" :key="habit.id">
        <v-card>
          <v-card-title>{{ habit.Name }}</v-card-title>
          <v-card-text>
            <HabitCalendar :habit_id="habit.id" />
          </v-card-text>
          <v-fab-transition>
            <v-btn color="secondary" fab dark absolute bottom right>
              <v-icon>mdi-plus-box</v-icon>
            </v-btn>
          </v-fab-transition>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import HabitsApi from "@/services/habits";

import HabitCalendar from "@/components/habit_calendar.vue";

export default {
  name: "Habits",
  data() {
    return {
      loading: true,
      habits: []
    };
  },
  created() {
    HabitsApi.get()
      .then(resp => {
        this.habits = resp.data;
      })
      .finally(() => {
        this.loading = false;
      });
  },
  components: {
    HabitCalendar
  }
};
</script>
