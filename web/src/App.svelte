<script lang="ts">
  import { pb } from './lib/pb';
  import { fetchStatus } from './lib/status';

  const status = fetchStatus(pb);
</script>

<main class="mx-auto flex min-h-dvh max-w-md flex-col justify-center gap-6 p-6">
  <h1 class="text-3xl font-semibold tracking-tight">plasto</h1>
  <p class="text-zinc-600 dark:text-zinc-400">
    Your app starts here. Edit <code class="font-mono text-sm">web/src/App.svelte</code> and save.
  </p>

  <div class="rounded-xl border border-zinc-200 p-4 dark:border-zinc-800">
    {#await status}
      <p class="text-zinc-500">Connecting to the backend…</p>
    {:then s}
      {#if s.ok}
        <p><span class="text-emerald-600 dark:text-emerald-400">●</span> Backend {s.version}</p>
      {:else}
        <p><span class="text-red-600 dark:text-red-400">●</span> Backend unreachable: {s.error}</p>
      {/if}
    {/await}
  </div>
</main>
