// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },
  devServer: {
    port: 5173,
  },

  runtimeConfig: {
    public: {
      apiBase: 'http://localhost:8080',
    },
  },
})
