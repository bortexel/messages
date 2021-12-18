FROM golang:latest
WORKDIR /home/container
ADD messages .
ADD index.gohtml .
HEALTHCHECK --interval=1m --timeout=3s \
  CMD curl -f http://localhost:6868/ || exit 1
EXPOSE 6868
CMD [ "./messages" ]
