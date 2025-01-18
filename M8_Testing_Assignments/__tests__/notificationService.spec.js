const { sendNotification } = require("../src/notificationService");

describe("Notification Service", () => {
  const mockNotificationService = {
    send: jest.fn(),
  };

  test("should return 'Notification Sent' for successful notification", () => {
    mockNotificationService.send.mockReturnValue(true);
    expect(sendNotification(mockNotificationService, "Test message")).toBe("Notification Sent");
  });

  test("should return 'Failed to Send' for failed notification", () => {
    mockNotificationService.send.mockReturnValue(false);
    expect(sendNotification(mockNotificationService, "Test message")).toBe("Failed to Send");
  });
});
