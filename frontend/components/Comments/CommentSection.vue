<template>
  <div class="comments-section bg-white rounded-2xl border border-gray-100 shadow-sm mt-8">
    <!-- Header -->
    <div class="px-8 pt-8 pb-6 border-b border-gray-100">
      <h3 class="text-2xl font-serif font-bold text-gray-900 flex items-center">
        <Icon icon="mdi:comment-text-outline" class="h-6 w-6 text-orange-500 mr-3" />
        Comments
        <span v-if="totalComments > 0" class="text-gray-500 text-lg font-sans ml-2">
          {{ totalComments }}
        </span>
      </h3>
    </div>

    <!-- Comments List -->
    <div v-if="comments.length > 0" class="divide-y divide-gray-100">
      <div v-for="comment in comments" :key="comment.id" class="px-8 py-6">
        <div class="flex items-start space-x-4">
          <!-- User Avatar -->
          <div class="w-10 h-10 bg-gradient-to-br from-orange-400 to-red-500 rounded-full flex items-center justify-center flex-shrink-0">
            <span class="text-white font-semibold text-sm">
              {{ comment.user?.username?.charAt(0)?.toUpperCase() || 'U' }}
            </span>
          </div>
          
          <!-- Comment Content -->
          <div class="flex-1">
            <div class="flex items-center space-x-2 mb-1">
              <span class="font-semibold text-gray-900">{{ comment.user?.username || 'User' }}</span>
              <span class="text-gray-400 text-sm">•</span>
              <span class="text-gray-500 text-sm">{{ formatDate(comment.created_at) }}</span>
            </div>
            <p class="text-gray-700">{{ comment.content }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- Empty State -->
    <div v-else-if="!loading" class="px-8 py-12 text-center">
      <div class="w-20 h-20 bg-gradient-to-br from-orange-100 to-red-100 rounded-full flex items-center justify-center mx-auto mb-4">
        <Icon icon="mdi:comment-outline" class="h-8 w-8 text-orange-400" />
      </div>
      <h4 class="text-lg font-semibold text-gray-900 mb-2">No comments yet</h4>
      <p class="text-gray-600">Be the first to share your thoughts about this recipe!</p>
    </div>

    <!-- Comment Input -->
    <div class="px-8 py-6 bg-gray-50 border-t border-gray-100">
      <div v-if="isAuthenticated" class="flex space-x-4">
        <!-- User Avatar -->
        <div class="w-12 h-12 bg-gradient-to-br from-orange-400 to-red-500 rounded-full flex items-center justify-center flex-shrink-0">
          <span class="text-white font-semibold text-sm">
            {{ user?.username?.charAt(0)?.toUpperCase() || 'U' }}
          </span>
        </div>
        
        <!-- Comment Form -->
        <div class="flex-1">
          <textarea
            v-model="newComment"
            placeholder="Add a comment..."
            rows="3"
            class="w-full px-4 py-3 border border-gray-300 rounded-xl focus:outline-none focus:ring-2 focus:ring-orange-500 focus:border-transparent resize-none"
          ></textarea>
          
          <div class="flex justify-end mt-3">
            <button
              @click="submitComment"
              :disabled="!newComment.trim() || creatingComment"
              class="px-6 py-2 bg-orange-500 text-white rounded-lg hover:bg-orange-600 disabled:opacity-50 flex items-center space-x-2"
            >
              <span v-if="creatingComment" class="animate-spin">⟳</span>
              <span>Post Comment</span>
            </button>
          </div>
        </div>
      </div>

      <!-- Login Prompt -->
      <div v-else class="text-center">
        <p class="text-gray-600 mb-3">Join the conversation</p>
        <NuxtLink 
          to="/login" 
          class="inline-flex items-center space-x-2 text-orange-500 hover:text-orange-600 font-medium"
        >
          <span>Sign in to comment</span>
          <Icon icon="mdi:arrow-right" class="h-4 w-4" />
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useAuth } from '~/composables/useAuth'
import { useComments } from '~/composables/useComments'

const props = defineProps({
  recipeId: {
    type: String,
    required: true
  }
})

const { user, isAuthenticated } = useAuth()
const { getRecipeComments, createComment, loading: commentsLoading } = useComments()

const comments = ref([])
const totalComments = ref(0)
const newComment = ref('')
const creatingComment = ref(false)
const loading = ref(false)

const loadComments = async () => {
  loading.value = true
  const result = await getRecipeComments(props.recipeId)
  if (result) {
    comments.value = result.comments
    totalComments.value = result.total
  }
  loading.value = false
}

const submitComment = async () => {
  if (!newComment.value.trim()) return
  
  creatingComment.value = true
  const result = await createComment(props.recipeId, newComment.value.trim())
  
  if (result) {
    newComment.value = ''
    // Add the new comment to the top immediately
    comments.value = [result, ...comments.value]
    totalComments.value += 1
    console.log('✅ Comment added to UI immediately')
  }
  
  creatingComment.value = false
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  })
}

onMounted(() => {
  loadComments()
})
</script>