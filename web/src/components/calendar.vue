<template>
  <div>
    <svg class="calendar-wrapper" :height="8 * squareSize">
      <!-- <g class="cal-months" v-for="(month, i) in ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']" :key="month">
        <text
          fill="#000000"
          :y="1.25 * squareSize"
          :x="(i * 2 * (squareSize + 1)) + (squareSize + 1)">{{ month }}</text>
      </g> -->
      <g class="calendar-week" v-for="(week, i) in values" :key="i">
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
      </g>
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
//import axios from "axios";

const DEFAULT_RANGE_COLOR = [
  "#ebedf0",
  "#c6e48b",
  "#7bc96f",
  "#239a3b",
  "#196127"
];

export default {
  name: "calendar",
  props: {
    rangeColors: {
      type: Array,
      default: () => DEFAULT_RANGE_COLOR
    }
  },
  data() {
    return {
      values: [],
      squareSize: 0
    };
  },
  methods: {
    contribCount(count) {
      return count >= this.rangeColors.length
        ? this.rangeColors.length - 1
        : count;
    }
  },
  mounted() {
    var data = {
      data: {
        user: {
          contributionsCollection: {
            contributionCalendar: {
              totalContributions: 579,
              weeks: [
                {
                  contributionDays: [
                    {
                      date: "2018-12-30T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2018-12-31T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-01T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-02T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-03T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-04T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-05T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-01-06T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-07T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-08T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-09T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-10T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-11T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-12T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-01-13T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-14T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-15T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-16T00:00:00.000+00:00",
                      contributionCount: 3
                    },
                    {
                      date: "2019-01-17T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-18T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-19T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-01-20T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-21T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-22T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-01-23T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-01-24T00:00:00.000+00:00",
                      contributionCount: 5
                    },
                    {
                      date: "2019-01-25T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-01-26T00:00:00.000+00:00",
                      contributionCount: 1
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-01-27T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-01-28T00:00:00.000+00:00",
                      contributionCount: 7
                    },
                    {
                      date: "2019-01-29T00:00:00.000+00:00",
                      contributionCount: 11
                    },
                    {
                      date: "2019-01-30T00:00:00.000+00:00",
                      contributionCount: 6
                    },
                    {
                      date: "2019-01-31T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-02-01T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-02-02T00:00:00.000+00:00",
                      contributionCount: 1
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-02-03T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-02-04T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-02-05T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-02-06T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-02-07T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-02-08T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-02-09T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-02-10T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-02-11T00:00:00.000+00:00",
                      contributionCount: 9
                    },
                    {
                      date: "2019-02-12T00:00:00.000+00:00",
                      contributionCount: 10
                    },
                    {
                      date: "2019-02-13T00:00:00.000+00:00",
                      contributionCount: 5
                    },
                    {
                      date: "2019-02-14T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-02-15T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-02-16T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-02-17T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-02-18T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-02-19T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-02-20T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-02-21T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-02-22T00:00:00.000+00:00",
                      contributionCount: 3
                    },
                    {
                      date: "2019-02-23T00:00:00.000+00:00",
                      contributionCount: 1
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-02-24T00:00:00.000+00:00",
                      contributionCount: 8
                    },
                    {
                      date: "2019-02-25T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-02-26T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-02-27T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-02-28T00:00:00.000+00:00",
                      contributionCount: 7
                    },
                    {
                      date: "2019-03-01T00:00:00.000+00:00",
                      contributionCount: 11
                    },
                    {
                      date: "2019-03-02T00:00:00.000+00:00",
                      contributionCount: 2
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-03-03T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-03-04T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-03-05T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-03-06T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-03-07T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-03-08T00:00:00.000+00:00",
                      contributionCount: 3
                    },
                    {
                      date: "2019-03-09T00:00:00.000+00:00",
                      contributionCount: 4
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-03-10T00:00:00.000+00:00",
                      contributionCount: 3
                    },
                    {
                      date: "2019-03-11T00:00:00.000+00:00",
                      contributionCount: 9
                    },
                    {
                      date: "2019-03-12T00:00:00.000+00:00",
                      contributionCount: 5
                    },
                    {
                      date: "2019-03-13T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-03-14T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-03-15T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-03-16T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-03-17T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-03-18T00:00:00.000+00:00",
                      contributionCount: 6
                    },
                    {
                      date: "2019-03-19T00:00:00.000+00:00",
                      contributionCount: 4
                    },
                    {
                      date: "2019-03-20T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-03-21T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-03-22T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-03-23T00:00:00.000+00:00",
                      contributionCount: 1
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-03-24T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-03-25T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-03-26T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-03-27T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-03-28T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-03-29T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-03-30T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-03-31T00:00:00.000+00:00",
                      contributionCount: 3
                    },
                    {
                      date: "2019-04-01T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-04-02T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-04-03T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-04-04T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-04-05T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-04-06T00:00:00.000+00:00",
                      contributionCount: 4
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-04-07T00:00:00.000+00:00",
                      contributionCount: 5
                    },
                    {
                      date: "2019-04-08T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-04-09T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-04-10T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-04-11T00:00:00.000+00:00",
                      contributionCount: 7
                    },
                    {
                      date: "2019-04-12T00:00:00.000+00:00",
                      contributionCount: 12
                    },
                    {
                      date: "2019-04-13T00:00:00.000+00:00",
                      contributionCount: 1
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-04-14T00:00:00.000+00:00",
                      contributionCount: 9
                    },
                    {
                      date: "2019-04-15T00:00:00.000+00:00",
                      contributionCount: 10
                    },
                    {
                      date: "2019-04-16T00:00:00.000+00:00",
                      contributionCount: 4
                    },
                    {
                      date: "2019-04-17T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-04-18T00:00:00.000+00:00",
                      contributionCount: 7
                    },
                    {
                      date: "2019-04-19T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-04-20T00:00:00.000+00:00",
                      contributionCount: 1
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-04-21T00:00:00.000+00:00",
                      contributionCount: 6
                    },
                    {
                      date: "2019-04-22T00:00:00.000+00:00",
                      contributionCount: 9
                    },
                    {
                      date: "2019-04-23T00:00:00.000+00:00",
                      contributionCount: 4
                    },
                    {
                      date: "2019-04-24T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-04-25T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-04-26T00:00:00.000+00:00",
                      contributionCount: 23
                    },
                    {
                      date: "2019-04-27T00:00:00.000+00:00",
                      contributionCount: 5
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-04-28T00:00:00.000+00:00",
                      contributionCount: 12
                    },
                    {
                      date: "2019-04-29T00:00:00.000+00:00",
                      contributionCount: 5
                    },
                    {
                      date: "2019-04-30T00:00:00.000+00:00",
                      contributionCount: 12
                    },
                    {
                      date: "2019-05-01T00:00:00.000+00:00",
                      contributionCount: 15
                    },
                    {
                      date: "2019-05-02T00:00:00.000+00:00",
                      contributionCount: 7
                    },
                    {
                      date: "2019-05-03T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-05-04T00:00:00.000+00:00",
                      contributionCount: 1
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-05-05T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-05-06T00:00:00.000+00:00",
                      contributionCount: 4
                    },
                    {
                      date: "2019-05-07T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-05-08T00:00:00.000+00:00",
                      contributionCount: 5
                    },
                    {
                      date: "2019-05-09T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-05-10T00:00:00.000+00:00",
                      contributionCount: 7
                    },
                    {
                      date: "2019-05-11T00:00:00.000+00:00",
                      contributionCount: 8
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-05-12T00:00:00.000+00:00",
                      contributionCount: 3
                    },
                    {
                      date: "2019-05-13T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-05-14T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-05-15T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-05-16T00:00:00.000+00:00",
                      contributionCount: 5
                    },
                    {
                      date: "2019-05-17T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-05-18T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-05-19T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-05-20T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-05-21T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-05-22T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-05-23T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-05-24T00:00:00.000+00:00",
                      contributionCount: 8
                    },
                    {
                      date: "2019-05-25T00:00:00.000+00:00",
                      contributionCount: 17
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-05-26T00:00:00.000+00:00",
                      contributionCount: 3
                    },
                    {
                      date: "2019-05-27T00:00:00.000+00:00",
                      contributionCount: 4
                    },
                    {
                      date: "2019-05-28T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-05-29T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-05-30T00:00:00.000+00:00",
                      contributionCount: 5
                    },
                    {
                      date: "2019-05-31T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-06-01T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-06-02T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-03T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-06-04T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-05T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-06T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-07T00:00:00.000+00:00",
                      contributionCount: 3
                    },
                    {
                      date: "2019-06-08T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-06-09T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-10T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-06-11T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-06-12T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-13T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-14T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-15T00:00:00.000+00:00",
                      contributionCount: 1
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-06-16T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-06-17T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-06-18T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-19T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-20T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-21T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-22T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-06-23T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-24T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-06-25T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-26T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-27T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-06-28T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-06-29T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-06-30T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-01T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-02T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-03T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-04T00:00:00.000+00:00",
                      contributionCount: 4
                    },
                    {
                      date: "2019-07-05T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-06T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-07-07T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-08T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-09T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-10T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-11T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-07-12T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-13T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-07-14T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-15T00:00:00.000+00:00",
                      contributionCount: 3
                    },
                    {
                      date: "2019-07-16T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-17T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-18T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-19T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-20T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-07-21T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-22T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-23T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-24T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-25T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-26T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-27T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-07-28T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-29T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-30T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-07-31T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-01T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-02T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-03T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-08-04T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-05T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-06T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-07T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-08T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-09T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-10T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-08-11T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-12T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-13T00:00:00.000+00:00",
                      contributionCount: 3
                    },
                    {
                      date: "2019-08-14T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-15T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-16T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-17T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-08-18T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-19T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-20T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-21T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-22T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-23T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-24T00:00:00.000+00:00",
                      contributionCount: 1
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-08-25T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-26T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-27T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-28T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-29T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-30T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-08-31T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-09-01T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-02T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-03T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-09-04T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-05T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-06T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-07T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-09-08T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-09T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-10T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-11T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-12T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-13T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-14T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-09-15T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-16T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-17T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-18T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-19T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-20T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-21T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-09-22T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-23T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-24T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-25T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-26T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-27T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-28T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-09-29T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-09-30T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-10-01T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-02T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-03T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-04T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-05T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-10-06T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-07T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-08T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-09T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-10T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-11T00:00:00.000+00:00",
                      contributionCount: 3
                    },
                    {
                      date: "2019-10-12T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-10-13T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-14T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-15T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-16T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-17T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-18T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-19T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-10-20T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-21T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-22T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-23T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-24T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-25T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-26T00:00:00.000+00:00",
                      contributionCount: 4
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-10-27T00:00:00.000+00:00",
                      contributionCount: 14
                    },
                    {
                      date: "2019-10-28T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-10-29T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-30T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-10-31T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-01T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-02T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-11-03T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-04T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-05T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-06T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-07T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-08T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-09T00:00:00.000+00:00",
                      contributionCount: 0
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-11-10T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-11T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-12T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-13T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-14T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-11-15T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-16T00:00:00.000+00:00",
                      contributionCount: 4
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-11-17T00:00:00.000+00:00",
                      contributionCount: 3
                    },
                    {
                      date: "2019-11-18T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-19T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-20T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-21T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-22T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-11-23T00:00:00.000+00:00",
                      contributionCount: 5
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-11-24T00:00:00.000+00:00",
                      contributionCount: 5
                    },
                    {
                      date: "2019-11-25T00:00:00.000+00:00",
                      contributionCount: 3
                    },
                    {
                      date: "2019-11-26T00:00:00.000+00:00",
                      contributionCount: 4
                    },
                    {
                      date: "2019-11-27T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-11-28T00:00:00.000+00:00",
                      contributionCount: 4
                    },
                    {
                      date: "2019-11-29T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-11-30T00:00:00.000+00:00",
                      contributionCount: 1
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-12-01T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-12-02T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-03T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-04T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-05T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-06T00:00:00.000+00:00",
                      contributionCount: 5
                    },
                    {
                      date: "2019-12-07T00:00:00.000+00:00",
                      contributionCount: 5
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-12-08T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-12-09T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-10T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-11T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-12T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-13T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-14T00:00:00.000+00:00",
                      contributionCount: 2
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-12-15T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-16T00:00:00.000+00:00",
                      contributionCount: 2
                    },
                    {
                      date: "2019-12-17T00:00:00.000+00:00",
                      contributionCount: 5
                    },
                    {
                      date: "2019-12-18T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-19T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-12-20T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-12-21T00:00:00.000+00:00",
                      contributionCount: 2
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-12-22T00:00:00.000+00:00",
                      contributionCount: 6
                    },
                    {
                      date: "2019-12-23T00:00:00.000+00:00",
                      contributionCount: 5
                    },
                    {
                      date: "2019-12-24T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-25T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-12-26T00:00:00.000+00:00",
                      contributionCount: 0
                    },
                    {
                      date: "2019-12-27T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-28T00:00:00.000+00:00",
                      contributionCount: 2
                    }
                  ]
                },
                {
                  contributionDays: [
                    {
                      date: "2019-12-29T00:00:00.000+00:00",
                      contributionCount: 8
                    },
                    {
                      date: "2019-12-30T00:00:00.000+00:00",
                      contributionCount: 1
                    },
                    {
                      date: "2019-12-31T00:00:00.000+00:00",
                      contributionCount: 8
                    },
                    {
                      date: "2020-01-01T00:00:00.000+00:00",
                      contributionCount: 9
                    }
                  ]
                }
              ]
            }
          }
        }
      }
    };

    this.values = data.data.user.contributionsCollection.contributionCalendar.weeks.slice(
      53 - this.months * 4
    );
    this.squareSize =
      document.getElementById("github-stats").offsetWidth / this.values.length -
      1;
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
