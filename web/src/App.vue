<template>
  <v-app>
    <Navigation />
    <v-content>
      <div
        ref="tooltip"
        class="tooltip v-tooltip__content"
        style="display:none"
      />
      <router-view />
    </v-content>
  </v-app>
</template>

<script>
import Navigation from "@/components/navigation.vue";

import { EventBus } from "@/event_bus";

export default {
  mounted() {
    EventBus.$on("showTooltip", event => {
      var tooltip = this.$refs.tooltip;
      tooltip.innerHTML = "<span>" + event.text + "</span>";
      tooltip.style.display = "initial";
      tooltip.style.top = event.top + "px";
      tooltip.style.left = event.left - tooltip.scrollWidth / 2 + 15 + "px";
    });
    EventBus.$on("hideTooltip", () => {
      this.$refs.tooltip.style.display = "none";
    });
  },
  components: {
    Navigation
  }
};
</script>

<style lang="scss">
.tooltip {
  z-index: 100;
  position: absolute;
}
</style>
