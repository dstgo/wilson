FROM steamcmd/steamcmd:debian

ARG APP_ID=343050
ARG DATA_DIR="/data"

ENV DATA_DIR=${DATA_DIR}
ENV DST_DIR="/var/lib/dst"
ENV DST_VERSION_FILE="${DST_DIR}/version.txt"
ENV TZ="Asia/Shanghai"

RUN dpkg --add-architecture i386  \
    # Prepare i386 dependency
    && apt update\
    && apt -y install lib32z1 lib32stdc++6 libcurl4-gnutls-dev:i386 \
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
      description="Dont Starve Together Dedicated Server for x86 arch"

USER root
WORKDIR ${DST_DIR}/bin/

ENTRYPOINT ["/bin/bash","-c"]
CMD ["cat ${DST_VERSION_FILE}"]