## Implement commands
import discord
from discord.ext import commands
from discord import app_commands
import logging
from typing import Literal
import os
import json
import requests

logger = logging.getLogger("sw_discord_bot.user_management")


class UserManagement(commands.Cog):
    """Cog extension that allows to manage SW users and store them into a database."""

    def __init__(self, bot: commands.Bot) -> None:
        """Initializes a UserManagement Cog extension.

        Args:
            bot (commands.Bot): bot instance to which the cog is added.
        """
        self.bot = bot
        self.backend_api_url = os.environ.get("REDEMPTION_SERVICE_URL", "")
        if self.backend_api_url == "":
            logger.error("No backend api url provided")
        self.users_api_url = self.backend_api_url + "/v1/users"

    @app_commands.command(
        name="register",
        description="Registers a hive user-id for automatic coupon claiming",
    )
    @app_commands.describe(
        hive_user_id="Your Hive user ID",
        server="Your game server",
    )
    async def register(
        self,
        interaction: discord.Interaction,
        hive_user_id: str,
        server: Literal["global", "korea", "japan", "china", "asia", "europe"],
    ):
        """Registers a hive user-id as well as relevant information into a database."""
        await interaction.response.defer(ephemeral=True)

        print(interaction)
        logger.info(
            f"User {interaction.user} registered Hive ID {hive_user_id} for server {server}"
        )
        user_entry = {
            "server": server,
            "hive_id": hive_user_id,
            "discord_id": interaction.user.id,
            "discord_username": str(interaction.user),
        }
        logger.debug(f"Posting to users api: {user_entry}")
        try:
            response = requests.post(
                self.users_api_url,
                data=json.dumps(user_entry),
                timeout=10,
            )
            response.raise_for_status()

        except requests.exceptions.HTTPError as err:
            logger.error(f"Got error while posting to users api: {err}")
            await interaction.followup.send(
                f"Error while registering Hive ID **{hive_user_id}** for server **{server}**"
            )
            return
        except Exception as err:
            logger.error(f"Got general exeption while posting to users api: {err}")
            return
        logger.info(
            f"User {interaction.user} registered Hive ID {hive_user_id} for server {server}"
        )
        await interaction.followup.send(
            f"Registered Hive ID **{hive_user_id}** for server **{server}**!"
        )


async def setup(bot: commands.Bot):
    """Adds the extension to a bot.

    Args:
        bot (commands.Bot): Bot instance to which the extension is added.
    """
    logger.debug("Loading UserManagement extension cog")
    await bot.add_cog(UserManagement(bot))
