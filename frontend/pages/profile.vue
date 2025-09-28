<template>
  <div class="min-h-screen bg-gradient-to-br from-orange-50 to-amber-50">
    <Navigation />

    <div class="max-w-2xl mx-auto px-4 py-12">
      <!-- Header -->
      <div class="text-center mb-12">
        <div class="w-24 h-24 bg-gradient-to-r from-orange-500 to-red-500 rounded-full flex items-center justify-center mx-auto mb-6 shadow-lg">
          <span class="text-white text-3xl font-bold">
            {{ user?.username?.charAt(0)?.toUpperCase() || 'U' }}
          </span>
        </div>
        <h1 class="text-4xl font-serif font-bold text-gray-900 mb-2">Profile Settings</h1>
        <p class="text-lg text-gray-600">Manage your account information</p>
      </div>

      <!-- Messages -->
      <div v-if="successMessage" class="mb-6 p-4 bg-green-100 border border-green-400 text-green-700 rounded-xl">
        {{ successMessage }}
      </div>
      <div v-if="errorMessage" class="mb-6 p-4 bg-red-100 border border-red-400 text-red-700 rounded-xl">
        {{ errorMessage }}
      </div>

      <!-- Profile Form -->
      <div class="bg-white rounded-2xl shadow-lg border border-orange-100 overflow-hidden">
        <div class="p-8">
          <form @submit.prevent="handleUpdateProfile" class="space-y-6">
            <!-- Username -->
            <div class="border-b border-gray-100 pb-6">
              <h3 class="text-lg font-semibold text-gray-900 mb-4">Basic Information</h3>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Username</label>
                <input
                  v-model="profileForm.username"
                  type="text"
                  class="w-full px-4 py-3 border border-gray-200 rounded-xl transition-all focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                  placeholder="Enter your username"
                />
                <p v-if="errors.username" class="text-red-500 text-sm mt-1">{{ errors.username }}</p>
              </div>

              <div class="flex justify-end mt-4">
                <button
                  type="submit"
                  :disabled="loading || !hasUsernameChanges"
                  class="px-6 py-2 bg-orange-500 text-white rounded-lg hover:bg-orange-600 disabled:opacity-50 transition-colors font-medium"
                >
                  <span v-if="loading">Updating...</span>
                  <span v-else>Update Username</span>
                </button>
              </div>
            </div>

            <!-- Password Change Section -->
            <div>
              <h3 class="text-lg font-semibold text-gray-900 mb-4">Change Password</h3>

              <!-- New Password -->
              <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-2">New Password</label>
                <input
                  v-model="profileForm.newPassword"
                  type="password"
                  class="w-full px-4 py-3 border border-gray-200 rounded-xl transition-all focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                  placeholder="Enter new password"
                />
                <p v-if="errors.newPassword" class="text-red-500 text-sm mt-1">{{ errors.newPassword }}</p>
              </div>

              <!-- Confirm Password -->
              <div v-if="profileForm.newPassword" class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-2">Confirm New Password</label>
                <input
                  v-model="profileForm.confirmPassword"
                  type="password"
                  class="w-full px-4 py-3 border border-gray-200 rounded-xl transition-all focus:ring-2 focus:ring-orange-500 focus:border-transparent"
                  placeholder="Confirm new password"
                />
                <p v-if="errors.confirmPassword" class="text-red-500 text-sm mt-1">{{ errors.confirmPassword }}</p>
              </div>

              <div class="flex justify-end">
                <button
                  type="button"
                  @click="updatePassword"
                  :disabled="!hasPasswordChanges || passwordLoading"
                  class="px-6 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 disabled:opacity-50 transition-colors font-medium"
                >
                  <span v-if="passwordLoading">Updating...</span>
                  <span v-else>Change Password</span>
                </button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>

    <Footer />
  </div>
</template>

<script setup>
const { user, isAuthenticated } = useAuth()

const loading = ref(false)
const passwordLoading = ref(false)
const successMessage = ref('')
const errorMessage = ref('')
const errors = ref({})

const profileForm = reactive({
  username: user.value?.username || '',
  newPassword: '',
  confirmPassword: ''
})

const originalValues = ref({ ...profileForm })

const hasUsernameChanges = computed(() => profileForm.username !== originalValues.value.username)
const hasPasswordChanges = computed(() => profileForm.newPassword.length > 0)

const validateForm = () => {
  errors.value = {}
  
  if (profileForm.newPassword && profileForm.newPassword.length < 6) {
    errors.value.newPassword = 'Password must be at least 6 characters'
  }
  
  if (profileForm.newPassword && profileForm.newPassword !== profileForm.confirmPassword) {
    errors.value.confirmPassword = 'Passwords must match'
  }
  
  return Object.keys(errors.value).length === 0
}

const handleUpdateProfile = async () => {
  if (!hasUsernameChanges.value) {
    successMessage.value = 'No changes to save'
    return
  }

  loading.value = true
  errorMessage.value = ''
  successMessage.value = ''

  try {
    // For now, just update locally since we don't have a backend endpoint
    // You'll need to add this to your useAuth composable
    successMessage.value = 'Profile update coming soon!'
    originalValues.value.username = profileForm.username
    setTimeout(() => successMessage.value = '', 3000)
  } catch (err) {
    errorMessage.value = err.message || 'Failed to update profile'
  } finally {
    loading.value = false
  }
}

const updatePassword = async () => {
  if (!hasPasswordChanges.value) return
  
  if (!validateForm()) {
    errorMessage.value = 'Please fix the validation errors'
    return
  }

  passwordLoading.value = true
  errorMessage.value = ''
  successMessage.value = ''

  try {
    successMessage.value = 'Password update coming soon!'
    profileForm.newPassword = ''
    profileForm.confirmPassword = ''
    setTimeout(() => successMessage.value = '', 5000)
  } catch (err) {
    errorMessage.value = err.message || 'Failed to update password'
  } finally {
    passwordLoading.value = false
  }
}

watch(user, (newUser) => {
  if (newUser) {
    profileForm.username = newUser.username
    originalValues.value = { ...profileForm }
  }
}, { immediate: true })

onMounted(() => {
  if (!isAuthenticated.value) {
    navigateTo('/login')
  }
})
</script>