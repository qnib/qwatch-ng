FROM qnib/uplain-init

COPY qwatch-ng.yml /etc/
COPY lib/handlers /opt/qwatch-ng/lib/handlers
COPY lib/collectors /opt/qwatch-ng/lib/collectors
COPY lib/filters /opt/qwatch-ng/lib/filters

ADD ./bin/* /usr/local/bin/

CMD ["qwatch-ng", "--config=/etc/qwatch-ng.yml", "--ld-path=/opt/qwatch-ng/lib/"]

