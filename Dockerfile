FROM centos:7
COPY server /root/server
EXPOSE 8080
CMD /root/server