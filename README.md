# erebrus
Anonymous Virtual Private Network for accessing internet in stealth mode bypassing filewalls and filters

# Watcher Setup

### For Ubuntu 21.04

After placing the .path and .service files in /etc/systemd/system, Run:

1. ```sudo systemctl daemon-reload```

2. ```sudo systemctl enable wg-watcher.path && sudo systemctl start wg-watcher.path```
Created symlink /etc/systemd/system/multi-user.target.wants/wg-watcher.path → /etc/systemd/system/wg-watcher.path.

3. ```sudo systemctl enable wg-watcher.service && sudo systemctl start wg-watcher.service```
Created symlink /etc/systemd/system/multi-user.target.wants/wg-watcher.service → /etc/systemd/system/wg-watcher.service.

4. ```sudo systemctl status wg-watcher.path```

5. ```sudo systemctl status wg-watcher.service```

# WireGuard Setup

1. ```sudo apt install wireguard```
2. ```modprobe wireguard```
3. ```lsmod | grep wireguard```
4. Set the Linux kernel to forward the traffic:
    >> cat << EOF >> /etc/sysctl.conf
    >> net.ipv4.ip_forward=1
    >> net.ipv6.conf.all.forwarding=1
    >> EOF
    >> sysctl -p
5. ```wg-quick up wg0```
6. ```chmod 600 /etc/wireguard/wg0.conf```
7. ```systemctl enable wg-quick@wg0```

# Docker Setup

### docker-cli

1. ```docker build -t erebrus .```
2. ```
    docker run -d -p 9080:9080/tcp -p 51820:51820/udp --cap-add=NET_ADMIN --cap-add=SYS_MODULE --sysctl="net.ipv4.conf.all.src_valid_mark=1" --sysctl="net.ipv6.conf.all.forwarding=1" \
    -e LOAD_CONFIG_FILE="FALSE" \
    -e RUNTYPE='debug' \
    -e SERVER='0.0.0.0' \
    -e GRPC_PORT='9080' \
    -e WG_CONF_DIR='/etc/wireguard' \
    -e WG_KEYS_DIR='/etc/wireguard/keys' \
    -e WG_INTERFACE_NAME='wg0.conf' \
    -e WG_ENDPOINT_HOST='region.lazarus.network' \
    -e WG_ENDPOINT_PORT='51820' \
    -e WG_IPv4_SUBNET='10.0.0.1/24' \
    -e WG_IPv6_SUBNET='fd9f:0000::10:0:0:1/64' \
    -e WG_DNS='1.1.1.1' \
    -e WG_ALLOWED_IP_1='0.0.0.0/0' \
    -e WG_ALLOWED_IP_2='::/0' \
    -e WG_PRE_UP='echo WireGuard PreUp' \
    -e WG_POST_UP='iptables -A FORWARD -i %i -j ACCEPT; iptables -A FORWARD -o %i -j ACCEPT; iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE' \
    -e WG_PRE_DOWN='echo WireGuard PreDown' \
    -e WG_POST_DOWN='iptables -D FORWARD -i %i -j ACCEPT; iptables -D FORWARD -o %i -j ACCEPT; iptables -t nat -D POSTROUTING -o eth0 -j MASQUERADE' \
    -e SMTP_HOST='smtp.mail-domain.com' \
    -e SMTP_PORT='465' \
    -e SMTP_USERNAME='erebrus@lazarus.network' \
    -e SMTP_PASSWORD='erebrus' \
    -e SMTP_FROM='Lazarus Network - Erebrus <erebrus@lazarus.network>' \
    --restart unless-stopped \
    --name erebrus-region \
    erebrus
    ```
3. ```docker exec -it erebrus bash```
4. ```sudo netstat -pna | grep 51820```
5. ```sudo lsof -i -P -n | grep 51820```
6. ```docker rm -f $(docker ps -aq)```