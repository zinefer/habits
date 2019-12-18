import { shallowMount } from "@vue/test-utils";
import Navigation from "@/components/navigation.vue";

describe("Application navigation", () => {
  let options, wrapper;

  beforeEach(() => {
    options = {
      stubs: [
        "v-app-bar",
        "v-img",
        "v-spacer",
        "v-menu",
        "v-list",
        "v-list-item",
        "v-list-item-title",
        "v-btn"
      ],
      mocks: {
        $store: {
          getters: {
            isLoggedIn: false
          }
        }
      }
    };
  });

  afterEach(() => {
    wrapper.destroy();
  });

  it("renders our name", () => {
    wrapper = shallowMount(Navigation, options);
    expect(wrapper.text()).toContain("Habits");
  });

  it("renders the sign in button when not logged in", () => {
    wrapper = shallowMount(Navigation, options);
    expect(wrapper.find("v-btn-stub").text()).toContain("Sign in");
  });

  it("renders the user dropdown when logged in", () => {
    options.mocks.$store.getters.isLoggedIn = true;
    wrapper = shallowMount(Navigation, options);
    expect(wrapper.find("v-menu-stub").text()).toContain("Sign out");
  });
});
