# Spotify Hidden Info

*Nov 2024*

*Unbeknownst to me, Spotify deprecated most of the API endpoints used in this project. While basic song search still works, most other functions will return an error 403 from spotify and my little application does not have a loading timeout or error handling for a Spotify 403 built in.*

This is a simple app that uses the Spotify API to get Spotify's song analysis as well as generate song recommendations based on user input. It uses the [Spotify](https://developer.spotify.com/) REST API and iFrame Embed.

Play with the app here (https://spotify-hidden-info.vercel.app/)

## Built with
Svelte - Frontend framework

Svelte Kit - Used for deploying to Vercel 

Go - Backend server technology deployed via Vercel Serverless Functions

Vercel - Hosting and deployment

Redis - Holds Spotify API tokens

Spotify API
