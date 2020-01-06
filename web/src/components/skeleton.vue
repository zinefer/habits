<script>
import { VSkeletonLoader } from "vuetify/lib";

export default {
  functional: true,
  render(h, { data, props, children }) {
    return h(
      VSkeletonLoader,
      {
        ...data,
        props: {
          types: {
            "habit-calendar-days": "avatar@132"
          },
          ...props
        }
      },
      children
    );
  }
};
</script>

<style lang="scss">
.v-skeleton-loader__habit-calendar-days {
  .v-skeleton-loader__avatar {
    display: inline-block;
    margin-right: 10px;
    margin-bottom: 4px;
    border-radius: 4px;
    height: 42px;
    width: 42px;
  }
}
</style>
