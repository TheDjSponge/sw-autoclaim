#!/bin/bash

/app/scraper

echo "Starting cronjob with supercronic"
exec supercronic /etc/cron.d/scraper-cron