FROM alpine:3.19

COPY app.sh /app.sh

RUN chmod +x /app.sh

CMD ["/app.sh"]
