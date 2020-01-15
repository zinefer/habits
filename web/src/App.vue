<template>
  <v-app>
    <Navigation />
    <v-content>
      <div ref="tooltip" class="tooltip" style="display:none" />
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
  background: rgba(97, 97, 97, 0.9);
  color: #ffffff;
  border-radius: 4px;
  font-size: 14px;
  line-height: 22px;
  display: inline-block;
  padding: 5px 16px;
  position: absolute;
  text-transform: initial;
  width: auto;
  opacity: 1;
  pointer-events: none;
}
</style>
