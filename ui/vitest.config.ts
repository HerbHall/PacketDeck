import { defineConfig, mergeConfig } from "vitest/config";
import viteConfig from "./vite.config";
import path from "path";

export default mergeConfig(
  viteConfig,
  defineConfig({
    test: {
      environment: "jsdom",
      globals: true,
      setupFiles: "./src/test-setup.ts",
    },
    resolve: {
      alias: {
        "@docker/extension-api-client": path.resolve(
          __dirname,
          "src/__mocks__/@docker/extension-api-client.ts",
        ),
      },
    },
  }),
);
