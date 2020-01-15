import { shallowMount } from "@vue/test-utils";
import Navigation from "@/components/navigation.vue";

describe("Application navigation", () => {
  let options, wrapper;

  beforeEach(() => {
    options = {
      stubs: {
        "v-app-bar": true,
        "v-img": true,
        "v-spacer": true,
        "v-menu": '<div class="v-menu"><slot/><slot name="activator"/></div>',
        "v-list": true,
        "v-list-item": true,
        "v-list-item-title": true,
        "v-icon": true,
        "v-btn": true
      },
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

  it("renders the app name", () => {
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
    expect(wrapper.find(".v-menu").text()).toContain("Sign out");
  });

  it("renders the user name when logged in", () => {
    let random = Math.random()
      .toString(36)
      .substring(7);
    options.mocks.$store.getters.isLoggedIn = true;
    options.mocks.$store.getters.currentUser = random;
    wrapper = shallowMount(Navigation, options);
    expect(wrapper.text()).toContain(random);
  });
});
