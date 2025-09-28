<template>
  <div class="comment-item px-8 py-6 hover:bg-gray-50/50 transition-colors duration-200">
    <div class="flex items-start justify-between">
      <div class="flex items-start space-x-4 flex-1">
        <!-- User Avatar -->
        <div class="w-12 h-12 bg-gradient-to-br from-orange-400 to-red-500 rounded-full flex items-center justify-center flex-shrink-0 shadow-sm">
          <span class="text-white font-semibold text-sm">
            {{ comment.user?.username?.charAt(0)?.toUpperCase() || 'U' }}
          </span>
        </div>
        
        <!-- Comment Content -->
        <div class="flex-1 min-w-0">
          <div class="flex items-center space-x-3 mb-2">
            <span class="font-semibold text-gray-900">{{ comment.user?.username || 'Anonymous' }}</span>
            <span class="text-gray-400 text-sm">•</span>
            <span class="text-gray-500 text-sm">{{ formatTime(comment.created_at) }}</span>
            
            
          </div>
          
          <p class="text-gray-700 whitespace-pre-wrap break-words leading-relaxed">{{ comment.content }}</p>
          
          <!-- Actions -->
          <div v-if="isCommentOwner" class="flex items-center space-x-4 mt-3">
            <button
              @click="startEditing"
              class="text-gray-500 hover:text-orange-500 text-sm font-medium transition-colors flex items-center space-x-1"
            >
              <Icon icon="mdi:pencil" class="h-4 w-4" />
              <span>Edit</span>
            </button>
            <button
              @click="handleDelete"
              class="text-gray-500 hover:text-red-500 text-sm font-medium transition-colors flex items-center space-x-1"
            >
              <Icon icon="mdi:delete" class="h-4 w-4" />
              <span>Delete</span>
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Edit Form -->
    <div v-if="isEditing" class="mt-4 pl-16">
      <div class="bg-gray-50 rounded-xl p-4">
        <textarea
          v-model="editContent"
          rows="3"
          class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent resize-none mb-3"
        ></textarea>
        
        <div class="flex justify-end space-x-3">
          <button
            @click="isEditing = false"
            class="px-4 py-2 text-gray-600 hover:text-gray-800 transition-colors"
          >
            Cancel
          </button>
          <button
            @click="handleUpdate"
            :disabled="!editContent.trim() || loading"
            class="px-6 py-2 bg-orange-500 text-white rounded-lg hover:bg-orange-600 disabled:opacity-50 transition-colors"
          >
            {{ loading ? 'Updating...' : 'Update' }}
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useAuth } from '~/composables/useAuth'

const props = defineProps({
  comment: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['update', 'delete'])

const { user } = useAuth()
const loading = ref(false)
const isEditing = ref(false)
const editContent = ref(props.comment.content)

// Check if current user is the comment owner
const isCommentOwner = computed(() => {
  return user.value && user.value.id === props.comment.user?.id
})

// Format relative time
const formatTime = (dateString) => {
  const date = new Date(dateString)
  const now = new Date()
  const diffInSeconds = Math.floor((now - date) / 1000)
  
  if (diffInSeconds < 60) return 'just now'
  if (diffInSeconds < 3600) return `${Math.floor(diffInSeconds / 60)}m ago`
  if (diffInSeconds < 86400) return `${Math.floor(diffInSeconds / 3600)}h ago`
  if (diffInSeconds < 604800) return `${Math.floor(diffInSeconds / 86400)}d ago`
  
  return date.toLocaleDateString('en-US', { 
    month: 'short', 
    day: 'numeric',
    year: date.getFullYear() !== now.getFullYear() ? 'numeric' : undefined
  })
}

const startEditing = () => {
  isEditing.value = true
  editContent.value = props.comment.content
}

const handleUpdate = async () => {
  if (!editContent.value.trim()) return
  
  loading.value = true
  await emit('update', props.comment.id, editContent.value.trim())
  loading.value = false
  isEditing.value = false
}

const handleDelete = async () => {
  if (confirm('Are you sure you want to delete this comment?')) {
    loading.value = true
    await emit('delete', props.comment.id)
    loading.value = false
  }
}
</script>