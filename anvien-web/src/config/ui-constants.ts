export const DEFAULT_BACKEND_URL = 'http://127.0.0.1:4848';
export const DEFAULT_OLLAMA_BASE_URL = 'http://127.0.0.1:11434';
export const DEFAULT_OPENROUTER_BASE_URL = 'https://openrouter.ai/api/v1';

/** Minimum Node.js version required by the anvien CLI (injected by Vite from package.json engines). */
declare const __REQUIRED_NODE_VERSION__: string;
export const REQUIRED_NODE_VERSION = __REQUIRED_NODE_VERSION__;
