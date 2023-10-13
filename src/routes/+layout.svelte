<script lang="ts">
  import "../app.postcss";
  import { AppBar, AppShell, initializeStores, Modal, storePopup } from "@skeletonlabs/skeleton";
  import { arrow, autoUpdate, computePosition, flip, offset, shift } from "@floating-ui/dom";

  storePopup.set({ computePosition, autoUpdate, offset, shift, flip, arrow });
  initializeStores();
</script>

<Modal />
<!-- App Shell -->
<AppShell>
  <!-- App Bar -->
  <svelte:fragment slot="header">
    <AppBar gridColumns="grid-cols-3" slotDefault="place-self-center" slotTrail="place-content-end">
      <svelte:fragment slot="lead">(icon)</svelte:fragment>
      <a href="/">Home</a>
      |
      <a href="/song_search">Spotify Song Hidden Info Search</a>
      <svelte:fragment slot="trail">(actions)</svelte:fragment>
    </AppBar>
  </svelte:fragment>

  <!-- Page Route Content -->
  <div class="container h-full mx-auto flex justify-center">
    <slot />
  </div>
</AppShell>
