FROM steamcmd/steamcmd:debian
USER root
WORKDIR /root

ARG APP_ID=343050
ARG DATA_DIR="/root/data"

ENV DST_DIR="/root/dst"
ENV DATA_DIR=${DATA_DIR}
ENV DST_VERSION_FILE="${DST_DIR}/version.txt"
ENV TZ="Asia/Shanghai"

RUN apt update \
    && apt install -y libcurl4-gnutls-dev \
    # Set up Timezone
    && ln -snf /usr/share/zoneinfo/${TZ} /etc/localtime  \
    && echo ${TZ} > /etc/timezone \
    # Create data dir
    && mkdir -p ${DATA_DIR} \
    # Download dedicated server
    && steamcmd +force_install_dir ${DST_DIR} \
        +login anonymous \
        +app_update ${APP_ID} validate \
        +quit \
    # Set up DST_VERSION
    && CUR_VERSION=$(cat ${DST_VERSION_FILE}) \
    && echo "DST_VERSION=${CUR_VERSION}" >> /etc/environment \
    && echo "export DST_VERSION=${CUR_VERSION}" >> ~/.bashrc \
    # Clean up
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

LABEL maintainer="dstgo" \
      description="Dont Starve Together Dedicated Server"

VOLUME ${DATA_DIR}

ENTRYPOINT ["/bin/bash","-c"]
CMD ["steamcmd +help"]