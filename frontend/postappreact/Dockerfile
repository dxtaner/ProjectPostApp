FROM node:14-alpine

# set working directory
WORKDIR /app

# ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package*.json ./
RUN npm install --silent
RUN npm install react-scripts@5.0.1 -g --silent

COPY . ./

# start app
CMD ["npm", "start"]
