<template>
  <v-dialog v-model="dialogVisible" max-width="290">
    <v-card>
      <v-card-title class="headline justify-center">Sign in</v-card-title>

      <v-card-text>
        <v-btn
          @click="dialog = false"
          class="mb-2"
          color="#444444"
          dark
          href="/auth/github"
        >
          <v-icon class="pr-2">mdi-github-circle</v-icon>
          Continue with GitHub
        </v-btn>
        <v-btn
          @click="dialog = false"
          class="mb-2"
          color="#DD4B39"
          dark
          href="/auth/google"
        >
          <v-icon class="pr-2">mdi-google</v-icon>
          Continue with Google
        </v-btn>
        <v-btn
          @click="dialog = false"
          color="#3B5998"
          dark
          href="/auth/facebook"
        >
          <v-icon class="pr-1">mdi-facebook-box</v-icon>
          Continue with Facebook
        </v-btn>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "Login",
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
