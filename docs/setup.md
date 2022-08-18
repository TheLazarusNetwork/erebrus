# Erebrus Setup Docs

# Watcher Setup

### For Ubuntu 21.04

After placing the .path and .service files in /etc/systemd/system, Run:

1. `sudo systemctl daemon-reload`

2. `sudo systemctl enable wg-watcher.path && sudo systemctl start wg-watcher.path`
   Created symlink /etc/systemd/system/multi-user.target.wants/wg-watcher.path → /etc/systemd/system/wg-watcher.path.

3. `sudo systemctl enable wg-watcher.service && sudo systemctl start wg-watcher.service`
   Created symlink /etc/systemd/system/multi-user.target.wants/wg-watcher.service → /etc/systemd/system/wg-watcher.service.

4. `sudo systemctl status wg-watcher.path`

5. `sudo systemctl status wg-watcher.service`

# WireGuard Setup

1. `sudo apt install wireguard`
2. `modprobe wireguard`
3. `lsmod | grep wireguard`
4. Set the Linux kernel to forward the traffic:

```bash
cat << EOF >> max.conf
net.ipv4.ip_forward=1
net.ipv6.conf.all.forwarding=1
EOF
sysctl -p
```

5. `wg-quick up wg0`
6. `chmod 600 /etc/wireguard/wg0.conf`
7. `systemctl enable wg-quick@wg0`
