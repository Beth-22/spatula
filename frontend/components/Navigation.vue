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

        <!-- Right-side content -->
        <div class="flex items-center space-x-4">
          <!-- Search Button - Always Visible -->
          <button 
            @click="$router.push('/search')"
            class="p-2 text-gray-600 hover:text-orange-600 transition-colors"
            title="Search Recipes"
          >
            <Icon icon="mdi:magnify" class="h-5 w-5" />
          </button>

          <!-- Show these only when authenticated -->
          <template v-if="isAuthenticated">
            <!-- My Recipes -->
            <NuxtLink 
              to="/my-recipes" 
              class="hidden md:inline-block text-black hover:text-orange-600 font-medium transition-colors"
            >
              My Recipes
            </NuxtLink>

            <!-- Bookmarks -->
            <button 
              @click="$router.push('/bookmarks')"
              class="p-2 text-gray-600 hover:text-orange-600 transition-colors"
              title="My Bookmarks"
            >
              <Icon icon="mdi:bookmark-outline" class="h-5 w-5" />
            </button>

            <!-- Profile Dropdown -->
            <div class="relative">
              <button
                @click="toggleProfileDropdown"
                class="flex items-center space-x-2 p-2 rounded-lg hover:bg-gray-50 transition-colors"
                title="Profile Menu"
              >
                <div class="w-8 h-8 bg-gradient-to-r from-orange-500 to-red-500 rounded-full flex items-center justify-center">
                  <span class="text-white font-semibold text-sm">
                    {{ user?.username?.charAt(0)?.toUpperCase() || '?' }}
                  </span>
                </div>
              </button>

              <!-- Dropdown Content for Authenticated Users -->
              <div v-if="showProfileDropdown" class="absolute right-0 mt-2 w-48 bg-white rounded-xl shadow-lg border border-gray-200 z-50 py-2">
                <div class="px-4 py-2 border-b border-gray-100">
                  <p class="text-sm font-medium text-gray-900">{{ user?.username || 'User' }}</p>
                  <p class="text-xs text-gray-500">{{ user?.email || 'user@example.com' }}</p>
                </div>
                
                <button
                  @click="navigateToProfile"
                  class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 transition-colors flex items-center space-x-2"
                >
                  <Icon icon="mdi:account" class="h-4 w-4" />
                  <span>Profile</span>
                </button>
                
                <button
                  @click="handleLogout"
                  class="w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-red-50 transition-colors flex items-center space-x-2"
                >
                  <Icon icon="mdi:logout" class="h-4 w-4" />
                  <span>Logout</span>
                </button>
              </div>
            </div>
          </template>

          <!-- Show Login/Signup buttons when NOT authenticated -->
          <template v-else>
            <div class="flex items-center space-x-3">
              <NuxtLink 
                to="/login" 
                class="text-gray-700 hover:text-orange-600 font-medium transition-colors"
              >
                Login
              </NuxtLink>
              <NuxtLink 
                to="/signup" 
                class="bg-gradient-to-r from-orange-500 to-red-500 text-white px-4 py-2 rounded-lg font-medium hover:from-orange-600 hover:to-red-600 transition-all transform hover:scale-105"
              >
                Sign Up
              </NuxtLink>
            </div>
          </template>
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

const { user, logout, isAuthenticated } = useAuth()

const showProfileDropdown = ref(false)

const toggleProfileDropdown = () => {
  showProfileDropdown.value = !showProfileDropdown.value
}

const navigateToProfile = () => {
  showProfileDropdown.value = false
  navigateTo('/profile')
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