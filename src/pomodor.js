module.exports = {
    tracker : function(message, time) {
        message.reply(`I will set a Pomodor timer for ${time} minutes`);
        console.log(`Timer for ${time} minutes`);
    }
}