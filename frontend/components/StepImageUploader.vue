<template>
  <div class="flex items-center gap-4">
    <!-- Image Preview -->
    <div v-if="image" class="relative">
      <img :src="image" class="w-16 h-16 object-cover rounded-lg" />
      <button
        @click="removeImage"
        class="absolute -top-2 -right-2 bg-red-500 text-white rounded-full p-1 hover:bg-red-600 transition-colors"
      >
        <Icon icon="mdi:close" class="h-3 w-3" />
      </button>
    </div>
    
    <!-- Upload Button -->
    <button
      @click="triggerFileInput"
      class="flex items-center gap-2 px-4 py-2 border border-gray-200 rounded-lg hover:border-orange-400 transition-colors"
    >
      <Icon icon="mdi:image-outline" class="h-4 w-4 text-gray-600" />
      <span class="text-sm text-gray-700">{{ image ? 'Change' : 'Add' }} Image</span>
    </button>
    
    <input
      ref="fileInput"
      type="file"
      accept="image/jpeg,image/png,image/webp,image/gif"
      @change="handleFileSelect"
      class="hidden"
    />
  </div>
</template>

<script setup>
const props = defineProps({
  stepIndex: Number,
  image: String
})

const emit = defineEmits(['image-uploaded'])

const fileInput = ref()

const triggerFileInput = () => {
  fileInput.value.click()
}

const handleFileSelect = async (event) => {
  const file = event.target.files[0]
  if (!file) return
  
  if (file.size > 10 * 1024 * 1024) {
    console.warn('File too large')
    return
  }
  
  try {
    const formData = new FormData()
    formData.append('file', file)
    
    const response = await $fetch('/actions/upload', {
      method: 'POST',
      body: formData
    })
    
    if (response.url) {
      emit('image-uploaded', props.stepIndex, response.url)
    }
  } catch (error) {
    console.error('Image upload failed:', error)
  }
  
  event.target.value = ''
}

const removeImage = () => {
  emit('image-uploaded', props.stepIndex, null)
}
</script>