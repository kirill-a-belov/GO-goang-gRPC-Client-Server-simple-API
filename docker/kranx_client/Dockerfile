FROM ubuntu:14.04
MAINTAINER Kirill Belov <kirill.a.belov@gmail.com>


ADD statistic.html /usr/bin/kranx/statistic.html 
ADD kranx_client /usr/bin/kranx/kranx_client
WORKDIR /usr/bin/kranx/


ENTRYPOINT ["kranx_client"]

