import { defineConfig } from 'vitest/config';

export default defineConfig({
  test: {
    environment: 'jsdom',
    globals: true,
    include: ['cmd/dashboard/static/js/**/*.test.js'],
    coverage: {
      reporter: ['text', 'html'],
      include: ['cmd/dashboard/static/js/**/*.js'],
      exclude: ['**/*.test.js', '**/node_modules/**'],
    },
  },
});
