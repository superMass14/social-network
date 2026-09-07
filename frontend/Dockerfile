FROM node:alpine3.18

LABEL captain=" Ibrahima GUEYE <https://learn.zone01dakar.sn/git/igueye>"
LABEL chat-feature="Vincent Félix NDOUR <https://learn.zone01dakar.sn/git/vindour> && Madické YADE <https://learn.zone01dakar.sn/git/myade> "
LABEL post-feature="Masseck THIAW <https://learn.zone01dakar.sn/git/mthiaw>"
LABEL user-feature="Mamour Ousmane DRAME <https://learn.zone01dakar.sn/git/mamdrame>"
LABEL group-feature="Ibrahima GUEYE <https://learn.zone01dakar.sn/git/igueye> && Nicolas Cor FAYE <https://learn.zone01dakar.sn/git/nifaye>"
LABEL Name="BABBLE SOCIAL"
LABEL service="front-end"
LABEL version="1.0.0"
LABEL License="BABBLE team"

# working directory 
WORKDIR /app
# Copy oproject file into container's working dir
COPY . .
# Install dependencies
RUN npm install

EXPOSE 3000
# command to launch front-end service
CMD ["npm", "run", "dev"]