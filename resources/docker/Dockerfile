FROM qnib/uplain-init

COPY qwatch-ng.yml /etc/
COPY lib/handlers /opt/qwatch-ng/lib/handlers
COPY lib/collectors /opt/qwatch-ng/lib/collectors
COPY lib/filters /opt/qwatch-ng/lib/filters
COPY opt/qnib/entry/* /opt/qnib/entry/
COPY test.log /resources/
ADD ./bin/* /usr/local/bin/
RUN echo "qwatch-ng --config=/etc/qwatch-ng.yml --ld-path=/opt/qwatch-ng/lib/" >> /root/.bash_history
CMD ["qwatch-ng", "--config=/etc/qwatch-ng.yml", "--ld-path=/opt/qwatch-ng/lib/"]
