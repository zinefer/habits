<template>
  <v-container>
    <div
      ref="tooltip"
      class="tooltip v-tooltip__content"
      style="display:none"
    ></div>
    <v-row dense>
      <v-col cols="12" v-for="habit in habits" :key="habit.ID">
        <v-card>
          <v-card-title>{{ habit.Name }}</v-card-title>
          <v-card-text>
            <HabitCalendar
              :habitID="habit.ID"
              v-on:showTooltip="showTooltip"
              v-on:hideTooltip="hideTooltip"
            />
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
  methods: {
    showTooltip(position) {
      var tooltip = this.$refs.tooltip;
      tooltip.innerHTML = "<span>" + position.text + "</span>";
      tooltip.style.display = "initial";
      tooltip.style.top = position.top + "px";
      tooltip.style.left = position.left - tooltip.scrollWidth / 2 + 15 + "px";
    },
    hideTooltip() {
      this.$refs.tooltip.style.display = "none";
    }
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

<style lang="scss">
.tooltip {
  z-index: 100;
  position: absolute;
}
</style>
