const { delayedGreeting } = require("../src/delayedGreeting");

jest.useFakeTimers();

describe("Delayed Greeting", () => {
  test("should resolve with correct greeting message after delay", async () => {
    const promise = delayedGreeting("John", 1000);

    jest.runAllTimers();

    await expect(promise).resolves.toBe("Hello, John!");
  });
});
