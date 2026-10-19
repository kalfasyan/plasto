import PocketBase from 'pocketbase';

// The web app is served by the backend, so the API is on the same origin.
export const pb = new PocketBase('/');
