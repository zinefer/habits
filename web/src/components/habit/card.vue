<template>
  <div class="mb-8">
    <v-card style="position:relative">
      <v-speed-dial
        style="left:-20px;top:10px"
        absolute
        v-model="optionsOpen"
        direction="left"
        transition="slide-x-reverse-transition"
        fab
      >
        <template v-slot:activator>
          <v-btn v-model="optionsOpen" color="primary" fab small>
            <v-icon v-if="optionsOpen">mdi-close</v-icon>
            <v-icon v-else>mdi-settings</v-icon>
          </v-btn>
        </template>
        <v-btn fab dark small color="green">
          <v-icon>mdi-pencil</v-icon>
        </v-btn>
        <v-btn
          fab
          dark
          small
          color="red"
          @click="deleteHabit"
          :loading="deleting"
        >
          <v-icon>mdi-delete</v-icon>
        </v-btn>
      </v-speed-dial>
      <v-card-title>
        <span class="pl-3">{{ habit.Name }}</span>
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
import HabitsApi from "@/services/habits";
import ActivitiesApi from "@/services/activities";
import Skeleton from "@/components/skeleton.vue";
import HabitCalendar from "@/components/habit/calendar.vue";

import { EventBus } from "@/event_bus";

export default {
  name: "HabitCard",
  props: ["habit"],
  data() {
    return {
      loading: true,
      deleting: false,
      optionsOpen: false,
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
    deleteHabit() {
      this.deleting = true;
      HabitsApi.delete(this.habit.ID)
        .then(resp => {
          if (resp.status == 200) {
            EventBus.$emit("reloadHabits");
          } else {
            alert("Error deleting habit");
          }
        })
        .finally(() => {
          this.deleting = false;
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
