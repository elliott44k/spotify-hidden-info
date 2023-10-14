<script>
  import { fade } from 'svelte/transition'
</script>

<div class="wrapper" transition:fade>
  <h1>
    <div class="loader"/>
  </h1>
</div>

<style>
    .wrapper {
        display: flex;
        flex-direction: column;
        align-items:center;
        justify-content: center;
        position: absolute;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
        color: slategray;
        background: rgb(33 33 33 / .46)
    }

    /* spinner from https://projects.lukehaas.me/css-loaders/ */
    .loader,
    .loader:after {
        border-radius: 50%;
        width: 8em;
        height: 8em;
    }
    .loader {
        margin: 60px auto;
        font-size: 8px;
        position: relative;
        text-indent: -9999em;
        border-top: 1.1em solid rgba(255, 255, 255, 0.2);
        border-right: 1.1em solid rgba(255, 255, 255, 0.2);
        border-bottom: 1.1em solid rgba(255, 255, 255, 0.2);
        border-left: 1.1em solid #ffffff;
        -webkit-transform: translateZ(0);
        -ms-transform: translateZ(0);
        transform: translateZ(0);
        -webkit-animation: load8 1.1s infinite linear;
        animation: load8 1.1s infinite linear;
    }
    @-webkit-keyframes load8 {
        0% {
            -webkit-transform: rotate(0deg);
            transform: rotate(0deg);
        }
        100% {
            -webkit-transform: rotate(360deg);
            transform: rotate(360deg);
        }
    }
    @keyframes load8 {
        0% {
            -webkit-transform: rotate(0deg);
            transform: rotate(0deg);
        }
        100% {
            -webkit-transform: rotate(360deg);
            transform: rotate(360deg);
        }
    }

</style>