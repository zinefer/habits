<template>
  <div ref="calendar">
    <svg
      :class="{ 'calendar-wrapper': true, isMobile: isMobile }"
      :height="height"
    >
      <g
        class="calendar-dow"
        v-for="(dayName, dow) in ['Mon', 'Wed', 'Fri']"
        :key="dayName"
      >
        <text
          text-anchor="middle"
          fill="#ccc"
          :x="calendarDowX(dow)"
          :y="calendarDowY(dow)"
        >
          {{ dayName }}
        </text>
      </g>
      <g class="calendar-week" v-for="(week, w) in displayedWeeks" :key="w">
        <g
          v-for="(day, d) in filteredValues.slice(w * 7, w * 7 + 7)"
          :key="day.Day"
        >
          <rect
            class="calendar-day"
            :day="day.Day"
            :count="day.Count"
            :style="{ fill: color(day.Count) }"
            :width="squareSize"
            :height="squareSize"
            :x="calendarDayX(w, d)"
            :y="calendarDayY(w, d)"
            v-on:mouseover="showDayTooltip"
            v-on:mouseleave="hideDayTooltip"
          />
          <text
            v-if="isSecondSundayOfMonth(day.Day)"
            text-anchor="middle"
            fill="#ccc"
            :x="calendarMonthX(w)"
            :y="calendarMonthY(w)"
          >
            {{ getMonthName(day.Day) }}
          </text>
        </g>
      </g>
    </svg>
  </div>
</template>

<script>
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

Date.prototype.addDays = function(days) {
  var dat = new Date(this.valueOf());
  dat.setDate(dat.getDate() + days);
  return dat;
};

export default {
  name: "HabitCalendar",
  props: {
    values: {
      type: Array,
      required: true
    },
    rangeColors: {
      type: Array,
      default: () => DEFAULT_RANGE_COLOR
    },
    isMobile: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
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
    showDayTooltip(event) {
      var transform = event.currentTarget.getBoundingClientRect();
      var date = event.currentTarget.getAttribute("day").split("T")[0];

      date = new Date(date)
        .toUTCString()
        .split(" ")
        .slice(0, 4)
        .join(" ");

      this.$emit("showTooltip", {
        top:
          transform.top +
          window.scrollY +
          (transform.bottom - transform.top) +
          1,
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
    },
    calendarDayX(w, d) {
      var n = w;
      if (this.isMobile) {
        n = d;
      }

      return n * this.squareSize + n + this.squareSize;
    },
    calendarDayY(w, d) {
      var n = d;
      if (this.isMobile) {
        n = w;
      }

      return n * (this.squareSize + 1) + this.headerHeight;
    },
    calendarDowX(dow) {
      if (this.isMobile) {
        return (
          2 * dow * (this.squareSize + 1) +
          this.headerHeight +
          this.squareSize * 2 -
          2
        );
      }
      return this.squareSize / 2 - 2;
    },
    calendarDowY(dow) {
      if (this.isMobile) {
        return this.squareSize / 2 - 2;
      }
      return (
        2 * dow * (this.squareSize + 1) +
        this.headerHeight +
        this.textHeight * 1.4 +
        this.squareSize
      );
    },
    calendarMonthX(w) {
      if (this.isMobile) {
        return this.textHeight;
      }
      return w * this.squareSize + w + this.squareSize + this.squareSize;
    },
    calendarMonthY(w) {
      if (this.isMobile) {
        var y =
          w * this.squareSize +
          w + // the gap between each row
          this.squareSize +
          this.squareSize -
          this.textHeight / 2;

        // Prevent a month getting displayed on the final week from being
        // pushed off the image due to trying to span the gap
        if (w == this.displayedWeeks - 1) {
          y -= this.textHeight / 2 + w - 2;
        }

        return y;
      }
      return this.textHeight;
    }
  },
  mounted() {
    if (this.isMobile) {
      this.squareSize = this.$refs.calendar.offsetWidth / 9;
    } else {
      this.squareSize = this.$refs.calendar.offsetWidth / 56;
    }
  },
  computed: {
    height: function() {
      if (this.isMobile)
        return this.displayedWeeks * (this.squareSize + 1) + this.headerHeight;
      return 7 * (this.squareSize + 1) + this.headerHeight;
    },
    max: function() {
      var filtered = this.filteredValues
        .filter(value => value.Count > 0)
        .map(function(value) {
          return value.Count;
        });

      var sorted = filtered.sort(function(a, b) {
        return a - b;
      });

      // 95th percentile max
      return sorted[Math.round(sorted.length * 0.95) - 1];
    },
    displayedWeeks: function() {
      if (this.isMobile) return 12;
      return 53;
    },
    filteredValues: function() {
      var data = [],
        today = new Date(),
        yearAgoTs =
          Math.round(today.getTime() / 1000) -
          (364 + today.getDay()) * 24 * 60 * 60,
        yearAgo = new Date(yearAgoTs * 1000),
        vIndex = 0,
        startDate = yearAgo,
        stopDate = today;

      startDate.setHours(0);
      stopDate.setHours(23, 59, 59);

      var date = startDate;
      while (date <= stopDate) {
        if (
          this.values[vIndex] &&
          new Date(this.values[vIndex].Day) < startDate
        ) {
          vIndex++;
          continue;
        }

        var datestr = date
          .toLocaleString("sv", { timeZoneName: "short" })
          .slice(0, 10);

        if (
          this.values[vIndex] &&
          datestr == this.values[vIndex].Day.slice(0, 10)
        ) {
          data.push(this.values[vIndex]);
          vIndex++;
        } else {
          data.push({
            Count: 0,
            Day: datestr
          });
        }

        date = date.addDays(1);
      }

      if (this.isMobile) {
        return data.slice(
          data.length - 1 - this.displayedWeeks * 7 + (7 - today.getDay())
        );
      }

      return data;
    }
  }
};
</script>

<style lang="scss">
.loader {
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
