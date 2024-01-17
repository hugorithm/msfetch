FROM golang:1.21 as build

WORKDIR /msfetch
COPY ./msfetch/* /msfetch/
RUN go build -o msfetch .

FROM alpine as runtime
COPY --from=build /msfetch/msfetch /usr/local/bin/msfetch
COPY run.sh /
RUN chmod +x /run.sh
ENTRYPOINT [ "./run.sh" ]
