FROM scratch

COPY . .

EXPOSE 8081

CMD ["/cmd/app", "--host", "0.0.0.0", "--port", "8081", "--dbHost", "file:/catalog.db?_mutex=full&_cslike=false"]