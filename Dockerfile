# docker build -t track_order_detail_api:v1 .


# docker run -d -p 12345:12345 -v /data/ambuj_poc/track_order_detail_api/logs:/track_order_detail_api/logs track_order_detail_api:v1


FROM golang:1.18.4-bullseye

ENV LD_LIBRARY_PATH=/track_order_detail_api/instantclient_21_3

WORKDIR /track_order_detail_api/
EXPOSE 12345
ADD . .
RUN dpkg -i libaio1_0.3.112-9_amd64.deb
RUN dpkg -i libaio-dev_0.3.112-9_amd64.deb
RUN go build -o /track_order_detail_api/main_app ./main.go
CMD [ "/track_order_detail_api/main_app" ]