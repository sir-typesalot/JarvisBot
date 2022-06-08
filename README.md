# JarvisBot
Simple Discord Bot with Multiple Functionalities. Originally just a side project to tinker with Discord's API, I quickly found that I may have more uses to it than I thought. With an expanding list of functionalities, I am working to make this more usable and appealing for everyone to enhance many more Discord channels.

## Features added so far ##
For more details on each of the commands, you can visit the [Wiki](https://github.com/sir-typesalot/JarvisBot/wiki) for more info.
- Pomodor timer
- Virtual Heads or Tails
- Simple Addition and Subtraction
- Active Minutes scoreboard (v1 out!)
- Stock Market Summary (In Progress)

Many more to come over the next few months!

## Architecture ##
JarvisBot is written in Go, and is hosted on a Kamatera VPS. The release process is handled by a couple of bash scripts that automate the process. JarvisBot relies on external APIs for several of it's features, like the Activity Scoreboard ([Exercise-API](https://github.com/sir-typesalot/ExerciseFinder-API)) and the upcoming Stock Monitor ([Polygon-API](https://polygon.io/)).

### Local Development ###
To run this project locally, you will need GO version > 1.17.    
- First, install the required packages, run `go get -u -v ./`    
- To run this project, navigate to the root directory of this project and run `sh scripts/run_local.sh` This will compile and run an executable that will be stored in the `bin` directory of this project.
