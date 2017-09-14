FROM 172.16.13.110:5000/centos:7.2.1511

RUN mkdir /app

COPY oracle-instantclient11.2-devel-11.2.0.1.0-1.x86_64.rpm /tmp

COPY oracle-instantclient11.2-sqlplus-11.2.0.1.0-1.x86_64.rpm /tmp
 
COPY oracle-instantclient11.2-basic-11.2.0.1.0-1.x86_64.rpm /tmp

RUN rpm -ivh  /tmp/oracle-instantclient11.2-basic-11.2.0.1.0-1.x86_64.rpm

RUN rpm -ivh   /tmp/oracle-instantclient11.2-devel-11.2.0.1.0-1.x86_64.rpm

RUN rpm -ivh  /tmp/oracle-instantclient11.2-sqlplus-11.2.0.1.0-1.x86_64.rpm

RUN mkdir -p /usr/lib/oracle/11.2/client64/network/admin

ENV LD_LIBRARY_PATH $LD_LIBRARY_PATH:/usr/lib/oracle/11.2/client64/lib:/usr/local/lib

ENV ORACLE_HOME /usr/lib/oracle/11.2/client64

RUN /sbin/ldconfig

ADD libaio.so.1.0.1 /lib64/

RUN ln -s /lib64/libaio.so.1.0.1 /lib64/libaio.so.1

ADD sqlplus /usr/bin/

ADD main /app/main

ADD template /app/template

WORKDIR /app

EXPOSE 8888

ENTRYPOINT /app/main

