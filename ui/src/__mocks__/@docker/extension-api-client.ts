import { vi } from "vitest";

const ddClient = {
  extension: {
    vm: {
      service: {
        get: vi.fn(),
        post: vi.fn(),
      },
    },
  },
  docker: {
    cli: {
      exec: vi.fn(),
    },
  },
  desktopUI: {
    toast: {
      success: vi.fn(),
      warning: vi.fn(),
      error: vi.fn(),
    },
  },
};

export function createDockerDesktopClient() {
  return ddClient;
}

export { ddClient };
