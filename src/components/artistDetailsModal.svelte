<script lang="ts">
  // Props
  /** Exposes parent props to this component. */

  // Stores
  import { getModalStore } from "@skeletonlabs/skeleton";

  const modalStore = getModalStore();

  // Base Classes
  const cBase = "card p-4 w-modal shadow-xl space-y-4";
  const cHeader = "text-2xl font-bold";

  const bodyData: bodyObject | null = $modalStore[0].body ? JSON.parse($modalStore[0].body) as unknown as bodyObject : null;

  interface bodyObject {
    artistName: string;
    artistDetails: {
      [key: string]: string;
    };
  }

  console.log(bodyData);
</script>

<!-- @component This example creates a simple form modal. -->

{#if bodyData}
  <div class={cBase}>
    <div class:cHeader>
      <h2 class="h2">{bodyData.artistName}</h2>
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
        {#each Object.keys(bodyData.artistDetails) as tableKey}
          <tr>
            <td>{tableKey}</td>
            <td>
              {bodyData.artistDetails[tableKey]}
            </td>
          </tr>
        {/each}
        </tbody>
      </table>
    </div>
  </div>
{/if}
