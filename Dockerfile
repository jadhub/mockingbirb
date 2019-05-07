FROM aoepeople/scratch-go-env

ADD mock_config /mock_config
ADD bin/mockingbirb_unix bin/mockingbirb_unix

ENTRYPOINT ["bin/mockingbirb_unix"]

EXPOSE 3322

CMD ["serve"]
