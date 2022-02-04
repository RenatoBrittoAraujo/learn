FROM nginx:latest

COPY /index.html /usr/share/nginx/html/index.html

ENTRYPOINT ["/docker-entrypoint.sh"]

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]