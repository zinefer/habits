<template>
  <div class="mb-6">
    <v-card>
      <v-card-title>
        <span>{{ habit.Name }}</span>
      </v-card-title>
      <v-card-text>
        <v-scale-transition hide-on-leave>
          <Skeleton v-if="this.loading" type="habit-calendar-days" />
          <HabitCalendar
            v-else
            v-bind:values="this.activities"
            v-on:showTooltip="showTooltip"
            v-on:hideTooltip="hideTooltip"
          />
        </v-scale-transition>
      </v-card-text>
      <v-fab-transition>
        <v-btn
          color="secondary"
          fab
          dark
          absolute
          bottom
          right
          :loading="addingNew"
          @click="addNewActivity"
        >
          <v-icon>mdi-plus-box</v-icon>
        </v-btn>
      </v-fab-transition>
    </v-card>
  </div>
</template>

<script>
import ActivitiesApi from "@/services/activities";
import Skeleton from "@/components/skeleton.vue";
import HabitCalendar from "@/components/habit/calendar.vue";

import { EventBus } from "@/main";

export default {
  name: "HabitCard",
  props: ["habit"],
  data() {
    return {
      loading: true,
      addingNew: false,
      activities: []
    };
  },
  methods: {
    addNewActivity() {
      this.addingNew = true;

      ActivitiesApi.create(this.habit.ID)
        .then(resp => {
          if (resp.status == 200) {
            this.getActivities();
          }
        })
        .finally(() => {
          this.addingNew = false;
        });
    },
    getActivities() {
      ActivitiesApi.get(this.habit.ID)
        .then(resp => {
          this.activities = resp.data;
        })
        .finally(() => {
          this.loading = false;
        });
    },
    showTooltip(event) {
      EventBus.$emit("showTooltip", event);
    },
    hideTooltip() {
      EventBus.$emit("hideTooltip");
    }
  },
  mounted() {
    this.getActivities();
  },
  components: {
    Skeleton,
    HabitCalendar
  }
};
</script>
