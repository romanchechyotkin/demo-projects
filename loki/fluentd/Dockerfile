FROM fluent/fluentd:v1.17.1-debian-amd64-1.1

USER root

RUN apt-get update && \
    apt-get install -y build-essential libgeoip-dev libmaxminddb-dev

RUN gem install excon -v 0.109.0
RUN gem install faraday -v 2.8.1 --no-document
RUN gem install faraday-net_http -v 3.0.2 --no-document
RUN gem install fluent-plugin-beats --no-document
RUN gem install fluent-plugin-docker --no-document
RUN gem install fluent-plugin-grafana-loki --no-document

USER fluent
