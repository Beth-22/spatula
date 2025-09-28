// frontend/nuxt.config.js
export default defineNuxtConfig({
  modules: [
    "@nuxtjs/tailwindcss",
    "@vee-validate/nuxt",
    "@vueuse/nuxt",
    "@nuxt/image",
  ],
  veeValidate: {
    autoImports: true,
  },
  runtimeConfig: {
    public: {
      chapaPublicKey: "",
      hasuraEndpoint: "http://localhost:8080/v1/graphql",
      authEndpoint: "http://localhost:8081/actions",
      goBackendUrl: "http://localhost:8081",
    },
  },

  compatibilityDate: "2025-09-23",
  devtools: { enabled: true },
});
