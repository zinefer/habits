<template>
  <div ref="calendar">
    <v-scale-transition hide-on-leave>
      <v-skeleton-loader
        v-if="loading"
        class="loader"
        :transition="scale - transition"
        type="heading, avatar@132"
      ></v-skeleton-loader>
      <svg
        class="calendar-wrapper"
        :height="7 * (squareSize + 1) + headerHeight"
        v-else
      >
        <g
          class="calendar-dow"
          v-for="(dayName, dow) in ['Mon', 'Wed', 'Fri']"
          :key="dayName"
        >
          <text
            text-anchor="middle"
            fill="#ccc"
            :x="squareSize / 2 - 2"
            :y="
              2 * dow * (squareSize + 1) +
                headerHeight +
                textHeight * 1.4 +
                squareSize
            "
          >
            {{ dayName }}
          </text>
        </g>
        <g class="calendar-week" v-for="(week, w) in 53" :key="w">
          <g v-for="(day, d) in values.slice(w * 7, w * 7 + 7)" :key="day.Day">
            <rect
              class="calendar-day"
              :day="day.Day"
              :count="day.Count"
              :style="{ fill: color(day.Count) }"
              :width="squareSize"
              :height="squareSize"
              :x="w * squareSize + w + squareSize"
              :y="d * (squareSize + 1) + headerHeight"
              v-on:mouseover="showDayTooltip"
              v-on:mouseleave="hideDayTooltip"
            />
            <text
              v-if="isSecondSundayOfMonth(day.Day)"
              text-anchor="middle"
              fill="#ccc"
              :x="w * squareSize + w + squareSize + squareSize / 2"
              :y="textHeight"
            >
              {{ getMonthName(day.Day) }}
            </text>
          </g>
        </g>
      </svg>
    </v-scale-transition>
  </div>
</template>

<script>
import ActivitiesApi from "@/services/activities";

const DEFAULT_RANGE_COLOR = [
  "#ebedf0",
  "#c6e48b", // 1-3
  "#7bc96f", // 4-7
  "#239a3b", // 8-10
  "#196127" // 11+
];

const months = [
  "Jan",
  "Feb",
  "Mar",
  "Apr",
  "May",
  "Jun",
  "Jul",
  "Aug",
  "Sep",
  "Oct",
  "Nov",
  "Dec"
];

export default {
  name: "HabitCalendar",
  props: {
    habitID: {
      type: Number,
      required: true
    },
    rangeColors: {
      type: Array,
      default: () => DEFAULT_RANGE_COLOR
    }
  },
  data() {
    return {
      max: 1,
      values: [],
      squareSize: 0,
      headerHeight: 20,
      textHeight: 16,
      loading: true
    };
  },
  methods: {
    color(count) {
      if (count == 0) return this.rangeColors[0];
      var index = Math.round((count / this.max) * 4);
      if (index < 1) return this.rangeColors[1];
      if (index > 4) return this.rangeColors[4];
      return this.rangeColors[index];
    },
    showDayTooltip(event) {
      var transform = event.currentTarget.getBoundingClientRect();
      var date = event.currentTarget.getAttribute("day").split("T")[0];

      date = new Date(date)
        .toUTCString()
        .split(" ")
        .slice(0, 4)
        .join(" ");

      this.$emit("showTooltip", {
        top: transform.top - 75,
        left: transform.left,
        text: date + ": " + event.currentTarget.getAttribute("count")
      });
    },
    hideDayTooltip() {
      this.$emit("hideTooltip");
    },
    isSecondSundayOfMonth(date) {
      date = date.split("T")[0];
      date = new Date(date);
      var day = date.getUTCDate();
      date.setDate(7);
      date.setDate(7 + 7 - date.getUTCDay());
      return date.getUTCDate() == day;
    },
    getMonthName(date) {
      return months[date.split("-")[1] - 1];
    }
  },
  mounted() {
    ActivitiesApi.get(this.habitID)
      .then(resp => {
        this.values = resp.data;

        var filtered = this.values
          .filter(value => value.Count > 0)
          .map(function(value) {
            return value.Count;
          });

        var sorted = filtered.sort(function(a, b) {
          return a - b;
        });

        this.max = sorted[Math.round(sorted.length * 0.95) - 1];
      })
      .finally(() => {
        this.loading = false;
      });

    this.squareSize = this.$refs.calendar.offsetWidth / 56;
  }
};
</script>

<style lang="scss">
.loader {
  .v-skeleton-loader__heading {
    margin-bottom: 10px;
  }
  .v-skeleton-loader__avatar {
    display: inline-block;
    margin-right: 10px;
    margin-bottom: 4px;
    border-radius: 4px;
    height: 42px;
    width: 42px;
  }
}

.calendar-wrapper {
  width: 100%;
  margin-bottom: 25px;
  .calendar-legend {
    margin-top: 50%;
  }
  .calendar-day {
    &:hover {
      stroke-width: 1;
      stroke: black;
    }
  }
}
</style>
