<template>
  <div
    class="relative transition-all duration-500 cursor-pointer group"
    :class="[size === 'large' ? 'scale-140 z-10' : 'scale-100 opacity-90']"
    @click="$emit('click')"
  >
    <!-- Main Container -->
    <div class="relative" :class="size === 'large' ? 'w-80 h-80' : 'w-64 h-64'">
      <!-- ⭐ ROTATING TEXT CIRCLE - ALWAYS VISIBLE ON ALL 3 IMAGES ⭐ -->
      <div
        class="absolute inset-0"
        :class="size === 'large' ? 'm-[-45px]' : 'm-[-40px]'"
      >
        <svg class="w-full h-full animate-spin-slow" viewBox="0 0 200 200">
          <defs>
            <path
              :id="'circlePath' + recipe.id"
              d="M 100, 100 m -85, 0 a 85,85 0 1,1 170,0 a 85,85 0 1,1 -170,0"
            />
          </defs>
          <text
            class="text-[9px] font-black fill-primary-500"
            text-anchor="middle"
            dominant-baseline="middle"
          >
            <textPath
              :href="'#circlePath' + recipe.id"
              startOffset="0%"
              method="align"
            >
              • {{ recipe.name.toUpperCase() }} •
              {{ recipe.name.toUpperCase() }} •
              {{ recipe.name.toUpperCase() }} •
            </textPath>
          </text>
        </svg>
      </div>

      <!-- Circular Image Container -->
      <div
        class="w-full h-full rounded-full overflow-hidden border-0 border-white shadow-2xl transition-all absolute inset-0 z-10"
        :class="size === 'large' ? 'm-0' : 'm-2'"
      >
        <img
          :src="recipe.image"
          :alt="recipe.name"
          class="w-full h-full object-cover"
        />
      </div>

      <!-- Premium Crown -->
      <div
        v-if="recipe.premium"
        class="absolute -top-2 -right-2 bg-yellow-400 rounded-full shadow-lg z-20 transition-all"
        :class="size === 'large' ? 'p-3' : 'p-2'"
      >
        <Icon
          name="lucide:crown"
          :class="size === 'large' ? 'h-6 w-6' : 'h-5 w-5'"
          class="text-white"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  recipe: {
    type: Object,
    required: true,
  },
  size: {
    type: String,
    default: "medium",
    validator: (value) => ["small", "medium", "large"].includes(value),
  },
  position: {
    type: String,
    default: "center",
    validator: (value) => ["left", "center", "right"].includes(value),
  },
  isActive: {
    type: Boolean,
    default: false,
  },
});

defineEmits(["click"]);
</script>

<style scoped>
.animate-spin-slow {
  animation: spin 25s linear infinite;
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

/* Ensure text is visible on all images */
text {
  font-size: 9px !important;
  font-weight: 900 !important;
  fill: #f97316 !important; /* primary-500 */
  letter-spacing: 1px;
}

/* Different rotation speeds for visual interest */
.group:nth-child(odd) .animate-spin-slow {
  animation-duration: 30s;
}

.group:nth-child(even) .animate-spin-slow {
  animation-duration: 20s;
}
</style>
