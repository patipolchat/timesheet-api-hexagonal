FROM rockylinux:9

RUN curl --proto '=https' --tlsv1.2 -sSf https://tiup-mirrors.pingcap.com/install.sh | sh

ENV PATH="${PATH}:/root/.tiup/bin/"

ENV VERSION=v7.5.0

RUN tiup install playground
RUN yum -y install mysql

ENTRYPOINT ["TIUP"]