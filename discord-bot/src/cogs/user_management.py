## Implement commands
import discord
from discord.ext import commands
from discord import app_commands
import logging
from typing import Literal

logger = logging.getLogger("sw_discord_bot.user_management")


class UserManagement(commands.Cog):
    """Cog extension that allows to manage SW users and store them into a database."""

    def __init__(self, bot: commands.Bot) -> None:
        """Initializes a UserManagement Cog extension.

        Args:
            bot (commands.Bot): bot instance to which the cog is added.
        """
        self.bot = bot

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

        await interaction.followup.send(
            f"Registered Hive ID **{hive_user_id}** for server **{server}**!"
        )
        print(interaction)
        logger.info(
            f"User {interaction.user} registered Hive ID {hive_user_id} for server {server}"
        )
        dummy_users_db_entry = {
            "server": server,
            "hiveid": hive_user_id,
            "active": True,
            "discord_id": interaction.user.id,
        }
        logger.debug(
            f"DUMMY LOG: ADDING ENTRY TO USERS DB WITH : \n {dummy_users_db_entry}"
        )
        dummy_discord_db_entry = {
            "discord_id": interaction.user.id,
            "discord_name": str(interaction.user),
        }
        logger.debug(
            f"DUMMY LOG: ADDING ENTRY TO DISCORD DB WITH : \n {dummy_discord_db_entry}"
        )


async def setup(bot: commands.Bot):
    """Adds the extension to a bot.

    Args:
        bot (commands.Bot): Bot instance to which the extension is added.
    """
    logger.debug("Loading UserManagement extension cog")
    await bot.add_cog(UserManagement(bot))
