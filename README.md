# Project: VibeShare

## Table of Contents

- [Project: VibeShare](#project-vibeshare)
  - [Table of Contents](#table-of-contents)
  - [Project Description](#project-description)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Running the Application Locally](#running-the-application-locally)
  - [Members](#members)
    - [Frontend](#frontend)
    - [Backend](#backend)
  - [Source and Documentation](#source-and-documentation)

## Project Description

VibeShare is a social media application that allows users to collaborate on Spotify playlists, comment, share, and more. The app also generates song suggestions for the combined playlist.

Some of the specific features that the project includes are:

- Developing group-centered playlists for events
- Voting on songs to add to the group's playlists
- Connect with friends who also have Spotify accounts

The goal of VibeShare is to increase the enjoyment level for parties and other events by creating playlists that are designed based on the environment and user similarities.

## Getting Started

### Prerequisites

- Download all necessary dependencies
- Open your Firestore database and get JSON Google credentials, then store them in a separate section of your local drive

### Running the Application Locally

1. Set the environment variable for Firestore credentials:

   ```
   export GOOGLE_APPLICATION_CREDENTIALS=path/to/JSON
   ```

2. Build the Go application:

   ```
   go build
   ```

3. Run the server:

   ```
   go run *
   ```

4. Start the Angular frontend:

   ```
   npm start
   ```

5. Follow instructions as prompted.

## Members

### Frontend

- Akshat Pant
- Michael Shaffer

### Backend

- Aryaan Verma
- Martin Kent

## Source and Documentation

- [Angular Documentation](https://angular.io/docs)
- [Go Documentation](https://golang.org/doc/)
- [Firestore Documentation](https://firebase.google.com/docs/firestore)
- [Spotify API Documentation](https://developer.spotify.com/documentation/web-api/)
- [Cypress Documentation](https://docs.cypress.io/)
