<svelte:head>
  <title>Spotify Hidden Info Search</title>
  <meta name="description" content="Find hidden info about your favorite songs on Spotify" />
</svelte:head>

<script lang="ts">
  import type { ModalComponent, ModalSettings } from "@skeletonlabs/skeleton";
  import { getModalStore } from "@skeletonlabs/skeleton";
  import modalComp from "../../components/artistDetailsModal.svelte";
  import { loader } from "../../components/loadingOverlay";
  import { writable } from "svelte/store";

  // loading overlay control
  let loading = writable(false);

  // handling for artist search
  let artist = "";
  let error = false;
  let spotifyArtists = {};
  let tableReady = false;

  async function getSpotifySearch() {
    if (artist.length === 0) {
      return;
    } else {
      return await fetch(`${window.location.origin}/api/getSpotifySearch`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          "artist_name": artist,
          "type": "artist",
          "market": "US"
        })
      }).then(response => response.json()
      ).then(data => {
        tableReady = true;
        const sortedArtists = data.artists.items;
        sortedArtists.sort(function(a, b) {
          return a.popularity >= b.popularity ? -1 : 1;
        });
        return sortedArtists;
      }).catch(error => {
      });
    }
  }

  function handleSubmit() {
    if (artist.length === 0) {
      //handle error state
      error = true;
    } else {
      error = false;
      spotifyArtists = getSpotifySearch();
    }
  }


  const modalStore = getModalStore();
  const modalComponent: ModalComponent = { ref: modalComp };

  async function getSpotifyArtist(artistId: string) {
    return await fetch(`${window.location.origin}/api/getSpotifyArtist`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        "artist_id": artistId
      })
    }).then(response => response.json()
    ).then(data => {
      const artistDetails = {
        "id": data.id,
        "followers": data.followers.total,
        "genres": data.genres,
        "popularity": data.popularity
      };
      console.log(data);
      console.log(artistDetails);
      const modal: ModalSettings = {
        type: "component",
        component: modalComponent,
        body: JSON.stringify({ artistName: data.name, artistDetails })
      };
      loading.update(() => false);
      modalStore.trigger(modal);
    });
  }

  function handleTableClick(artistId: string) {
    loading.update(() => true);
    getSpotifyArtist(artistId);
  }


</script>

<div class="space-y-10 w-11/12 m-6 text-left flex flex-col grow items-center" use:loader={loading}>
  <h1 class="h1">Hidden Song Info Search</h1>

  <form class="w-9/12" on:submit|preventDefault={handleSubmit}>
    <label class="label mt-2">
      <span>Artist Name</span>
      <input class="input mt-0" bind:value={artist} type="text" placeholder={"ESKM"} />
    </label>
    {#if error}
      <p class="text-red-500">Please enter an artist</p>
    {/if}
    <button class="btn variant-filled mt-4" type="submit">
      Search
    </button>
  </form>

  {#await spotifyArtists}
    <div role="status">
      <svg aria-hidden="true" class="w-8 h-8 mr-2 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600"
           viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path
          d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
          fill="currentColor" />
        <path
          d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
          fill="currentFill" />
      </svg>
      <span class="sr-only">Loading...</span>
    </div>
  {:then artists}
    {#if tableReady}
      <div class="table-container w-11/12">
        <table class="table table-interactive" role="grid">
          <thead class="table-head ">
          <tr>
            <th class="w-60">Artist Photo</th>
            <th class="">Artist Name</th>
            <th class="">Genres</th>
          </tr>
          </thead>
          <tbody class="table-body">
          {#each Object.values(artists) as artist}
            <tr on:click={() => handleTableClick(artist.id)}>
              {#if artist.images.length === 0}
                <td>No image</td>
              {:else}
                <td><img class="object-scale-down" src={artist.images[0].url}
                         alt={"Artist picture for: " + artist.name}></td>
              {/if}
              <td>{artist.name}</td>
              <td class="table-cell-fit">
                <p>
                {#if artist.genres.length > 1}
                  {artist.genres.shift()}
                  {#each artist.genres as genre}
                    {"| " + genre + " "}
                  {/each}
                {:else if artist.genres.length === 1}
                  {artist.genres}
                {/if}
                </p>
              </td>
            </tr>
          {/each}
          </tbody>
        </table>
      </div>
    {/if}
  {:catch error}
    <p style="color: red">{error.message}</p>
  {/await}
</div>