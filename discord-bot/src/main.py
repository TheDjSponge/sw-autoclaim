## Declare bot
import os
import logging

import discord
from discord.ext import commands
from dotenv import load_dotenv

logger = logging.getLogger("sw_discord_bot")


class SWAutoClaimBot(commands.Bot):
    """Discord bot used to register Summoners war users for automatic coupon claiming."""

    def __init__(self) -> None:
        """Initialises an SWAutoClaimBot"""
        intents = discord.Intents.default()
        intents.message_content = True
        super().__init__(command_prefix="$", intents=intents)

    async def setup_hook(self) -> None:
        """Hook called before the bot is ready for interaction. Loads extensions and syncs the command tree."""
        await self.load_extension("cogs.user_management")
        await self.tree.sync()

    async def on_ready(self) -> None:
        """Event called whenever the bot is ready for usage."""
        logger.info(f"Logged in as {self.user} (ID: {self.user.id})")


def main():
    load_dotenv()
    bot_key = os.environ.get("DISCORD_BOT_KEY")
    bot = SWAutoClaimBot()
    bot.run(bot_key)


if __name__ == "__main__":
    logging.basicConfig(level=logging.DEBUG)
    main()
