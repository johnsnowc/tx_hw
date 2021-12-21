FROM centos:7
COPY server /root/server
COPY client-rpc /root/client-rpc
EXPOSE 8080
CMD /root/server