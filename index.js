const INFO = require('./info.json');
const clock = require('./src/pomodor.js');
const math = require('./src/math.js');
const TOKEN = INFO["BOT_TOKEN"];

const Discord = require('discord.js');
const bot = new Discord.Client();


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
    case 'ping': // !ping
      console.log(`command invoked: ${primeCom}`);
      baseComms(receivedMessage);
      break;
    case 'help': // !help 
      console.log(`command invoked: ${primeCom}`);
      baseComms(receivedMessage);
      break;
    case 'heads': // !heads
      console.log(`command invoked ${primeCom}`);
      math.calc.commandEval(receivedMessage, splitCommand);
      break;
    case 'pomodor': // !pomodor 20
      console.log(`command invoked: ${primeCom}`);
      clock.t.commandEval(receivedMessage, splitCommand[1]);
      break;
    case 'add': // !add x + y
      console.log(`command invoked: ${primeCom}`);
      math.calc.commandEval(receivedMessage, splitCommand);
      break;
    case 'minus': //!minus x - y
      console.log(`command invoked: ${primeCom}`);
      math.calc.commandEval(receivedMessage, splitCommand);
      break;
  }
};

// Basic functions for the bot
function baseComms(message) {
  if (message.content == '!ping') {
    // Reply to message and log response time
    message.reply('Pong');
    var now = new Date();
    var timeRn = now.getHours() + ":" + now.getMinutes() + ":" + now.getSeconds();
    // Log time
    console.log(`Ping responded to at ${timeRn}`);
  } else if (message.content == '!help') {
    // Send user link to wiki page
    refLink = 'https://github.com/sir-typesalot/JarvisBot/wiki'
    message.reply(`Here is a simple reference sheet ${refLink}`);
    console.log('Sent Link to Wiki');
  } else {
    console.log('Message not understood');
  }
};

bot.login(TOKEN);