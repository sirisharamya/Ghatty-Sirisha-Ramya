const { toggleVisibility } = require("../src/domManipulation");

describe("DOM Manipulation", () => {
  test("should toggle visibility correctly", () => {
    const element = { style: { display: "none" } };

    toggleVisibility(element);
    expect(element.style.display).toBe("block");

    toggleVisibility(element);
    expect(element.style.display).toBe("none");
  });
});
