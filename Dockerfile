FROM golang AS builder
COPY game_of_life.go .
RUN go build game_of_life.go

FROM scratch
COPY --from=builder go/game_of_life .
CMD ["./game_of_life"]