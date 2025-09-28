<template>
  <div class="space-y-4">
    <!-- Image Grid -->
    <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-4">
      <!-- Add Image Card -->
      <div 
        class="border-2 border-dashed border-gray-300 rounded-xl h-32 flex items-center justify-center cursor-pointer hover:border-orange-400 transition-colors"
        @click="triggerFileInput"
      >
        <div class="text-center">
          <Icon icon="mdi:plus" class="h-8 w-8 text-gray-400 mx-auto mb-2" />
          <span class="text-sm text-gray-500">Add Image</span>
        </div>
      </div>
      
      <!-- Image Previews -->
      <div
        v-for="(image, index) in images"
        :key="index"
        class="relative group rounded-xl overflow-hidden h-32"
        :class="featuredIndex === index ? 'ring-2 ring-orange-500' : ''"
      >
        <img :src="image.preview || image.url" class="w-full h-full object-cover" />
        
        <!-- Overlay Actions -->
        <div class="absolute inset-0 bg-black/50 opacity-0 group-hover:opacity-100 transition-opacity flex items-center justify-center space-x-2">
          <button
            @click.stop="setFeatured(index)"
            class="p-2 bg-white/20 backdrop-blur-sm rounded-lg hover:bg-white/30 transition-colors"
            :class="featuredIndex === index ? 'text-yellow-400' : 'text-white'"
          >
            <Icon icon="mdi:star" class="h-4 w-4" />
          </button>
          <button
            @click.stop="removeImage(index)"
            class="p-2 bg-white/20 backdrop-blur-sm rounded-lg hover:bg-white/30 transition-colors text-white"
          >
            <Icon icon="mdi:trash-can-outline" class="h-4 w-4" />
          </button>
        </div>
        
        <!-- Featured Badge -->
        <div v-if="featuredIndex === index" class="absolute top-2 left-2">
          <span class="bg-orange-500 text-white text-xs px-2 py-1 rounded-full">Featured</span>
        </div>
      </div>
    </div>
    
    <!-- Hidden File Input -->
    <input
      ref="fileInput"
      type="file"
      multiple
      accept="image/jpeg,image/png,image/webp,image/gif"
      @change="handleFileSelect"
      class="hidden"
    />
  </div>
</template>

<script setup>
const props = defineProps({
  images: Array,
  maxImages: {
    type: Number,
    default: 10
  },
  featuredIndex: Number
})

const emit = defineEmits(['update:images', 'featured-changed', 'images-updated'])

const fileInput = ref()

const triggerFileInput = () => {
  if (props.images.length < props.maxImages) {
    fileInput.value.click()
  }
}

const handleFileSelect = (event) => {
  const files = Array.from(event.target.files)
  const remainingSlots = props.maxImages - props.images.length
  
  if (files.length > remainingSlots) {
    // Show error message
    console.warn(`Can only upload ${remainingSlots} more images`)
    files.splice(remainingSlots)
  }
  
  files.forEach(file => {
    if (file.size > 10 * 1024 * 1024) {
      console.warn('File too large:', file.name)
      return
    }
    
    const reader = new FileReader()
    reader.onload = (e) => {
      const newImages = [...props.images, {
        file: file,
        preview: e.target.result,
        url: null
      }]
      emit('update:images', newImages)
      emit('images-updated', newImages)
    }
    reader.readAsDataURL(file)
  })
  
  event.target.value = ''
}

const setFeatured = (index) => {
  emit('featured-changed', index)
}

const removeImage = (index) => {
  const newImages = props.images.filter((_, i) => i !== index)
  emit('update:images', newImages)
  emit('images-updated', newImages)
  
  if (props.featuredIndex === index) {
    emit('featured-changed', 0)
  }
}
</script>