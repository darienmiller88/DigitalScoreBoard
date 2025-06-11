# Score board for Jeopardy
[![Netlify Status](https://api.netlify.com/api/v1/badges/1715535d-5c1e-46d4-80e2-fba5816cf2ca/deploy-status)](https://app.netlify.com/sites/messengerv2/deploys)

I volunteered to assist a coworker by designing a scoreboard to efficiently track scores for Jeopardy games at our organization. The app allows adding players to different ADAPT locations, creating and managing teams, saving game results, and viewing every game ever played, with useful information like the total amount of points scored, the winner of the game! The core functionality is complete, and I'm currently finalizing the team features and polishing the UI to perfect it for another upcoming tournament.

Here is a brief summary of the impact this app had:

* Replaced paper based scorekeeping with a digital system.
* Saved dozens of staff hours over the course of the event due to simplified score keeping.
* Enabled consistent and simple score keeping, complete with timestamps and winners!
* Efficient score keeping allowed rounds to be lengtened from 45 to 60 minutes.
* Successfully used in a citywide, multi-borough tournament streamed to hundreds using company morning show.

### Built With:
* [Vue](https://vuejs.org/)
* [Go](https://github.com/go-chi/chi)
* [MongoDB-Atlas](https://www.mongodb.com/cloud/atlas)
* [Google Cloud Services](https://cloud.google.com/)
* [Netlify](https://www.netlify.com/)

## Home Page Dark Mode, team mode
<img width="959" alt="image" src="https://github.com/user-attachments/assets/4e030cc6-a890-4872-b207-5043f5a33748" />

## Home Page Light Mode, team mode
<img width="955" alt="image" src="https://github.com/user-attachments/assets/087d5c21-9542-4d82-9650-750a84741d5c" />

## Home Page, Add new user
<img width="955" alt="image" src="https://github.com/user-attachments/assets/72e1446c-1c56-42c2-9062-8ee77d5050f1" />

## Add player to team
<img width="956" alt="image" src="https://github.com/user-attachments/assets/3d6c8a6b-9052-4542-8830-6a027da0d0e7" />

## View team
<img width="959" alt="image" src="https://github.com/user-attachments/assets/d1d37f6c-38e8-427e-a8a8-591eb49053f1" />

## Saved games
<img width="959" alt="image" src="https://github.com/user-attachments/assets/f28c5dae-dcfb-493c-befd-532216690329" />

## Features
So far, this application has the following features:

* Single player Jeopardy. Client can add and remove Users to a specific ADAPT location, and add and subtract points to them.
* Tournament/Team jeopardy mode where the client can choose the host location for the game, and which teams are playing.
* The ability to save games, both single player and team jeopardy.
* Viewing games, which will include the winner of the game, and the total and average amount of points earned during the game.

## View api here:
https://adaptscoreboard-678166011633.us-east4.run.app/

I made the choice to migrate the backend from fly.io to Google cloud due to cost overruns with the former, and the latter providing a free tier. Thankfully, deploying to Google Cloud was not hard at all! The features provided for the free tier seem to be reasonable too.

### Requirements:

* Clone repo using `git clone https://github.com/darienmiller88/DigitalScoreBoard`
* Migrate the necessary information to your local `.env` as described in the `.env_sample` file
* Run go build to create a root level `DigitalScoreBoard.exe` file, and then run `.\DigitalScoreBoard` to run the executable. If an executable is not needed, simply input `go run main.go` instead, or `.\fresh` to enable a server restart on change.
* `cd client` to access the Vue package, and run `npm run dev` to start vite server for Vue, which should be on port `5173`.

*NOTE: Please run backend before frontend as there are onLoad calls that the frontend makes to the backend that will fail if the front end is run first.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Feel free to leave suggestions as well, I'm always looking for ways to improve!

<p align="right">(<a href="#top">back to top</a>)</p>

## License
[MIT](https://choosealicense.com/licenses/mit/)