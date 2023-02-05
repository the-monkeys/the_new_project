#!/bin/bash

source common.sh

# STEP 3: Generate environment vars.
# echo $ALL_ENVS > /the_monkeys/etc/dev.env

# STEP 4: Setup DB (TBD)
# Install postgres and OpenSearch in the deployment container.

# Final step.

# reload systemd manager configuration
systemctl daemon-reload

# start the service
systemctl start $PROGRAM_NAME

# enable the service to start on boot
systemctl enable $PROGRAM_NAME
