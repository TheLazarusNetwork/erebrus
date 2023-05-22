export function getBaseUrl(): string {
    const { protocol, host } = window.location;
    return `${protocol}//${host}`;
  }