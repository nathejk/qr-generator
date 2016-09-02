FROM busybox

ADD bin/qr-generator /bin/
ADD zoneinfo.zip /usr/local/go/lib/time/

EXPOSE 80
CMD qr-generator
