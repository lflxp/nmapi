FROM centos:7

RUN yum install https://nmap.org/dist/nmap-7.70-1.x86_64.rpm -y
RUN yum install https://nmap.org/dist/ncat-7.70-1.x86_64.rpm -y
RUN yum install https://nmap.org/dist/nping-0.7.70-1.x86_64.rpm -y

ADD conf /opt/conf
ADD nmapi /opt/nmapi
ADD swagger /opt/swagger
ENV GOPATH /opt/
WORKDIR /opt/
EXPOSE 1987

CMD ["./nmapi"]
