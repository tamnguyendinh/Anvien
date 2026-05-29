import { beforeEach } from 'vitest';
import '@testing-library/jest-dom/vitest';

// Reset storage between tests
beforeEach(() => {
  sessionStorage.removeItem('anvien-llm-settings');
  localStorage.removeItem('anvien-llm-settings'); // legacy key (migration)
  localStorage.removeItem('anvien.graphLinksVisible');
});
