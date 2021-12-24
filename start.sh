#!/usr/bin/env bash
set -eo pipefail

modprobe wireguard

cat <<EOF >>/etc/sysctl.conf
net.ipv4.ip_forward=1
net.ipv6.conf.all.forwarding=1
EOF

wg-quick up wg0
chmod 600 /etc/wireguard/wg0.conf

mkdir -p $WG_KEYS_DIR
/app/erebrus &
./wg-watcher.sh
sleep infinity
