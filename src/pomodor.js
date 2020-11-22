
let ACTIVE = false;
var user = "name";
var queue = [];

class Timer{
    constructor() {
    }
    commandEval(message, comms) {
        if (comms.toLowerCase() === "stop"){
            this.endTimer(message);
            console.log('Cancelled Pomodor');
        } else if (comms.toLowerCase() === "status") {
            console.log(`Time remaining in Pomodor: ${10}`);
        } else {
            this.timer(message, comms);
            console.log('Called Pomodor');
        }
    };
    timer(message, time) {
        message.reply(`I will set a Pomodor timer for ${time} minutes, starting now`);
        console.log(`Timer for ${time} minutes`);
        var now = (new Date().getTime() / 1000);
        var minsToSecs = function(time) {
            return time * 60;
        };
        let endTime = minsToSecs(time);
        ACTIVE = true;
        this.clock = setInterval(function() {
            var check = (new Date().getTime()/1000);
            if (check >= (now + endTime)) {
                message.reply('Time to take a break!');
                console.log('Time limit reached');
                clearInterval(clock);
            }}, 1000);
    };
    endTimer(message) {
        try{
            clearInterval(this.clock);
            message.reply('Your Pomodor timer has been cancelled...');
            console.log('Pomodor cleared...');
            return 1;
        } catch(err) {
            console.log(`Error processing request: ${err}`);
            return 0;
        }    
    };
    timerStatus() {
        console.log("Pass");
    }
}
module.exports = {
    t : new Timer()
}