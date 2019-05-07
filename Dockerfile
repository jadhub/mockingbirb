FROM aoepeople/scratch-go-env

ADD config /config
ADD bin/mockingbirb_unix /mockingbirb_unix

ENTRYPOINT ["/mockingbirb_unix"]

EXPOSE 3210

CMD ["serve"]
