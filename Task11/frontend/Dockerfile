FROM node:21-alpine3.19

WORKDIR /react-example 

COPY package*.json ./

RUN npm install

COPY . . 


# Build the React app
# RUN npm run build

CMD [ "npm", "start" ]

# Stage 2: Production stage using Nginx base image
# FROM nginx:alpine

# # Copy the built app from the 'build' stage into the Nginx server directory
# COPY --from=build /app/build /usr/share/nginx/html

# # Expose port 80 to traffic
# EXPOSE 80

# # Containers run Nginx with global directives and daemon off
# CMD ["nginx", "-g", "daemon off;"]