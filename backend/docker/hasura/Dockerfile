FROM hasura/graphql-engine:v2.32.0

# Copy metadata if you have any custom configurations
COPY metadata /hasura-metadata

EXPOSE 8080

CMD ["graphql-engine", "serve"]