# build stage
FROM golang:1.11-alpine AS build-env

ARG LD_FLAGS
ARG GOTAGS

WORKDIR "/go/src/awssm2env"

COPY . .

RUN \
  CGO_ENABLED="0" \
  go build -a -o "/awssm2env" 
#-ldflags "${LD_FLAGS}" -tags "${GOTAGS}"


# final stage
FROM alpine:3.9

RUN apk add --no-cache ca-certificates git

COPY --from=build-env /awssm2env /bin/awssm2env

ARG BUILD_DATE
ARG VCS_REF

ENTRYPOINT ["/bin/awssm2env"]
