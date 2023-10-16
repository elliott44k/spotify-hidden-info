<svelte:head>
  <title>Spotify Hidden Info Search</title>
  <meta name="description" content="Find hidden info about your favorite songs on Spotify" />
</svelte:head>

<script lang="ts">
  import type { ModalComponent, ModalSettings, PopupSettings } from "@skeletonlabs/skeleton";
  import { getModalStore, popup } from "@skeletonlabs/skeleton";
  import modalComp from "../../components/songDetailsModal.svelte";
  import { loader } from "../../components/loadingOverlay";
  import { writable } from "svelte/store";

  // loading overlay control
  let loading = writable(false);

  // Input info popups
  const popupFocusBlur1: PopupSettings = {
    event: "focus-blur",
    target: "popupFocusBlur1",
    placement: "right"
  };
  const popupFocusBlur2: PopupSettings = {
    event: "focus-blur",
    target: "popupFocusBlur2",
    placement: "right"
  };

  // handling for song search
  let songs = "";
  let artists = "";
  let error = false;
  let spotifyTracks = {};
  let tableReady = false;

  // spotify api call for track search
  async function getSpotifySearch() {
    if (songs.length === 0) {
      return;
    } else {
      return await fetch(`${window.location.origin}/api/getSpotifySearch`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          "track_name": songs,
          "artist_name": artists,
          "type": "track",
          "market": "US"
        })
      }).then(response => response.json()
      ).then(data => {
        tableReady = true;
        return data.tracks.items;
      }).catch(error => {
      });
    }
  }

  // form submit handler
  function handleSubmit() {
    if (songs.length === 0 && artists.length === 0) {
      //handle error state
      error = true;
    } else {
      error = false;
      spotifyTracks = getSpotifySearch();
    }
  }

  // modal control
  const modalStore = getModalStore();
  const modalComponent: ModalComponent = { ref: modalComp };

  // spotify api call for track audio features
  async function getSpotifyTrackAudioFeatures(trackName: string, trackId: string) {
    return await fetch(`${window.location.origin}/api/getSpotifyTrackAudioFeatures`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        "track_id": trackId
      })
    }).then(response => response.json()
    ).then(data => {
      const trackAudioFeatures = {
        "danceability": Math.round(data.danceability * 10000) / 100,
        "energy": Math.round(data.energy * 10000) / 100,
        "key": data.key,
        "loudness": Math.round(data.loudness * 100) / 100 + " db",
        "mode": data.mode,
        "speechiness": Math.round(data.speechiness * 10000) / 100,
        "acousticness": Math.round(data.acousticness * 10000) / 100,
        "instrumentalness": Math.round(data.instrumentalness * 10000) / 100,
        "liveness": Math.round(data.liveness * 10000) / 100,
        "valence": data.valence,
        "tempo": Math.round(data.tempo) + " bpm",
        "duration": new Date(data.duration_ms).toISOString().slice(11, 19) + (" (hh:mm:ss)"),
        "time_signature": data.time_signature
      };
      const modal: ModalSettings = {
        type: "component",
        component: modalComponent,
        body: JSON.stringify({ trackName, trackId, trackAudioFeatures })
      };
      loading.update(() => false);
      modalStore.trigger(modal);
    });
  }

  // handle table row click
  function handleTableClick(trackName: string, trackId: string) {
    loading.update(() => true);
    getSpotifyTrackAudioFeatures(trackName, trackId);
  }
</script>

<div class="space-y-10 w-11/12 m-6 text-left flex flex-col grow items-center" use:loader={loading}>
  <h1 class="h1">Hidden Song Info Search</h1>

  <form class="w-9/12" on:submit|preventDefault={handleSubmit}>
    <label class="label">
      <span class="">Song Name </span>
      <input class="input mt-0" bind:value={songs} type="text" placeholder={"Summer ('til the end)"}
             oninvalid="this.setCustomValidity('At least one song is required')" use:popup={popupFocusBlur1} />
      <div class="card p-3 variant-filled text-sm mr-10" data-popup="popupFocusBlur1">
        <p>Input multiple songs separated by a comma ie: Hello,Get Low</p>
        <div class="arrow variant-filled" />
      </div>
    </label>
    <label class="label mt-2">
      <span>Artists</span>
      <input class="input mt-0" bind:value={artists} type="text" placeholder={"ESKM"} use:popup={popupFocusBlur2} />
      <div class="card p-3 variant-filled text-sm" data-popup="popupFocusBlur2">
        <p>Input multiple artists separated by a comma ie: ESKM,C.SWAG</p>
        <div class="arrow variant-filled" />
      </div>
    </label>
    {#if error}
      <p class="text-red-500">Please enter a song name</p>
    {/if}
    <button class="btn variant-filled mt-4" type="submit">
      Search
    </button>
  </form>

  {#await spotifyTracks}
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
  {:then tracks}
    {#if tableReady}
      <div class="table-container w-11/12">
        <table class="table table-interactive" role="grid">
          <thead class="table-head ">
          <tr>
            <th class="w-60">Album Art</th>
            <th class="">Track Name</th>
            <th class="">Artist</th>
          </tr>
          </thead>
          <tbody class="table-body">
          {#each Object.values(tracks) as track}
            <tr on:click={() => handleTableClick(track.name, track.id)}>
              <td><img class="object-scale-down" src={track.album.images[0].url}
                       alt={"Album art for: " + track.album.name}></td>
              <td>{track.name}</td>
              <td>
                {#if track.artists.length > 1}
                  {track.artists.shift().name}
                  {#each track.artists as artist}
                    {"| " + artist.name + " "}
                  {/each}
                {:else}
                  {track.artists[0].name}
                {/if}
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