<template>
  <div class="habits">
    <v-container>
      <v-row dense>
        <v-col cols="12">
          <v-card>
            <v-card-title>Coding</v-card-title>
            <Calendar></Calendar>
            <v-fab-transition>
              <v-btn color="secondary" fab dark absolute bottom right>
                <v-icon>mdi-plus-box</v-icon>
              </v-btn>
            </v-fab-transition>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import Calendar from "@/components/calendar.vue";

export default {
  name: "habits",
  components: {
    Calendar
  }
};
</script>
