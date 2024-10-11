# Score board for Jeopardy
[![Netlify Status](https://api.netlify.com/api/v1/badges/1715535d-5c1e-46d4-80e2-fba5816cf2ca/deploy-status)](https://app.netlify.com/sites/messengerv2/deploys)

I am volunteering to assist a co worker in designing them a score board to more efficiently keep track of the 
score of everyone participating!

### Built With:
* [Vue](https://vuejs.org/)
* [Go](https://github.com/go-chi/chi)
* [MongoDB-Atlas](https://www.mongodb.com/cloud/atlas)

## Home Page
<img width="959" alt="project-screenshot" src="https://github.com/user-attachments/assets/f12d1f05-8dc5-43ac-a59e-88cb9b68330d">


### Requirements:

* Clone repo using `git clone https://github.com/darienmiller88/DigitalScoreBoard`
* Migrate the necessary information to your local `.env` as described in the `.env_sample` file
* Run go build to create a root level `DigitalScoreBoard.exe` file, and then run `.\DigitalScoreBoard` to run the executable. If an executable is not needed, simply input `go run main.go` instead, or `.\fresh` to enable a server restart on change.
* `cd client` to access the Vue package, and run `npm run dev` to start vite server for Vue, which should be on port `5173`.

### Notes
I do have a Dockerfile for my go part of this project, but including causes the application to use it to deploy instead of pushes to github for whatever reason, so I'm not including it for now.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Feel free to leave suggestions as well, I'm always looking for ways to improve!

<p align="right">(<a href="#top">back to top</a>)</p>

## License
[MIT](https://choosealicense.com/licenses/mit/)