import type PocketBase from 'pocketbase';
import { describe, expect, it } from 'vitest';
import { fetchStatus } from './status';

const client = (send: () => Promise<unknown>) => ({ send }) as unknown as PocketBase;

describe('fetchStatus', () => {
  it('returns the backend version', async () => {
    const status = await fetchStatus(client(async () => ({ version: '1.2.3' })));
    expect(status).toEqual({ ok: true, version: '1.2.3' });
  });

  it('reports an unreachable backend instead of throwing', async () => {
    const status = await fetchStatus(
      client(async () => {
        throw new Error('Failed to fetch');
      }),
    );
    expect(status).toEqual({ ok: false, error: 'Failed to fetch' });
  });
});
