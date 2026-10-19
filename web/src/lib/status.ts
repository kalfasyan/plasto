import type PocketBase from 'pocketbase';

export type Status = { ok: true; version: string } | { ok: false; error: string };

/** Asks the backend whether it is up and which version it runs. */
export async function fetchStatus(client: PocketBase): Promise<Status> {
  try {
    const info = await client.send<{ version: string }>('/api/app/info', {});
    return { ok: true, version: info.version };
  } catch (err) {
    return { ok: false, error: err instanceof Error ? err.message : String(err) };
  }
}
