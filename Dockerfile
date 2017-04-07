FROM qnib/uplain-init

RUN apt-get update \
 && apt-get install -y wget \
 && wget -qO /usr/local/bin/go-github https://github.com/qnib/go-github/releases/download/0.2.2/go-github_0.2.2_Linux \
 && chmod +x /usr/local/bin/go-github \
 && echo "Download '$(/usr/local/bin/go-github rLatestUrl --ghorg qnib --ghrepo qwatch-ng --regex "qwatch-ng.*" --limit 1)'" \
 && wget -qO /usr/local/bin/qwatch-ng $(/usr/local/bin/go-github rLatestUrl --ghorg qnib --ghrepo qwatch-ng --regex "qwatch-ng.*" --limit 1) \
 && chmod +x /usr/local/bin/qwatch-ng \
 && mkdir -p /opt/qwatch-ng/ \
 && echo "Download '$(/usr/local/bin/go-github rLatestUrl --ghorg qnib --ghrepo qwatch-ng --regex "plugins.tar" --limit 1)'" \
 && wget -qO- $(/usr/local/bin/go-github rLatestUrl --ghorg qnib --ghrepo qwatch-ng --regex "plugins.tar" --limit 1) |tar xf - -C /opt/qwatch-ng/ \
 && apt-get purge -y wget \
 && apt-get autoclean \
 && rm -f /usr/local/bin/go-github

COPY resources/docker/opt/qnib/entry/* /opt/qnib/entry/
COPY resources/docker/qwatch-ng.yml /etc/
COPY resources/docker/test.log /resources/
RUN echo "qwatch-ng --config=/etc/qwatch-ng.yml --ld-path=/opt/qwatch-ng/lib/" >> /root/.bash_history
CMD ["qwatch-ng", "--config=/etc/qwatch-ng.yml", "--ld-path=/opt/qwatch-ng/lib/"]
