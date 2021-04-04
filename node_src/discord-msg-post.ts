import { Client } from 'discord.js';
import { config } from 'dotenv';

config();

const client = new Client();

client.once('ready', () => {
  console.log('Connected to Discord server');
});

client.login(process.env.DISCORD_BOT_TOKEN);

client.on('message', (message) => {
  if (message.content === 'Bread') message.channel.send('Sucks');
});
