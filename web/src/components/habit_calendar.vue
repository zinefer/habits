<template>
  <div ref="calendar">
    <svg class="calendar-wrapper" :height="7 * (squareSize + 1) + headerHeight">
      <!-- <g class="cal-months" v-for="(month, i) in ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']" :key="month">
        <text
          fill="#000000"
          :y="1.25 * squareSize"
          :x="(i * 2 * (squareSize + 1)) + (squareSize + 1)">{{ month }}</text>
      </g> -->

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
            :id="day.Day"
            :style="{ fill: color(day.Count) }"
            :width="squareSize"
            :height="squareSize"
            :x="w * squareSize + w + squareSize"
            :y="d * (squareSize + 1) + headerHeight"
          />
          <text
            v-if="isSecondSundayOfMonth(day.Day)"
            text-anchor="middle"
            fill="#ccc"
            :x="w * squareSize + w + squareSize / 2 - 2"
            :y="textHeight"
          >
            {{ getMonthName(day.Day) }}
          </text>
        </g>
      </g>

      <!--<g class="calendar-week" v-for="(week, i) in values" :key="i">
        <rect
          class="calendar-day"
          v-for="(day, j) in week.contributionDays"
          :key="j"
          :style="{ fill: rangeColors[contribCount(day.contributionCount)] }"
          :width="squareSize"
          :height="squareSize"
          :x="i * (squareSize + 1)"
          :y="j * (squareSize + 1)"
        />
      </g>-->
      <!-- <g class="calendar-legend" :y="9 * squareSize">
        <rect v-for="(color, i) in rangeColors" :key="color"
          :style="{fill: color}"
          :width="squareSize"
          :height="squareSize"
          :x="i * (squareSize + 1)"
        /> 
      </g> -->
    </svg>
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
      groups: [],
      values: [],
      squareSize: 0,
      headerHeight: 20,
      textHeight: 16
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
    isSecondSundayOfMonth(date) {
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
#github-stats {
  width: 90%;
}
.calendar-wrapper {
  width: 100%;
  // border: 1px solid black;

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
