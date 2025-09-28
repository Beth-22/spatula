<template>
  <nav class="bg-[#F5F5F7] shadow-sm border-b border-gray-100 z-50 relative">
    <div class="max-w-7xl mx-auto px-4">
      <div class="flex justify-between items-center h-16">
        <!-- Logo -->
        <div class="flex items-center">
          <NuxtLink to="/" class="flex items-center space-x-3">
            <div class="w-8 h-8 bg-gradient-to-r from-orange-500 to-red-500 rounded-lg flex items-center justify-center">
              <Icon icon="mdi:chef-hat" class="h-5 w-5 text-white" />
            </div>
            <span class="text-2xl font-serif font-bold text-gray-900">SEASONED</span>
          </NuxtLink>
        </div>

        <!-- Right-side: My Recipes + Search + Bookmarks + Profile -->
        <div class="flex items-center space-x-4">
          <!-- My Recipes next to Search -->
          <NuxtLink 
            to="/my-recipes" 
            class="hidden md:inline-block text-black hover:text-orange-600 font-medium transition-colors"
          >
            My Recipes
          </NuxtLink>

          <button 
            @click="$router.push('/search')"
            class="p-2 text-gray-600 hover:text-orange-600 transition-colors"
          >
            <Icon icon="mdi:magnify" class="h-5 w-5" />
          </button>

          <button 
            @click="$router.push('/bookmarks')"
            class="p-2 text-gray-600 hover:text-orange-600 transition-colors"
          >
            <Icon icon="mdi:bookmark-outline" class="h-5 w-5" />
          </button>

          <!-- Profile Dropdown -->
          <div class="relative">
            <button
              @click="toggleProfileDropdown"
              class="flex items-center space-x-2 p-2 rounded-lg hover:bg-gray-50 transition-colors"
            >
              <div class="w-8 h-8 bg-gradient-to-r from-orange-500 to-red-500 rounded-full flex items-center justify-center">
                <span class="text-white font-semibold text-sm">
                  {{ user?.username?.charAt(0)?.toUpperCase() || '?' }}
                </span>
              </div>
            </button>

            <div v-if="showProfileDropdown" class="absolute right-0 mt-2 w-48 bg-white rounded-xl shadow-lg border border-gray-200 z-50">
              <!-- Dropdown content remains the same -->
            </div>
          </div>
        </div>
      </div>
    </div>

    <div 
      v-if="showProfileDropdown" 
      class="fixed inset-0 z-40" 
      @click="showProfileDropdown = false"
    ></div>
  </nav>
</template>

<script setup>
import { Icon } from '@iconify/vue'

const { user, logout } = useAuth()

const showProfileDropdown = ref(false)

const toggleProfileDropdown = () => {
  showProfileDropdown.value = !showProfileDropdown.value
}

const navigateToProfile = () => {
  showProfileDropdown.value = false
  navigateTo('/profile')
}

const navigateToLogin = () => {
  showProfileDropdown.value = false
  navigateTo('/login')
}

const handleLogout = async () => {
  showProfileDropdown.value = false
  await logout()
  await navigateTo('/')
}

const closeDropdown = (event) => {
  if (!event.target.closest('.relative')) {
    showProfileDropdown.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', closeDropdown)
})

onUnmounted(() => {
  document.removeEventListener('click', closeDropdown)
})
</script>