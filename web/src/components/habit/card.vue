<template>
  <v-card style="position:relative">
    <v-speed-dial
      v-if="showActions"
      class="options"
      absolute
      v-model="optionsOpen"
      :direction="speedDialDirection"
      transition="slide-x-reverse-transition"
      fab
    >
      <template v-slot:activator>
        <v-btn v-model="optionsOpen" color="primary" fab small>
          <v-icon v-if="optionsOpen">mdi-close</v-icon>
          <v-icon v-else>mdi-settings</v-icon>
        </v-btn>
      </template>
      <v-btn fab dark small color="green" @click.stop="editHabit">
        <v-icon>mdi-pencil</v-icon>
      </v-btn>
      <v-btn
        fab
        dark
        small
        color="red"
        @click.stop="deleteHabit"
        :loading="deleting"
      >
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </v-speed-dial>
    <v-card-title class="pb-2">
      <v-container fluid dense class="pa-0">
        <v-row dense>
          <span class="pl-4 col-sm-8">{{ habit.Name }}</span>
          <v-spacer />
          <div
            :class="{ 'cols-sm-2': true, streak: true, mobile: this.isMobile }"
          >
            <template v-if="streaks == null">
              <Skeleton type="text" width="75" class="pushdown" />
              <div class="divider" />
              <Skeleton type="text" width="75" class="pushdown" />
            </template>
            <template v-else>
              <span
                >{{ streaks.Longest.Streak }} Day <br />
                Longest Streak</span
              >
              <div class="divider" />
              <span
                >{{ streaks.Current.Streak }} Day <br />
                Current Streak</span
              >
            </template>
          </div>
        </v-row>
      </v-container>
    </v-card-title>
    <v-card-text>
      <v-scale-transition hide-on-leave>
        <Skeleton v-if="this.loading" type="habit-calendar-days" />
        <HabitCalendar
          v-else
          :values="this.activities"
          :isMobile="isMobile"
          v-on:showTooltip="showTooltip"
          v-on:hideTooltip="hideTooltip"
        />
      </v-scale-transition>
    </v-card-text>
    <v-fab-transition v-if="showActions">
      <v-btn
        color="secondary"
        fab
        dark
        absolute
        bottom
        right
        :loading="addingNew"
        @click.stop="addNewActivity"
      >
        <v-icon>mdi-plus-box</v-icon>
      </v-btn>
    </v-fab-transition>
  </v-card>
</template>

<script>
import HabitsApi from "@/services/habits";
import ActivitiesApi from "@/services/activities";
import Skeleton from "@/components/skeleton.vue";
import HabitCalendar from "@/components/habit/calendar.vue";

import { EventBus } from "@/event_bus";

export default {
  name: "HabitCard",
  props: ["habit", "isMobile", "showActions"],
  data() {
    return {
      loading: true,
      deleting: false,
      optionsOpen: false,
      addingNew: false,
      activities: [],
      streaks: null
    };
  },
  computed: {
    speedDialDirection: function() {
      return this.isMobile ? "bottom" : "left";
    }
  },
  methods: {
    addNewActivity() {
      this.addingNew = true;

      ActivitiesApi.create(this.habit.ID)
        .then(resp => {
          if (resp.status == 200) {
            this.loadData();
          }
        })
        .finally(() => {
          this.addingNew = false;
        });
    },
    loadData() {
      ActivitiesApi.get(this.habit.ID)
        .then(resp => {
          this.activities = resp.data;
        })
        .finally(() => {
          this.loading = false;
        });

      ActivitiesApi.streaks(this.habit.ID).then(resp => {
        this.streaks = resp.data;
      });
    },
    editHabit() {
      var event = { habit: this.habit };
      EventBus.$emit("editHabit", event);
      this.optionsOpen = false;
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
    this.loadData();
  },
  components: {
    Skeleton,
    HabitCalendar
  }
};
</script>

<style lang="scss">
.streak {
  font-size: 0.65em;

  span {
    text-align: center;
    display: inline-block;
    line-height: 1.2em;
  }

  &.mobile {
    margin: 0 auto;
  }
}
.divider {
  height: 2.5em;
  display: inline-block;
  border-left: thin solid rgba(0, 0, 0, 0.12);
  margin: 0 10px;
  position: relative;
  top: 5px;
}
.pushdown {
  display: inline-block;
  position: relative;
  top: 15px;
}
.options {
  left: -20px;
  top: 15px;
}
</style>
