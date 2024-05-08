FROM debian:bookworm
EXPOSE 8080:8080

# Copy the backend binary
COPY ./card-jong-be .
ENTRYPOINT ["card-jong-be"]
