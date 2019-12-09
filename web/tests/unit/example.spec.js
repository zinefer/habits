import { shallowMount } from "@vue/test-utils";
import HelloWorld from "@/components/HelloWorld.vue";

let wrapper = null;

beforeEach(() => {
  wrapper = shallowMount(HelloWorld, {
    stubs: ["v-flex", "v-img", "v-layout", "v-container"]
  });
});

afterEach(() => {
  wrapper.destroy();
});

describe("HelloWorld.vue", () => {
  it("renders welcome", () => {
    expect(wrapper.text()).toContain("Welcome to Vuetify");
  });
});
