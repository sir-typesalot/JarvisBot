const INFO = require('./info.json');
const clock = require('./src/pomodor.js');
const Discord = require('discord.js');
const bot = new Discord.Client();
const TOKEN = INFO["BOT_TOKEN"];
var commList = ['ping', 'time']

bot.on('ready', () => {
  var d = new Date();
  var n = d.getHours() + ":" + d.getMinutes() + ":" + d.getSeconds();
  console.info(`Logged in as ${bot.user.tag}!`);
  console.log(`Time : ${n}`);
});

bot.on('message', msg => {
  if (msg.author == bot.user || msg.content.charAt(0) != '!') {
    return
  } else {
    processCommand(msg);
  } 
});

function processCommand(receivedMessage) {
  let fullCommand = receivedMessage.content.substr(1)
  let splitCommand = fullCommand.split(" ");
  let primeCom = splitCommand[0];
  
  switch (primeCom){
    case 'ping':
      console.log(`command invoked: ${primeCom}`);
      baseComms(receivedMessage);
      break;
    case 'time':
      console.log(`command invoked: ${primeCom}`);
      baseComms(receivedMessage);
      break;
    case 'ideas':
      console.log(`command invoked: ${primeCom}`);
      postLink(receivedMessage);
    case 'pomodor':
      console.log(`command invoked: ${primeCom}`);
      clock.t.commandEval(receivedMessage, splitCommand[1]);
      break;
  }
};
function baseComms(message) {
  if (message.content == '!ping') {
    message.reply('Pong');
    var now = new Date();
    var timeRn = now.getHours() + ":" + now.getMinutes() + ":" + now.getSeconds();
    console.log(`Ping responded to at ${timeRn}`);
  } else if (message.content == '!time') {
    var rn = new Date();
    var rightNow = rn.getHours() + ":" + rn.getMinutes() + ":" + rn.getSeconds();
    message.reply(`The time is ${rightNow}`);
    if (rn.getHours() >= 20 || rn.getHours() <= 5) {
      message.channel.send("You should get some rest");
    } else {
      message.channel.send("Hope you're having a great day");
    }
  } else {
    console.log('Message not understood');
  }
};

function postLink(message) {
  if (message.content == '!ideas') {
    let link = 'https://www.mensjournal.com/health-fitness/top-workout-routines-according-science/'
    message.reply(`Here is a link ${link}`);
    console.log('Sent link');
  }
};

bot.login(TOKEN);