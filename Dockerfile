FROM debian:stretch-slim
RUN mkdir /app
COPY bin/app /app/app
ENV DB_HOST="" DB_PORT=3306 DB_USER="" DB_PASS="" DB_NAME=""
EXPOSE 9090
WORKDIR /app
CMD ["/app/app"]