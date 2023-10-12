<script lang="ts">
  // Props
  /** Exposes parent props to this component. */
  export let parent: any;

  // Stores
  import { getModalStore } from '@skeletonlabs/skeleton';
  const modalStore = getModalStore();

  // Form Data
  const formData = {
    name: 'Jane Doe',
    tel: '214-555-1234',
    email: 'jdoe@email.com'
  };

  // Base Classes
  const cBase = 'card p-4 w-modal shadow-xl space-y-4';
  const cHeader = 'text-2xl font-bold';
  const cForm = 'border border-surface-500 p-4 space-y-4 rounded-container-token';
</script>

<!-- @component This example creates a simple form modal. -->

{#if $modalStore[0]}
  <div class="modal-example-form {cBase}">
    <h1 class="{cHeader}">{$modalStore[0].body.trackName}</h1>
    <div class="table-container">
      <table class="table table-interactive" role="grid">
        <thead class="table-head ">
          <tr>
            <th>Key</th>
            <th>Value</th>
          </tr>
        </thead>
        <tbody class="table-body ">
        {#each Object.keys($modalStore[0].body.trackAudioFeatures) as tableKey}
          <tr>
            <td>{tableKey}</td>
            <td>
              {$modalStore[0].body.trackAudioFeatures[tableKey]}
            </td>
          </tr>
        {/each}
        </tbody>
      </table>
    </div>
  </div>
{/if}
