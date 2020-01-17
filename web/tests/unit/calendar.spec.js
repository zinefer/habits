import { shallowMount } from "@vue/test-utils";
import HabitCalendar from "@/components/habit/calendar.vue";

describe("Habit Calendar", () => {
  let options, wrapper;

  beforeEach(() => {
    options = {
      propsData: {
        isMobile: false,
        values: [{ Count: 1, Day: new Date().toISOString() }]
      }
    };
  });

  afterEach(() => {
    wrapper.destroy();
  });

  it("to filter down to the correct number of days on desktop", () => {
    var days = 364 + new Date().getDay() + 1;
    wrapper = shallowMount(HabitCalendar, options);
    expect(wrapper.vm.filteredValues.length).toBe(days);
  });

  it("to filter down to the correct number of days on mobile", () => {
    options.propsData.isMobile = true;
    wrapper = shallowMount(HabitCalendar, options);
    var days = 77 + new Date().getDay() + 1;
    expect(wrapper.vm.filteredValues.length).toBe(days);
  });
});
