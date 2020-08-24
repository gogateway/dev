FROM scratch
COPY ./dev /usr/local/bin/dev
WORKDIR /usr/local/bin/
EXPOSE 80
CMD [ "dev" ]