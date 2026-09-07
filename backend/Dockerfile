FROM golang:alpine3.19

LABEL captain=" Ibrahima GUEYE <https://learn.zone01dakar.sn/git/igueye>"
LABEL chat-feature="Vincent Félix NDOUR <https://learn.zone01dakar.sn/git/vindour> & Madické YADE <https://learn.zone01dakar.sn/git/myade> "
LABEL post-feature="Masseck THIAW <https://learn.zone01dakar.sn/git/mthiaw>"
LABEL user-feature="Mamour Ousmane DRAME <https://learn.zone01dakar.sn/git/mamdrame>"
LABEL group-feature="Ibrahima GUEYE <https://learn.zone01dakar.sn/git/igueye> Nicolas Cor FAYE <https://learn.zone01dakar.sn/git/nifaye>"
LABEL Name="BABBLE SOCIAL"
LABEL service="back-end"
LABEL version="1.0.0"
LABEL License="BABBLE team"

#activate CGO
ENV CGO_ENABLED=1

# install dependencies
RUN apk update && apk add --no-cache git gcc musl-dev

# working directory 
WORKDIR /backend
# Copy oproject file into container's working dir
COPY . .

# build main.go file
RUN go build -o main .

EXPOSE 8080
# Command to launch backend service
CMD ["./main"]
