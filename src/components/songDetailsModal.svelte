<script lang="ts">
  // Stores
  import { getModalStore } from "@skeletonlabs/skeleton";
  import { Spotify } from "sveltekit-embed";

  const modalStore = getModalStore();

  // Base Classes
  const cBase = "card p-4 w-modal shadow-xl space-y-4";

  const bodyData: bodyObject | null = $modalStore[0].body ? JSON.parse($modalStore[0].body) as unknown as bodyObject : null;

  interface bodyObject {
    trackName: string;
    trackId: string;
    trackAudioFeatures: {
      [key: string]: string;
    };
  }
</script>

{#if bodyData}
  <div class={cBase}>
    <div class="-mb-10">
      <Spotify spotifyLink="track/{bodyData.trackId}" height="200px" width="100%" />
    </div>
    <div class="table-container">
      <table class="table table-interactive" role="grid">
        <thead class="table-head ">
        <tr>
          <th>Key</th>
          <th>Value</th>
        </tr>
        </thead>
        <tbody class="table-body ">
        {#each Object.keys(bodyData.trackAudioFeatures) as tableKey}
          <tr>
            <td>{tableKey}</td>
            <td>
              {bodyData.trackAudioFeatures[tableKey]}
            </td>
          </tr>
        {/each}
        </tbody>
      </table>
    </div>
  </div>
{/if}
