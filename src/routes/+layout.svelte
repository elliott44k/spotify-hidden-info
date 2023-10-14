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
      <svelte:fragment slot="lead">
        <div class="place-content-start flex-nowrap">
          <a class="inline-block btn btn-md px-1" href="/">Home</a>
          <a class="inline-block btn btn-md px-1" href="/song_search">Song Search</a>
          <a class="inline-block btn btn-md px-1" href="/playlist_generator">Playlist Generator</a>
          <a class="inline-block btn btn-md px-1" href="/artist_id_search">Artist ID Lookup</a>
        </div>
      </svelte:fragment>
      <h1 class="h3">Spotify Hidden Info Search</h1>
      <svelte:fragment slot="trail">Designed and Coded by: Elliott Kim</svelte:fragment>
    </AppBar>
  </svelte:fragment>

  <!-- Page Route Content -->
  <div class="container h-full mx-auto flex justify-center">
    <slot />
  </div>
</AppShell>
