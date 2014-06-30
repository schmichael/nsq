# Ubuntu 14.04, Go 1.3, NSQ 0.2.28 (from source)
# Example commands:
#  docker run \
#    -p 4160:4160 -p 4161:4161 \
#    --net="host" \
#    --name nsqlookupd schmichael/nsq nsqlookupd
#  docker run \
#    -p 4150:4150 -p 4151:4151 \
#    --net="host" \
#    -v /tmp:/data --name nsqd schmichael/nsq \
#      nsqd --lookupd-tcp-address=$NSQLOOKUPD_IP:4160 --data-path=/data
#  docker run \
#    -p 4171:4171 \
#    --net="host" \
#    --name nsqadmin schmichael/nsq \
#      nsqadmin --lookupd-http-address=$NSQLOOKUPD_IP:4161
FROM schmichael/ubuntu-go:v14.04-1.3.0
MAINTAINER Michael Schurter <schmichael@lytics.io>
RUN go get github.com/kr/godep
ADD . /opt/go/src/github.com/bitly/nsq
RUN cd /opt/go/src/github.com/bitly/nsq && git checkout v0.2.28 && godep get ./apps/... && godep get ./nsqadmin
