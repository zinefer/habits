<template>
  <v-app>
    <Navigation />

    <v-content>
      <router-view />
    </v-content>
  </v-app>
</template>

<script>
import Navigation from "@/components/navigation.vue";

export default {
  components: {
    Navigation
  }
};
</script>
