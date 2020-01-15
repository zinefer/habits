<template>
  <v-dialog v-model="dialogVisible" max-width="700">
    <v-card>
      <v-card-title class="headline justify-center">
        {{ haveHabitID ? "Edit" : "Add New" }} Habit
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-text-field
                label="Name*"
                v-model="habit.Name"
                required
              ></v-text-field>
            </v-col>
          </v-row>
        </v-container>
        <small>*indicates required field</small>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="primary" @click="close">Close</v-btn>
        <v-btn color="secondary" @click="save" :loading="loading">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import HabitsApi from "@/services/habits";

import { EventBus } from "@/event_bus";

export default {
  name: "HabitDialog",
  props: {
    show: {
      type: Boolean,
      required: true
    },
    habit: {
      type: Object,
      default: () => ({ Name: "" })
    }
  },
  data() {
    return {
      loading: false
    };
  },
  methods: {
    save() {
      this.loading = true;
      if (this.haveHabitID) {
        HabitsApi.update(this.habit)
          .then(resp => {
            if (resp.status == 200) {
              this.dialogVisible = false;
              this.habit = {};
              EventBus.$emit("reloadHabits");
            } else {
              alert("Unknown error updating habit");
            }
          })
          .finally(() => {
            this.loading = false;
          });
      } else {
        HabitsApi.create(this.habit)
          .then(resp => {
            if (resp.status == 201) {
              this.dialogVisible = false;
              this.habit = {};
              EventBus.$emit("reloadHabits");
            } else {
              alert("Unknown error creating habit");
            }
          })
          .finally(() => {
            this.loading = false;
          });
      }
    },
    close() {
      this.dialogVisible = false;
      this.habit = {};
    }
  },
  computed: {
    haveHabitID: function() {
      return typeof this.habit.ID != "undefined";
    },
    dialogVisible: {
      get: function() {
        return this.show;
      },
      set: function(value) {
        if (!value) {
          this.$emit("close");
        }
      }
    }
  }
};
</script>
