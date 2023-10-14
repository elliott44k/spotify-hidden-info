<!-- YOU CAN DELETE EVERYTHING IN THIS PAGE -->

<div class="space-y-10 text-center flex flex-col items-center">
  <div class="mt-10 justify-center space-x-2">
    <h2 class="h2">Welcome to the Spotify Hidden Info Search</h2>
    <p class="pt-5">A place to generate playlist recommendations</p>
    <p class="p-">and see Spotify's audio analysis of the songs you love</p>
  </div>
  <div class="justify-center items-center space-x-2">
    <p>Try finding hidden info on your favorite song:</p>
    <a class="mt-2 btn variant-filled" href="/song_search">Song Search</a>
    <p class="m-2">Or perhaps you want to generate a playlist based on your favorite artists?</p>
    <a class="btn variant-filled" href="/playlist_generator">Playlist Generator</a>
    <p class="m-2">If you want to generate Artist based recommendations, you'll need to check this page to get their IDs</p>
    <a class="btn variant-filled" href="/artist_id_search">Artist ID Lookup</a>

  </div>
  <div class="space-y-2">
  </div>
</div>

<style lang="postcss">
    figure {
        @apply flex relative flex-col;
    }

    figure svg,
    .img-bg {
        @apply w-64 h-64 md:w-80 md:h-80;
    }

    .img-bg {
        @apply absolute z-[-1] rounded-full blur-[50px] transition-all;
        animation: pulse 5s cubic-bezier(0, 0, 0, 0.5) infinite,
        glow 5s linear infinite;
    }

    @keyframes glow {
        0% {
            @apply bg-primary-400/50;
        }
        33% {
            @apply bg-secondary-400/50;
        }
        66% {
            @apply bg-tertiary-400/50;
        }
        100% {
            @apply bg-primary-400/50;
        }
    }

    @keyframes pulse {
        50% {
            transform: scale(1.5);
        }
    }
</style>
