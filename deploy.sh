#!/bin/bash

source common.sh

# change these variables to match your setup
PROGRAM_NAME="api_gateway"
PROGRAM_BINARY="/usr/local/bin/the_monkeys/api_gateway"
PROGRAM_DIR="/usr/local/bin/the_monkeys/"

# create a new user and group for the service
useradd -r -s /bin/false $PROGRAM_NAME

# create a new service file
touch /etc/systemd/system/$PROGRAM_NAME.service

# TODO: Skip if the-monkey is already installed.
# write the service file
cat > /etc/systemd/system/$PROGRAM_NAME.service <<EOF
[Unit]
Description=$PROGRAM_NAME daemon
[Service]
User=$PROGRAM_NAME
WorkingDirectory=$PROGRAM_DIR
ExecStart=$PROGRAM_BINARY
Restart=always
RestartSec=5s
[Install]
WantedBy=multi-user.target
EOF

# reload systemd manager configuration
systemctl daemon-reload

# start the service
systemctl start $PROGRAM_NAME

# enable the service to start on boot
systemctl enable $PROGRAM_NAME
