FROM ubuntu:20.04
EXPOSE 22 
# For timezone
RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends tzdata \
    && TZ=Asia/Taipei \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && dpkg-reconfigure -f noninteractive tzdata
#install ssh docker git curl protobuf-compiler golang-goprotobuf-dev
#add user ubuntu pid 1000 join sudo , add workspace and chown prmisson , ssh always start
RUN apt-get update && apt-get install -y sudo openssh-server git curl gcc protobuf-compiler golang-goprotobuf-dev \
    && service ssh start \
    && groupadd -g 1000 ubuntu \
    && useradd -rm -d /home/ubuntu -s /bin/bash -g 1000 -u 1000 ubuntu  -p "$(openssl passwd -1 ubuntu)" \
    && usermod -aG sudo ubuntu && mkdir workspace && chown -R ubuntu:ubuntu workspace \
    && touch /home/ubuntu/start_ssh.sh && echo "#!/bin/bash" && echo "echo 'ubuntu' | sudo -S service ssh start" >> /home/ubuntu/start_ssh.sh \
    && echo ". /home/ubuntu/start_ssh.sh" >> /home/ubuntu/.bashrc 
# For timezone

WORKDIR /workspace
## git toolkits and install golang 1.17.6
RUN wget https://go.dev/dl/go1.19.5.linux-amd64.tar.gz \
    && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.5.linux-amd64.tar.gz \
    && rm -r go1.19.5.linux-amd64.tar.gz 
    && echo "export PATH=$PATH:/usr/local/go/bin" >> /home/ubuntu/.bashrc  

## change user and install yarn 
USER ubuntu

ENV PATH=/usr/local/go/bin:/home/ubuntu/go/bin:$PATH 
RUN echo "export PATH=/home/ubuntu/go/bin:$PATH" >> /home/ubuntu/.bashrc 

# RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2
# # SQLC
# RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.15.0
# # Mockgen
# RUN go install github.com/golang/mock/mockgen@v1.6.0

RUN go get github.com/allegro/bigcache/v3
RUN go get -u github.com/go-ego/gse
RUN go get github.com/bytedance/sonic
RUN go get  github.com/stretchr/testify/require
#ENTRYPOINT service ssh restart && bash
#CMD ["/usr/sbin/sshd", "-D"]