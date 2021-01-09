FROM node:alpine as build

WORKDIR /app

COPY package.json .
RUN yarn install

COPY . .

RUN yarn build


FROM nginx:1.18.0-alpine as production

RUN mkdir /app

# Copy the conf and content to docker host
COPY --from=build /app/dist /app
COPY nginx/nginx.conf /etc/nginx/nginx.conf

EXPOSE 80
EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]
