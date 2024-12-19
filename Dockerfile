FROM debian:stable-slim
COPY simple-go-server /bin/simple-go-server
ENV PORT=8080
CMD [ "/bin/simple-go-server" ]
