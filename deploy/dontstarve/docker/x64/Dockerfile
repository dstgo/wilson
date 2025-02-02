FROM steamcmd/steamcmd:debian

ARG APP_ID=343050
ARG DATA_DIR="/data"

ENV DATA_DIR=${DATA_DIR}
ENV DST_DIR="/var/lib/dst"
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
    # Clean up
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

LABEL maintainer="dstgo" \
      description="Dont Starve Together Dedicated Server for x64 arch"

USER root
WORKDIR ${DST_DIR}/bin64

ENTRYPOINT ["/bin/bash","-c"]
CMD ["cat ${DST_VERSION_FILE}"]