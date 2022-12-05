FROM        ubuntu:latest
# install oracle
RUN         mkdir -p /opt/oracle
RUN         cd /opt/oracle
RUN         apt-get update
RUN         apt-get install -y libaio1 wget unzip
RUN         wget https://download.oracle.com/otn_software/linux/instantclient/218000/instantclient-basic-linux.x64-21.8.0.0.0dbru.zip
RUN         unzip instantclient-basic-linux.x64-21.8.0.0.0dbru.zip
RUN         rm -f instantclient-basic-linux.x64-21.8.0.0.0dbru.zip
RUN         echo /opt/oracle/instantclient_21_8 > /etc/ld.so.conf.d/oracle-instantclient.conf
RUN         ldconfig
RUN         export LD_LIBRARY_PATH=/opt/oracle/instantclient_21_8:$LD_LIBRARY_PATH
# install time zone
RUN         apt-get install tzdata
# copy file
COPY        jgo /app/
COPY        dist/ /app/dist/
COPY        dao/ /app/dao/
COPY        config/ /app/config/
COPY        sitemap.xml /app/sitemap.xml
# setting
WORKDIR     /app
EXPOSE      7777
ENTRYPOINT  ["/app/jgo"]