<template>
  <div ref="calendar">
    <svg class="calendar-wrapper" :height="8 * squareSize - squareSize * 0.8">
      <!-- <g class="cal-months" v-for="(month, i) in ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']" :key="month">
        <text
          fill="#000000"
          :y="1.25 * squareSize"
          :x="(i * 2 * (squareSize + 1)) + (squareSize + 1)">{{ month }}</text>
      </g> -->

      <g class="calendar-week" v-for="(week, i) in 53" :key="i">
        <rect
          class="calendar-day"
          v-for="(day, j) in values.slice(i * 7, i * 7 + 7)"
          :id="day.Day"
          :key="j"
          :style="{ fill: color(day.Count) }"
          :width="squareSize"
          :height="squareSize"
          :x="i * (squareSize + 1)"
          :y="j * (squareSize + 1)"
        />
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
      squareSize: 0
    };
  },
  methods: {
    color(count) {
      if (count == 0) return this.rangeColors[0];
      //var index = Math.round((count + (this.max * 0.25) / this.max) * 4)
      var index = Math.round((count / this.max) * 4);
      if (index < 1) return this.rangeColors[1];
      if (index > 4) return this.rangeColors[4];
      return this.rangeColors[index];
    }
  },
  mounted() {
    ActivitiesApi.get(this.habitID)
      .then(resp => {
        this.values = resp.data;

        var filtered = this.values
          .filter(value => value.Count > 1)
          .map(function(value) {
            return value.Count;
          });

        var sorted = filtered.sort(function(a, b) {
          return a - b;
        });

        var len = sorted.length;
        this.max = sorted[Math.round(len * 0.95)];
      })
      .finally(() => {
        this.loading = false;
      });

    this.squareSize = this.$refs.calendar.offsetWidth / 55;
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
