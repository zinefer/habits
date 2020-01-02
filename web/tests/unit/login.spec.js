import { shallowMount } from "@vue/test-utils";
import Login from "@/components/dialogs/login.vue";

describe("Login dialog", () => {
  let wrapper;

  beforeEach(() => {
    wrapper = shallowMount(Login, {
      stubs: [
        "v-btn",
        "v-icon",
        "v-card",
        "v-card-title",
        "v-dialog",
        "v-card-text"
      ]
    });
  });

  afterEach(() => {
    wrapper.destroy();
  });

  it("renders a github button", () => {
    expect(
      wrapper
        .findAll("v-btn-stub")
        .at(0)
        .text()
    ).toContain("Continue with GitHub");
  });

  it("renders a google button", () => {
    expect(
      wrapper
        .findAll("v-btn-stub")
        .at(1)
        .text()
    ).toContain("Continue with Google");
  });

  it("renders a facebook button", () => {
    expect(
      wrapper
        .findAll("v-btn-stub")
        .at(2)
        .text()
    ).toContain("Continue with Facebook");
  });
});
