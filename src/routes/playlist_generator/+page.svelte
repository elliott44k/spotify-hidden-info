<svelte:head>
  <title>Spotify Hidden Info Search</title>
  <meta name="description" content="Find hidden info about your favorite songs on Spotify" />
</svelte:head>

<script lang="ts">
  import type { ModalComponent } from "@skeletonlabs/skeleton";
  import { Accordion, AccordionItem, getModalStore } from "@skeletonlabs/skeleton";
  import modalComp from "../../components/songDetailsModal.svelte";
  import { loader } from "../../components/loadingOverlay";
  import { writable } from "svelte/store";

  // loading overlay control
  let loading = writable(false);

  // handling for song search
  let artists = "";
  let genres = "";
  let tracks = "";
  let error = false;
  let errorMessage = "";
  let spotifyTracks = {};
  let tableReady = false;

  async function getSpotifyRecommendations() {
    return await fetch(`${window.location.origin}/api/getSpotifyRecommendations`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        "market": "US",
        "seed_artists": artists ? artists.replace(" ", "") : "",
        "seed_tracks": tracks ? tracks.replace(" ", "") : "",
        "seed_genres": genres ? genres : "",
        "limit": limit.toString(),
        "min_key": minKey ? minKey.toString() : "",
        "max_key": maxKey ? maxKey.toString() : "",
        "target_key": targetKey ? targetKey : "",
        "min_acousticness": minAcousticness ? (minAcousticness / 100).toString() : "",
        "max_acousticness": maxAcousticness ? (maxAcousticness / 100).toString() : "",
        "target_acousticness": targetAcousticness ? (targetAcousticness / 100).toString() : "",
        "min_instrumentalness": minInstrumentalness ? (minInstrumentalness / 100).toString() : "",
        "max_instrumentalness": maxInstrumentalness ? (maxInstrumentalness / 100).toString() : "",
        "target_instrumentalness": targetInstrumentalness ? (targetInstrumentalness / 100).toString() : "",
        "target_mode": targetMode ? targetMode.toString() : "",
        "min_tempo": minTempo ? minTempo.toString() : "",
        "max_tempo": maxTempo ? maxTempo.toString() : "",
        "target_tempo": targetTempo ? targetTempo.toString() : "",
        "min_time_sig": minTimeSig ? minTimeSig.toString() : "",
        "max_time_sig": maxTimeSig ? maxTimeSig.toString() : "",
        "target_time_sig": targetTimeSig ? targetTimeSig.toString() : "",
        "min_duration": minDuration ? (minDuration * 1000).toString() : "",
        "max_duration": maxDuration ? (maxDuration * 1000).toString() : "",
        "target_duration": targetDuration ? (targetDuration * 1000).toString() : "",
        "min_energy": minEnergy ? (minEnergy / 100).toString() : "",
        "max_energy": maxEnergy ? (maxEnergy / 100).toString() : "",
        "target_energy": targetEnergy ? (targetEnergy / 100).toString() : "",
        "min_danceability": minDanceability ? (minDanceability / 100).toString() : "",
        "max_danceability": maxDanceability ? (maxDanceability / 100).toString() : "",
        "target_danceability": targetDanceability ? (targetDanceability / 100).toString() : "",
        "min_liveness": minLiveness ? (minLiveness / 100).toString() : "",
        "max_liveness": maxLiveness ? (maxLiveness / 100).toString() : "",
        "target_liveness": targetLiveness ? (targetLiveness / 100).toString() : "",
        "min_speechiness": minSpeechiness ? (minSpeechiness / 100).toString() : "",
        "max_speechiness": maxSpeechiness ? (maxSpeechiness / 100).toString() : "",
        "target_speechiness": targetSpeechiness ? (targetSpeechiness / 100).toString() : "",
        "min_popularity": minPopularity ? minPopularity.toString() : "",
        "max_popularity": maxPopularity ? maxPopularity.toString() : "",
        "target_popularity": targetPopularity ? targetPopularity.toString() : "",
        "min_loudness": minLoudness ? minLoudness.toString() : "",
        "max_loudness": maxLoudness ? maxLoudness.toString() : "",
        "target_loudness": targetLoudness ? targetLoudness.toString() : ""
      })
    }).then(response => response.json()
    ).then(data => {
      tableReady = true;
      console.log(data);
      return data.tracks;
    }).catch(error => {
      console.log(error);
    });
  }

  function handleSubmit() {
    //data validation
    if (tracks.length === 0 && genres.length === 0 && artists.length === 0) {
      error = true;
      errorMessage = "Please enter at least 1 artist, genre, or track";
    } else if (tracks.split(",").length + genres.split(",").length + artists.split(",").length > 5) {
      error = true;
      errorMessage = "Please enter no more than 5 total artists, genres, and tracks";
    } else if (minAcousticness && maxAcousticness && minAcousticness < maxAcousticness) {
      error = true;
      errorMessage = "Min Acousticness must be less than Max Acousticness";
    } else if (targetAcousticness && minAcousticness && targetAcousticness < minAcousticness) {
      error = true;
      errorMessage = "Target Acousticness must be greater than or equal to Min Acousticness";
    } else if (targetAcousticness && maxAcousticness !== 0 && targetAcousticness > maxAcousticness) {
      error = true;
      errorMessage = "Target Acousticness must be less than or equal to Max Acousticness";
    } else if (minInstrumentalness && maxInstrumentalness && minInstrumentalness < maxInstrumentalness) {
      error = true;
      errorMessage = "Min Instrumentalness must be less than Max Instrumentalness";
    } else if (targetInstrumentalness && minInstrumentalness && targetInstrumentalness < minInstrumentalness) {
      error = true;
      errorMessage = "Target Instrumentalness must be greater than or equal to Min Instrumentalness";
    } else if (targetInstrumentalness && maxInstrumentalness && targetInstrumentalness > maxInstrumentalness) {
      error = true;
      errorMessage = "Target Instrumentalness must be less than or equal to Max Instrumentalness";
    } else if (targetMode < 0 || targetMode > 2) {
      error = true;
      errorMessage = "Please select valid Target Mode";
    } else if (minTempo && maxTempo && minTempo < maxTempo) {
      error = true;
      errorMessage = "Min Tempo must be less than Max Tempo";
    } else if (targetTempo && minTempo && targetTempo < minTempo) {
      error = true;
      errorMessage = "Target Tempo must be greater than or equal to Min Tempo";
    } else if (targetTempo && maxTempo && targetTempo > maxTempo) {
      error = true;
      errorMessage = "Target Tempo must be less than or equal to Max Tempo";
    } else if (minTimeSig && maxTimeSig && minTimeSig < maxTimeSig) {
      error = true;
      errorMessage = "Min Time Signature must be less than Max Time Signature";
    } else if (targetTimeSig && minTimeSig && targetTimeSig < minTimeSig) {
      error = true;
      errorMessage = "Target Time Signature must be greater than or equal to Min Time Signature";
    } else if (targetTimeSig && maxTimeSig && targetTimeSig > maxTimeSig) {
      error = true;
      errorMessage = "Target Time Signature must be less than or equal to Max Time Signature";
    } else if (minDuration && maxDuration && minDuration < maxDuration) {
      error = true;
      errorMessage = "Min Duration must be less than Max Duration";
    } else if (targetDuration && minDuration && targetDuration < minDuration) {
      error = true;
      errorMessage = "Target Duration must be greater than or equal to Min Duration";
    } else if (targetDuration && maxDuration && targetDuration > maxDuration) {
      error = true;
      errorMessage = "Target Duration must be less than or equal to Max Duration";
    } else if (minEnergy && maxEnergy && minEnergy < maxEnergy) {
      error = true;
      errorMessage = "Min Energy must be less than Max Energy";
    } else if (targetEnergy && minEnergy && targetEnergy < minEnergy) {
      error = true;
      errorMessage = "Target Energy must be greater than or equal to Min Energy";
    } else if (targetEnergy && maxEnergy && targetEnergy > maxEnergy) {
      error = true;
      errorMessage = "Target Energy must be less than or equal to Max Energy";
    } else if (minDanceability && maxDanceability && minDanceability < maxDanceability) {
      error = true;
      errorMessage = "Min Danceability must be less than Max Danceability";
    } else if (targetDanceability && minDanceability && targetDanceability < minDanceability) {
      error = true;
      errorMessage = "Target Danceability must be greater than or equal to Min Danceability";
    } else if (targetDanceability && maxDanceability && targetDanceability > maxDanceability) {
      error = true;
      errorMessage = "Target Danceability must be less than or equal to Max Danceability";
    } else if (minLiveness && maxLiveness && minLiveness < maxLiveness) {
      error = true;
      errorMessage = "Min Liveness must be less than Max Liveness";
    } else if (targetLiveness && minLiveness && targetLiveness < minLiveness) {
      error = true;
      errorMessage = "Target Liveness must be greater than or equal to Min Liveness";
    } else if (targetLiveness && maxLiveness && targetLiveness > maxLiveness) {
      error = true;
      errorMessage = "Target Liveness must be less than or equal to Max Liveness";
    } else if (minSpeechiness && maxSpeechiness && minSpeechiness < maxSpeechiness) {
      error = true;
      errorMessage = "Min Speechiness must be less than Max Speechiness";
    } else if (targetSpeechiness && minSpeechiness && targetSpeechiness < minSpeechiness) {
      error = true;
      errorMessage = "Target Speechiness must be greater than or equal to Min Speechiness";
    } else if (targetSpeechiness && maxSpeechiness && targetSpeechiness > maxSpeechiness) {
      error = true;
      errorMessage = "Target Speechiness must be less than or equal to Max Speechiness";
    } else if (minPopularity && maxPopularity && minPopularity < maxPopularity) {
      error = true;
      errorMessage = "Min Popularity must be less than Max Popularity";
    } else if (targetPopularity && minPopularity && targetPopularity < minPopularity) {
      error = true;
      errorMessage = "Target Popularity must be greater than or equal to Min Popularity";
    } else if (targetPopularity && maxPopularity && targetPopularity > maxPopularity) {
      error = true;
      errorMessage = "Target Popularity must be less than or equal to Max Popularity";
    } else if (minLoudness && maxLoudness && minLoudness < maxLoudness) {
      error = true;
      errorMessage = "Min Loudness must be less than Max Loudness";
    } else if (targetLoudness && minLoudness && targetLoudness < minLoudness) {
      error = true;
      errorMessage = "Target Loudness must be greater than or equal to Min Loudness";
    } else if (targetLoudness && maxLoudness && targetLoudness > maxLoudness) {
      error = true;
      errorMessage = "Target Loudness must be less than or equal to Max Loudness";
    } else {
      error = false;
      errorMessage = "";
      loading.set(true);
      getSpotifyRecommendations().then(data => {
        spotifyTracks = data;
        loading.set(false);
      });
    }
  }

  let minKey: number, maxKey: number, targetKey: number;
  let limit: number = 20;
  let minAcousticness: number, maxAcousticness: number, targetAcousticness: number;
  let minInstrumentalness: number, maxInstrumentalness: number, targetInstrumentalness: number;
  let targetMode: number = 2;
  let minTempo: number, maxTempo: number, targetTempo: number;
  let minTimeSig: number, maxTimeSig: number, targetTimeSig: number;
  let minDuration: number, maxDuration: number, targetDuration: number;
  let minEnergy: number, maxEnergy: number, targetEnergy: number;
  let minDanceability: number, maxDanceability: number, targetDanceability: number;
  let minLiveness: number, maxLiveness: number, targetLiveness: number;
  let minSpeechiness: number, maxSpeechiness: number, targetSpeechiness: number;
  let minPopularity: number, maxPopularity: number, targetPopularity: number;
  let minLoudness: number, maxLoudness: number, targetLoudness: number;

  const modalStore = getModalStore();
  const modalComponent: ModalComponent = { ref: modalComp };

</script>

<div class="space-y-10 w-11/12 m-6 text-left flex flex-col grow items-center" use:loader={loading}>
  <h1 class="h1">Playlist Recommendations</h1>
  <div class="w-10/12">
    <form class="" on:submit|preventDefault={handleSubmit}>
      <Accordion autocollapse>
        <AccordionItem open>
          <svelte:fragment slot="lead">Basic Info:</svelte:fragment>
          <svelte:fragment slot="summary">Artists, Genres, Tracks</svelte:fragment>
          <svelte:fragment slot="content">
            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group input-group-divider grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim">Artists</div>
                  <input type="text" bind:value={artists} placeholder="artistID1,artistID2">
                </div>
                <div class="input-group input-group-divider grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim">Genres</div>
                  <input type="text" bind:value={genres} placeholder="Pop, Rock">
                </div>
                <div class="input-group input-group-divider grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim">Songs</div>
                  <input type="text" bind:value={tracks} placeholder="trackID1,trackID2">
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Max Recs</div>
                  <input class="input text-left w-auto" bind:value={limit} min="1" max="100" type="number" step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">1-100</div>
                </div>
              </div>
            </div>
          </svelte:fragment>
        </AccordionItem>
        <AccordionItem>
          <svelte:fragment slot="lead">Musical Filters:</svelte:fragment>
          <svelte:fragment slot="summary">Key, Mode, Acousticness, Instrumentalness</svelte:fragment>
          <svelte:fragment slot="content">
            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Min Key</div>
                  <input class="input text-left w-auto" bind:value={minKey} min="0" max="11" type="number" step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-11</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Max Key</div>
                  <input class="input text-left w-auto" bind:value={maxKey} min="0" max="11" type="number" step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-11</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Target Key</div>
                  <input class="input text-left w-auto" bind:value={targetKey} min="0" max="11" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-11</div>
                </div>
              </div>
            </div>

            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Min Acousticness</div>
                  <input class="input text-left w-auto" bind:value={minAcousticness} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Max Acousticness</div>
                  <input class="input text-left w-auto" bind:value={maxAcousticness} min="0" max="100"
                         type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Target Acousticness</div>
                  <input class="input text-left w-auto" bind:value={targetAcousticness} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
              </div>
            </div>
            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Min Instrumentalness</div>
                  <input class="input text-left w-auto" bind:value={minInstrumentalness} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Max Instrumentalness</div>
                  <input class="input text-left w-auto" bind:value={maxInstrumentalness} min="0" max="100"
                         type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Target Instrumentalness</div>
                  <input class="input text-left w-auto" bind:value={targetInstrumentalness} min="0" max="100"
                         type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
              </div>
            </div>
            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim">Target Mode</div>
                  <select class="select" bind:value={targetMode}>
                    <option selected value={2}>Both</option>
                    <option value={0}>Minor</option>
                    <option value={1}>Major</option>
                  </select>
                </div>
              </div>
            </div>
          </svelte:fragment>
        </AccordionItem>

        <AccordionItem>
          <svelte:fragment slot="lead">Rhythmic Filters:</svelte:fragment>
          <svelte:fragment slot="summary">Tempo (BPM), Time Signature, Duration (s)</svelte:fragment>
          <svelte:fragment slot="content">
            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Min Tempo</div>
                  <input class="input text-left w-auto" bind:value={minTempo} min="50" max="200" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">50-200</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Max Tempo</div>
                  <input class="input text-left w-auto" bind:value={maxTempo} min="50" max="200" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">50-200</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Target Tempo</div>
                  <input class="input text-left w-auto" bind:value={targetTempo} min="50" max="200" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">50-200</div>
                </div>
              </div>
            </div>
            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Min Time Signature</div>
                  <input class="input text-left w-auto" bind:value={minTimeSig} min="3" max="7" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">X/4</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Max Time Signature</div>
                  <input class="input text-left w-auto" bind:value={maxTimeSig} min="3" max="7"
                         type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">X/4</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Target Time Signature</div>
                  <input class="input text-left w-auto" bind:value={targetTimeSig} min="3" max="7" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">X/4</div>
                </div>
              </div>
            </div>
            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Min Duration</div>
                  <input class="input text-left w-auto" bind:value={minDuration} min="1" max="999" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">1-999</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Max Duration</div>
                  <input class="input text-left w-auto" bind:value={maxDuration} min="1" max="999" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">1-999</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Target Duration</div>
                  <input class="input text-left w-auto" bind:value={targetDuration} min="1" max="999" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">1-999</div>
                </div>
              </div>
            </div>
          </svelte:fragment>
        </AccordionItem>
        <AccordionItem>
          <svelte:fragment slot="lead">Style Filters:</svelte:fragment>
          <svelte:fragment slot="summary">Energy, Danceability, Liveness <span class="italic"> (80 or higher likely live)</span>,
            Speechiness <span class="italic">(0-33 music, 33-66 music and speech, 66-100 mostly speech)</span>
          </svelte:fragment>
          <svelte:fragment slot="content">
            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Min Energy</div>
                  <input class="input text-left w-auto" bind:value={minEnergy} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Max Energy</div>
                  <input class="input text-left w-auto" bind:value={maxEnergy} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Target Energy</div>
                  <input class="input text-left w-auto" bind:value={targetEnergy} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
              </div>
            </div>
            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Min Danceability</div>
                  <input class="input text-left w-auto" bind:value={minDanceability} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Max Danceability</div>
                  <input class="input text-left w-auto" bind:value={maxDanceability} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Target Danceability</div>
                  <input class="input text-left w-auto" bind:value={targetDanceability} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
              </div>
            </div>
            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Min Liveness</div>
                  <input class="input text-left w-auto" bind:value={minLiveness} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Max Liveness</div>
                  <input class="input text-left w-auto" bind:value={maxLiveness} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Target Liveness</div>
                  <input class="input text-left w-auto" bind:value={targetLiveness} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
              </div>
            </div>
            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Min Speechiness</div>
                  <input class="input text-left w-auto" bind:value={minSpeechiness} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Max Speechiness</div>
                  <input class="input text-left w-auto" bind:value={maxSpeechiness} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Target Speechiness</div>
                  <input class="input text-left w-auto" bind:value={targetSpeechiness} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
              </div>
            </div>
          </svelte:fragment>
        </AccordionItem>
        <AccordionItem>
          <svelte:fragment slot="lead">Misc Filters:</svelte:fragment>
          <svelte:fragment slot="summary">Popularity, Loudness (dB)</svelte:fragment>
          <svelte:fragment slot="content">
            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Min Popularity</div>
                  <input class="input text-left w-auto" bind:value={minPopularity} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Max Popularity</div>
                  <input class="input text-left w-auto" bind:value={maxPopularity} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Target Popularity</div>
                  <input class="input text-left w-auto" bind:value={targetPopularity} min="0" max="100" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">0-100</div>
                </div>
              </div>
            </div>
            <div class="flex justify-center items-center mx-auto">
              <div class="w-full grid grid-cols-1 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-2">
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Min Loudness</div>
                  <input class="input text-left w-auto" bind:value={minLoudness} min="-60" max="0" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">-60-0</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Max Loudness</div>
                  <input class="input text-left w-auto" bind:value={maxLoudness} min="-60" max="0" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">-60-0</div>
                </div>
                <div class="input-group grid-cols-[auto_1fr_auto]">
                  <div class="input-group-shim truncate">Target Loudness</div>
                  <input class="input text-left w-auto" bind:value={targetLoudness} min="-60" max="0" type="number"
                         step="1" />
                  <div class="input-group-divider" style="padding-left: 0px">-60-0</div>
                </div>
              </div>
            </div>
          </svelte:fragment>
        </AccordionItem>
      </Accordion>
      {#if error}
        <p class="text-red-500">{errorMessage}</p>
      {/if}
      <button class="btn variant-filled-secondary ml-4" type="submit">
        Search
      </button>
    </form>
  </div>

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
          <tbody class="table-body ">
          {#each Object.values(tracks) as track}
            <tr>
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


<!--<div class="w-9/12">-->
<!--  <Spotify spotifyLink="track/3xh41S4f3xUo4o4Ag8Mpx7" height="200px" width="100%" />-->
<!--</div>-->