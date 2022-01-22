# Erebrus Deployment Docs

## Install and Deploy using binary

1. Make sure all setup were done
2. Download the suitable binary for your operating system from [here](https://github.com/TheLazarusNetwork/erebrus/releases/)
3. create a .env file in same directory and define the environment for erebrus . you can use template from [.sample-env](https://github.com/TheLazarusNetwork/erebrus/blob/main/.sample-env)
4. Run

## Install and Deploy using Docker

1. Make sure all setup were done
2. Pull the ererbus docker image
    ```
    docker pull lazarusnetwork/erebrus:latest
    ```
3. Run the Image
    ```
    docker run -d -p 9080:9080/tcp -p 51820:51820/udp --cap-add=NET_ADMIN --cap-add=SYS_MODULE --sysctl="net.ipv4.conf.all.src_valid_mark=1" --sysctl="net.ipv6.conf.all.forwarding=1" \
    -e LOAD_CONFIG_FILE="FALSE" \
    -e RUNTYPE='debug' \
    -e SERVER='0.0.0.0' \
    -e GRPC_PORT='9080' \
    -e WG_CONF_DIR='/etc/wireguard' \
    -e WG_KEYS_DIR='/etc/wireguard/keys' \
    -e WG_INTERFACE_NAME='wg0.conf' \
    -e WG_ENDPOINT_HOST='your endpoint' \
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
    -e SMTP_PORT='smtp port' \
    -e SMTP_USERNAME='username' \
    -e SMTP_PASSWORD='password' \
    -e SMTP_FROM='from' \
    --restart unless-stopped \
    --name erebrus-region \
    erebrus
    ```
4. Use the following commands   
    ```
    docker exec -it erebrus bash
    ```
    ```
    sudo netstat -pna | grep 51820
    ```
    ```
    sudo lsof -i -P -n | grep 51820
    ```
    ```
    docker rm -f $(docker ps -aq)
    ```
