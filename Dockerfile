FROM debian AS trust-bootnode

WORKDIR /trust

COPY /bin/linux/bootnode /deployment/boot.key ./

RUN apt-get update && apt-get install -y net-tools iputils-ping curl

ENTRYPOINT ["./bootnode", "-nodekey", "boot.key"]

FROM debian AS trust-node

WORKDIR /trust

COPY /bin/linux/geth /deployment/secret.txt /scripts/res-monitor-local.sh ./

RUN apt-get update && apt-get install -y net-tools iputils-ping curl procps tcpdump #libpcap-dev
#RUN ln -s /usr/lib/libpcap.so /usr/lib/libpcap.so.0.8
