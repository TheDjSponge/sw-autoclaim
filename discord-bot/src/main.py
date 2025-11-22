## Declare bot
import os
import logging
from dotenv import load_dotenv
import discord
from discord.ext import commands
from cogs.user_management import UserManagement

logger = logging.getLogger("sw_discord_bot")


class SWAutoClaimBot(commands.Bot):
    def __init__(self) -> None:
        intents = discord.Intents.default()
        intents.message_content = True
        super().__init__(command_prefix="$", intents=intents)

    async def setup_hook(self):
        await self.load_extension("cogs.user_management")
        await self.tree.sync()

    async def on_ready(self):
        logger.info(f"Logged in as {self.user} (ID: {self.user.id})")


def main():
    load_dotenv()
    bot_key = os.environ.get("DISCORD_BOT_KEY")
    bot = SWAutoClaimBot()
    bot.run(bot_key)


if __name__ == "__main__":
    logging.basicConfig(level=logging.DEBUG)
    main()
