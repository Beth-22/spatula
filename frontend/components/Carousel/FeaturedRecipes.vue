<template>
  <section class="w-full py-6 bg-#FFFFFFFF overflow-hidden">
    <div class="text-center mb-16 px-4">
      <h2 class="text-5xl font-serif font-bold text-gray-900 mb-0">TOP RATED RECIPES</h2>
      <p class="text-gray-600 text-xl max-w-2xl mx-auto">
        Discover our highest rated recipes loved by the community
      </p>
    </div>

    <div class="w-full">
      <Carousel :recipes="featuredRecipes" />
    </div>
    
    <div class="text-center mt-10 px-4">
      <NuxtLink 
        to="/recipes/"
        class="bg-primary-500 text-white px-8 py-3 rounded-lg hover:bg-primary-600 transition-colors font-medium text-md inline-block shadow-xl hover:shadow-2xl"
      >
        VIEW ALL TOP RATED RECIPES
      </NuxtLink>
    </div>
  </section>
</template>

<script setup>
// Featured recipes data
const featuredRecipes = [
  {
    id: 1,
    name: 'PENANG HOKKIEN PRAWN MEE',
    description: 'Authentic Malaysian street food with rich prawn broth, fresh noodles, bean sprouts, and hard-boiled eggs. A symphony of flavors that will transport you to the streets of Penang.',
    image: 'https://www.tasteatlas.com/Images/Dishes/1da821e8c98a42f38ede3fd8556652d2.jpg',
    time: '45 min',
    servings: '4 servings',
    rating: '4.9',
    premium: true
  },
  {
    id: 2,
    name: 'HEARTY TOMATO SOUP BOWL',
    description: 'Creamy roasted tomato soup with fresh basil, garlic croutons, and a drizzle of olive oil. Perfect for cozy evenings and family dinners.',
    image: 'https://loveonetoday.com/wp-content/uploads/avocado-bowl.png',
    time: '30 min',
    servings: '2 servings',
    rating: '4.7',
    premium: false
  },
  {
    id: 3,
    name: 'ETHIOPIAN INJERA PLATTER',
    description: 'Traditional fermented teff flatbread served with eight different stews including doro wat, misir wat, and shiro. A colorful and flavorful feast.',
    image: 'https://upload.wikimedia.org/wikipedia/commons/thumb/f/f6/Injera_with_eight_kinds_of_stew.jpg/500px-Injera_with_eight_kinds_of_stew.jpg',
    time: '2 hours',
    servings: '6 servings',
    rating: '4.8',
    premium: true
  },
  {
    id: 4,
    name: 'ITALIAN TRUFFLE PASTA',
    description: 'Fresh handmade pasta tossed in a creamy parmesan sauce with black truffle shavings and fresh herbs. An elegant dish for special occasions.',
    image: 'https://images.immediate.co.uk/production/volatile/sites/30/2022/04/5x-Reasons-to-Buy-Images-Web-1-1-9c00dac.png',
    time: '25 min',
    servings: '2 servings',
    rating: '4.9',
    premium: true
  },
  {
    id: 5,
    name: 'JAPANESE SUSHI PLATTER',
    description: 'Assorted fresh sushi including nigiri, maki rolls, and sashimi. Served with pickled ginger, wasabi, and soy sauce for an authentic experience.',
    image: 'https://thumbs.dreamstime.com/b/captivating-circular-food-photography-showcases-harmonious-blend-aesthetic-appeal-culinary-artistry-image-s-382088302.jpg',
    time: '1 hour',
    servings: '4 servings',
    rating: '4.6',
    premium: false
  }
]
</script>