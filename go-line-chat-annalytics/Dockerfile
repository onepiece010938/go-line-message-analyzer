FROM ubuntu:20.04

ARG DOCKER_GID

EXPOSE 22 7777 3000

# For timezone
RUN apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends tzdata \
    && TZ=Asia/Taipei \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && dpkg-reconfigure -f noninteractive tzdata

#install ssh docker git curl protobuf-compiler golang-goprotobuf-dev
#add user ubuntu pid 1000 join sudo , add workspace and chown prmisson , ssh always start

RUN apt-get update && apt-get install -y sudo openssh-server docker.io git curl gcc protobuf-compiler protobuf-compiler \
    && service ssh start \
    && groupadd -g 1000 ubuntu \
    && useradd -rm -d /home/ubuntu -s /bin/bash -g 1000 -u 1000 ubuntu  -p "$(openssl passwd -1 ubuntu)" \
    && usermod -aG sudo ubuntu && mkdir workspace && chown -R ubuntu:ubuntu workspace \
    && touch /home/ubuntu/start_ssh.sh && echo "#!/bin/bash" && echo "echo 'ubuntu' | sudo -S service ssh start" >> /home/ubuntu/start_ssh.sh \
    && echo ". /home/ubuntu/start_ssh.sh" >> /home/ubuntu/.bashrc

WORKDIR /workspace

## install golang 1.19.2
RUN wget https://go.dev/dl/go1.19.2.linux-amd64.tar.gz \
    && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz \
    && echo "export PATH=$PATH:/usr/local/go/bin" >> /home/ubuntu/.bashrc  \
    && rm -r go1.19.2.linux-amd64.tar.gz

# install caddy
RUN sudo apt install -y debian-keyring debian-archive-keyring apt-transport-https curl \
    && curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/gpg.key' | sudo gpg --dearmor -o /usr/share/keyrings/caddy-stable-archive-keyring.gpg\
    && curl -1sLf 'https://dl.cloudsmith.io/public/caddy/stable/debian.deb.txt' | sudo tee /etc/apt/sources.list.d/caddy-stable.list \
    && sudo apt update && sudo apt install caddy

## change user and install yarn
USER ubuntu
RUN curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.39.1/install.sh | bash
ENV NVM_DIR /home/ubuntu/.nvm
ENV NODE_VERSION v16.14.0
RUN /bin/bash -c "source $NVM_DIR/nvm.sh && nvm install $NODE_VERSION && nvm use --delete-prefix $NODE_VERSION"
ENV NODE_PATH $NVM_DIR/versions/node/$NODE_VERSION/lib/node_modules
ENV PATH      $NVM_DIR/versions/node/$NODE_VERSION/bin:$PATH
RUN npm install -g yarn

#SWAGGER
ENV PATH=/usr/local/go/bin:$PATH
RUN go install github.com/swaggo/swag/cmd/swag@latest \
&& echo "export PATH=/home/ubuntu/go/bin:$PATH" >> /home/ubuntu/.bashrc

# go get -u github.com/go-echarts/go-echarts/v2/...