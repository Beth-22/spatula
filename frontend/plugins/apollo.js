// frontend/plugins/apollo.js
import {
  ApolloClient,
  InMemoryCache,
  createHttpLink,
} from "@apollo/client/core";
import { setContext } from "@apollo/client/link/context";
import { provideApolloClient } from "@vue/apollo-composable";

export default defineNuxtPlugin((nuxtApp) => {
  const httpLink = createHttpLink({
    uri: "http://localhost:8080/v1/graphql",
  });

  // Auth link to add JWT token to requests
  const authLink = setContext((_, { headers }) => {
    // Get token from cookie
    const token = useCookie("auth-token").value;

    return {
      headers: {
        ...headers,
        ...(token ? { Authorization: `Bearer ${token}` } : {}),
        "x-hasura-role": token ? "user" : "anonymous",
      },
    };
  });

  const apolloClient = new ApolloClient({
    link: authLink.concat(httpLink),
    cache: new InMemoryCache(),
  });

  provideApolloClient(apolloClient);

  // Make Apollo client available via $apollo
  nuxtApp.provide("apollo", apolloClient);
});
