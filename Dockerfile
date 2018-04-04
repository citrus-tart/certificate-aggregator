FROM scratch
ADD main.exe /
ENV PORT=80
ENTRYPOINT ["/main.exe"]
EXPOSE 80