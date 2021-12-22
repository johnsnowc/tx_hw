FROM centos:7
COPY server /root/server
COPY /client/client /root/client
EXPOSE 8080
CMD /root/server