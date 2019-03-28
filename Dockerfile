# Use aoepeople scratch with installed root certificates
FROM aoepeople/scratch-ca:latest

EXPOSE 8080
ADD mockingbirb /

CMD ["/mockingbirb"]
