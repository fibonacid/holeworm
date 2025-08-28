#!/bin/bash
#
# This script uses magic-wormhole to transfer files between two hosts with a flipped interaction.
# Flow:
# 1. Host A specifies a path on Host B and forwards it to a web server with a public IP.
# 2. Web server sends an event to an agent running on Host B.
# 3. Agent notifies user and waits for acceptance/rejection.
# 4. If accepted, Host B uses magic-wormhole to receive the file.
# 5. Web server provides the wormhole code back to Host A.
# 6. Host A uses the code to transfer file directly to Host B.
# The transfer happens peer-to-peer using magic-wormhole.
#
