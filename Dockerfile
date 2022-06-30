FROM centos:centos7

RUN yum install -y\
    git\
    gcc\
    gcc-c++\
    make

RUN curl https://storage.googleapis.com/golang/go1.17.4.linux-amd64.tar.gz -o /var/tmp/go.tar.gz && \
    tar xzf /var/tmp/go.tar.gz -C /usr/local

ENV PATH=$PATH:/usr/local/bin:/usr/local/go/bin:/usr/local/go/path/bin 
