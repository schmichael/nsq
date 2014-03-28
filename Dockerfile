# Ubuntu 12.04, Go 1.2.1, NSQ 0.2.27 (from source)
# Example commands:
#  docker run \
#    -p 4160:4160 -p 4161:4161 \
#    --name nsqlookupd schmichael/nsq nsqlookupd
#  docker run \
#    -p 4150:4150 -p 4151:4151 \
#    -v /tmp:/data --name nsqd schmichael/nsq \
#      nsqd --lookupd-tcp-address=$NSQLOOKUPD_IP:4160 --data-path=/data
#  docker run \
#    -p 4171:4171 \
#    --name nsqadmin schmichael/nsq \
#      nsqadmin --lookupd-http-address=192.168.1.101:4161
FROM schmichael/ubuntu-go
MAINTAINER Michael Schurter <schmichael@lytics.io>
RUN go get github.com/kr/godep
ADD . /opt/go/src/github.com/bitly/nsq
RUN cd /opt/go/src/github.com/bitly/nsq && git checkout v0.2.27 && godep get ./...
