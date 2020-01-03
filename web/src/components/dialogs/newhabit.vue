<template>
  <v-dialog v-model="dialogVisible" max-width="700">
    <v-card>
      <v-card-title class="headline justify-center">Add New Habit</v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-text-field label="Name*" required></v-text-field>
            </v-col>
          </v-row>
        </v-container>
        <small>*indicates required field</small>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="primary" @click="dialogVisible = false">Close</v-btn>
        <v-btn color="secondary" @click="dialogVisible = false">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "NewHabit",
  props: ["show"],
  computed: {
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
