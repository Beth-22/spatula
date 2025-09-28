<template>
  <div class="min-h-screen flex items-center justify-center bg-white">
    <div class="w-full h-screen grid grid-cols-1 lg:grid-cols-2">
      <!-- Left Column - Steak Image Section -->
      <div class="relative h-full">
        <img 
          src="https://e1.pxfuel.com/desktop-wallpaper/766/853/desktop-wallpaper-steak-rib-eye-steak-animal-source-foods-cooking.jpg" 
          alt="Delicious steak"
          class="w-full h-full object-cover"
        >
        <div class="absolute inset-0 bg-gradient-to-r from-black/40 to-transparent"></div>
        <div class="absolute bottom-8 left-8 text-white">
          <h2 class="text-3xl font-bold mb-2">Welcome Back Chef!</h2>
          <p class="text-orange-200 text-lg">Continue your culinary journey with amazing recipes</p>
        </div>
      </div>

      <!-- Right Column - Login Form -->
      <div class="p-8 lg:p-12 flex flex-col justify-center bg-white">
        <div class="text-center mb-8">
          <div class="w-16 h-16 bg-gradient-to-r from-orange-500 to-amber-500 rounded-full flex items-center justify-center mx-auto mb-4">
            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
            </svg>
          </div>
          <h1 class="text-3xl font-bold text-gray-900 mb-2">Welcome Back</h1>
          <p class="text-gray-600">Sign in to your account</p>
        </div>

        <Form @submit="handleLogin" class="space-y-6 max-w-sm mx-auto w-full">
          <div class="space-y-4">
            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"></path>
                </svg>
              </div>
              <Field 
                v-model="form.email"
                name="email"
                type="email"
                placeholder="Email address"
                :rules="validateEmail"
                class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all duration-200"
              />
              <ErrorMessage name="email" class="text-red-500 text-sm mt-1" />
            </div>

            <div class="relative">
              <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
                </svg>
              </div>
              <Field 
                v-model="form.password"
                name="password"
                type="password"
                placeholder="Password"
                :rules="validateRequired"
                class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-all duration-200"
              />
              <ErrorMessage name="password" class="text-red-500 text-sm mt-1" />
            </div>
          </div>

          <button 
            type="submit"
            :disabled="loading"
            class="w-full bg-gradient-to-r from-orange-500 to-amber-500 text-white py-3 px-4 rounded-lg font-semibold hover:from-orange-600 hover:to-amber-600 transform hover:scale-105 transition-all duration-200 shadow-lg hover:shadow-xl disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <span v-if="loading" class="flex items-center justify-center">
              <svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              Signing In...
            </span>
            <span v-else>Sign In</span>
          </button>
        </Form>

        <div v-if="error" class="mt-4 p-4 bg-red-50 border border-red-200 rounded-lg max-w-sm mx-auto w-full">
          <div class="flex items-center">
            <svg class="w-5 h-5 text-red-500 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
            </svg>
            <span class="text-red-700 text-sm">{{ error }}</span>
          </div>
        </div>

        <div class="mt-8 text-center max-w-sm mx-auto w-full">
          <p class="text-gray-600">
            Don't have an account?
            <NuxtLink to="/signup" class="text-orange-500 hover:text-orange-600 font-semibold transition-colors">
              Sign up here
            </NuxtLink>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Form, Field, ErrorMessage } from 'vee-validate'

const form = ref({
  email: '',
  password: ''
})
const error = ref('')
const loading = ref(false)

// Custom validation functions
const validateRequired = (value) => {
  if (!value) {
    return 'This field is required'
  }
  return true
}

const validateEmail = (value) => {
  if (!value) {
    return 'Email is required'
  }
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  if (!emailRegex.test(value)) {
    return 'Please enter a valid email address'
  }
  return true
}

const handleLogin = async (values) => {
  try {
    loading.value = true
    error.value = ''
    
    const { login } = useAuth()
    await login(values.email, values.password)
    await navigateTo('/')
  } catch (err) {
    error.value = err.message
  } finally {
    loading.value = false
  }
}
</script>