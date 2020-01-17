<template>
  <div class="home">
    <v-container>
      <v-layout text-center wrap>
        <v-flex xs12>
          <v-img
            :src="require('../assets/logo.svg')"
            class="my-3"
            contain
            height="200"
          ></v-img>
        </v-flex>

        <v-flex mb-4>
          <h1 class="display-2 font-weight-bold mb-3">
            Welcome to Habits
          </h1>
          <p class="subheading font-weight-regular">
            This website can help track streaks in your habits.
          </p>
          <p class="caption">
            Made with <v-icon color="red" small>mdi-heart</v-icon> on
            <a href="https://github.com/zinefer/habits">Github</a>
          </p>
        </v-flex>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "Home"
};
</script>
