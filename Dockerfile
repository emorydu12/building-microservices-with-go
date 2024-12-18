FROM scratch

MAINTAINER orangeduxiaocheng@gmail.com

WORKDIR /usr/local/bin

EXPOSE 8080

COPY server .

ENTRYPOINT ["/usr/local/bin/server"]