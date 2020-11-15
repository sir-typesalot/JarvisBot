const INFO = require('./info.json');
const Discord = require('discord.js');
const bot = new Discord.Client();
const TOKEN = INFO["BOT_TOKEN"];

bot.on('ready', () => {
  var d = new Date();
  var n = d.getHours() + ":" + d.getMinutes() + ":" + d.getSeconds();
  console.info(`Logged in as ${bot.user.tag}!`);
  console.log(`Time : ${n}`);
});

bot.on('message', msg => {
  if (msg.author == bot.user){
    return
  } else {
    processCommand(msg);
  } 
});

function processCommand(receivedMessage) {
  let fullCommand = receivedMessage.content.substr(1)
  let splitCommand = fullCommand.split(" ");
  let primeCom = splitCommand[0];
  if (primeCom == 'ping') {
    ping(receivedMessage);
  } else {
    receivedMessage.reply('I did not understand what you said...');
  }
};
function ping(message) {
  if (message.content == '!ping') {
    message.reply('Pong');
    var now = new Date();
    var timeRn = now.getHours() + ":" + now.getMinutes() + ":" + now.getSeconds();
    console.log(`Ping responded to at ${timeRn}`);
  } else {
    console.log('Message not understood')
  }
};

bot.login(TOKEN);