FROM        ubuntu:latest
# install oracle
RUN         mkdir -p /opt/oracle && \
            cd /opt/oracle && \
            apt-get update && \
            apt-get install -y libaio1 wget unzip && \
            wget https://download.oracle.com/otn_software/linux/instantclient/218000/instantclient-basic-linux.x64-21.8.0.0.0dbru.zip && \
            unzip instantclient-basic-linux.x64-21.8.0.0.0dbru.zip && \
            rm -f instantclient-basic-linux.x64-21.8.0.0.0dbru.zip && \
            cd /opt/oracle/instantclient_21_8 && \
            rm -f *jdbc* *occi* *mysql* *mql1* *ipc1* *jar uidrvci genezi adrci && \
            echo /opt/oracle/instantclient_21_8 > /etc/ld.so.conf.d/oracle-instantclient.conf && \
            export LD_LIBRARY_PATH=/opt/oracle/instantclient_21_8:$LD_LIBRARY_PATH && \
            ldconfig
# install time zone
RUN         apt-get install tzdata
# install Node.js
RUN         apt-get install -y curl supervisor && \
            curl -sL https://deb.nodesource.com/setup_lts.x | bash - && \
            apt-get install -y nodejs
# copy file
COPY        jgo /app/
COPY        dist/ /app/dist/
COPY        dao/ /app/dao/
COPY        config/ /app/config/
COPY        sitemap.xml /app/sitemap.xml
COPY        supervisord.conf /etc/supervisor/conf.d/supervisord.conf
# env
ENV         SSR_DIST=/dist/browser
ENV         SSR_PORT=8888
ENV         GIN_MODE=release
# setting
WORKDIR     /app
EXPOSE      7777
CMD         ["/usr/bin/supervisord"]