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

      <!--i = 0
      0, 6

      i = 1
      7, 13-->

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

//const arrAvg = arr => arr.reduce((a, b) => a + b, 0) / arr.length;

/*function quickselect_median(arr) {
  const L = arr.length,
    halfL = L / 2;
  if (L % 2 == 1) return quickselect(arr, halfL);
  else return 0.5 * (quickselect(arr, halfL - 1) + quickselect(arr, halfL));
}

function quickselect(arr, k) {
  // Select the kth element in arr
  // arr: List of numerics
  // k: Index
  // return: The kth element (in numerical order) of arr
  if (arr.length == 1) return arr[0];
  else {
    const pivot = arr[0];
    const lows = arr.filter(e => e < pivot);
    const highs = arr.filter(e => e > pivot);
    const pivots = arr.filter(e => e == pivot);
    if (k < lows.length)
      // the pivot is too high
      return quickselect(lows, k);
    else if (k < lows.length + pivots.length)
      // We got lucky and guessed the median
      return pivot;
    // the pivot is too low
    else return quickselect(highs, k - lows.length - pivots.length);
  }
}*/

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
      squareSize: 0
    };
  },
  methods: {
    color(count) {
      if (count == 0) return this.rangeColors[0];
      //var index = Math.round((count + (this.max * 0.25) / this.max) * 4)
      var index = Math.round((count / this.max) * 4 + 0.49);
      if (index < 1) return this.rangeColors[1];
      return this.rangeColors[index];
    },
    contribCount(count) {
      return count >= this.rangeColors.length
        ? this.rangeColors.length - 1
        : count;
    }
  },
  mounted() {
    ActivitiesApi.get(this.habitID)
      .then(resp => {
        this.values = resp.data;
        this.max =
          Math.max.apply(
            Math,
            this.values.map(function(o) {
              return o.Count;
            })
          ) || 1;

        /*var filtered = this.values
          .filter(value => value.Count > 0)
          .map(function(value) {
            return value.Count;
          });

        var sorted = filtered.sort(function(a, b) {
          return a - b;
        });
        var pieces = 4;
        var size = Math.floor(sorted.length / pieces);
        for (let index = 0; index < pieces; index++) {
          let start = index * size;
          let end = index * size + size;
          if (index == pieces - 1) {
            end = sorted.length;
          }

          //let v = sorted.slice(start, end);
          //let median = v[Math.round((end - start) / 2)];
          //let avg = arrAvg(v);
        }*/
      })
      .finally(() => {
        this.loading = false;
      });

    /*this.values = data.data.user.contributionsCollection.contributionCalendar.weeks.slice(
      53 - this.months * 4
    );*/
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
