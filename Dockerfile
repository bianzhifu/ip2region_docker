FROM alpine:3.12
COPY ip2region.db /data/ip2region.db
COPY ip2region /ip2region
EXPOSE 8080
CMD /ip2region