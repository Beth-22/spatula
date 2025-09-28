<template>
  <div>
    <!-- Banner Section -->
    <div
      ref="banner"
      class="relative bg-cover bg-center h-[670px] overflow-visible"
      :style="{ backgroundImage: `url(${bgImage})` }"
    >
      <!-- Dropping Icons -->
      <img
        ref="pepper"
        src="@/assets/images/pepper.png"
        class="icon absolute w-[80px] h-auto"
      />
      <img
        ref="chips1"
        src="@/assets/images/chips.png"
        class="icon absolute w-[80px] h-auto"
      />
      <img
        ref="pepperExtra"
        src="@/assets/images/pepper2.png"
        class="icon absolute w-[80px] h-auto"
      />
      <!-- new 3rd icon -->
      <img
        ref="pepper2"
        src="@/assets/images/pepper2.png"
        class="icon absolute w-[80px] h-auto"
      />
      <img
        ref="tomato"
        src="@/assets/images/tomato.png"
        class="icon absolute w-[90px] h-auto"
      />
      <img
        ref="chips2"
        src="@/assets/images/chips.png"
        class="icon absolute w-[80px] h-auto"
      />
      <img
        ref="pepper3"
        src="@/assets/images/pepper2.png"
        class="icon absolute w-[80px] h-auto"
      />
      <img
        ref="chips3"
        src="@/assets/images/chips.png"
        class="icon absolute w-[80px] h-auto"
      />
      <img
        ref="chips4"
        src="@/assets/images/chips.png"
        class="icon absolute w-[80px] h-auto"
      />

      <!-- Hand + Noodle Icon -->
      <img
        ref="handNoodle"
        src="@/assets/images/icon2.png"
        alt="Hand holding noodle"
        class="absolute bottom-0 right-0 mr-0 transition-opacity transition-transform duration-1000 ease-out w-[340px] md:w-[580px] h-auto"
      />

      <!-- Title -->
      <div
        class="absolute inset-0 flex items-center justify-center -translate-y-24"
      >
        <h1
          class="font-poppins text-6xl md:text-9xl font-bold relative double-text"
        >
          <span class="layer1">SEASONED</span>
          <span class="layer2">SEASONED</span>
        </h1>
      </div>
    </div>

    <!-- FeaturedRecipes Section -->
    <FeaturedRecipes ref="featuredRef" />
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from "vue";
import FeaturedRecipes from "@/components/Carousel/FeaturedRecipes.vue";
import bgImage from "@/assets/images/orange2.png";

const banner = ref(null);
const handNoodle = ref(null);
const featuredRef = ref(null);

// Icon refs
const pepper = ref(null);
const chips1 = ref(null);
const pepperExtra = ref(null); // new icon
const pepper2 = ref(null);
const tomato = ref(null);
const chips2 = ref(null);
const pepper3 = ref(null);
const chips3 = ref(null);
const chips4 = ref(null);

// Clustered positions
const iconsOrder = [
  { ref: pepper, left: "8%", disappear: false },
  { ref: chips1, left: "16%", disappear: false },
  { ref: pepperExtra, left: "22%", disappear: false, higher: true }, // new icon, resting slightly higher
  { ref: pepper2, left: "28%", disappear: true }, // original 3rd icon falls & fades
  { ref: tomato, left: "32%", disappear: false },
  { ref: chips2, left: "40%", disappear: false },
  { ref: pepper3, left: "48%", disappear: false },
  { ref: chips3, left: "56%", disappear: false },
  { ref: chips4, left: "64%", disappear: true }, // last icon falls & fades
];

function triggerDrops() {
  if (!banner.value) return;

  iconsOrder.forEach((icon) => {
    if (!icon.ref.value) return;

    // Set horizontal position
    icon.ref.value.style.left = icon.left;

    // Reset animation
    icon.ref.value.classList.remove("drop");
    icon.ref.value.classList.add("reset");
    void icon.ref.value.offsetWidth;
    icon.ref.value.classList.remove("reset");
    icon.ref.value.classList.add("drop");

    // Rotation & fall distance
    let fallY;
    if (icon.disappear) {
      fallY = 1000 + Math.random() * 200; // fall past banner
    } else if (icon.higher) {
      fallY = 455 + Math.random() * 30; // slightly higher than other resting icons
    } else {
      fallY = 550 + Math.random() * 50; // resting just above bottom
    }
    const rotateDeg = Math.random() * 40 - 20;
    icon.ref.value.style.setProperty("--fallY", `${fallY}px`);
    icon.ref.value.style.setProperty("--rotate", `${rotateDeg}deg`);

    // Set disappear property
    icon.ref.value.dataset.disappear = icon.disappear ? "true" : "false";
  });
}

function onScroll() {
  if (!handNoodle.value || !featuredRef.value || !banner.value) return;

  const bannerRect = banner.value.getBoundingClientRect();

  if (bannerRect.top >= 0 && bannerRect.top < window.innerHeight * 0.7) {
    triggerDrops();
  }

  // Hand + noodle fade
  const featuredRect = featuredRef.value.$el.getBoundingClientRect();
  if (featuredRect.top < window.innerHeight * 0.8) {
    handNoodle.value.classList.add("opacity-0", "translate-y-8");
  } else {
    handNoodle.value.classList.remove("opacity-0", "translate-y-8");
  }
}

onMounted(() => {
  window.addEventListener("scroll", onScroll);
  triggerDrops();
});

onUnmounted(() => {
  window.removeEventListener("scroll", onScroll);
});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@400;600;700&display=swap");

.font-poppins {
  font-family: "Poppins", sans-serif;
}

/* Reset animation */
.reset {
  animation: slideUp 0.6s ease-out forwards;
}
@keyframes slideUp {
  0% {
    transform: translateY(600px);
    opacity: 1;
  }
  100% {
    transform: translateY(-200px);
    opacity: 0;
  }
}

/* Drop animation for normal icons */
.icon.drop {
  animation: dropAnim 3.5s cubic-bezier(0.3, 0.7, 0.4, 1) forwards;
}
@keyframes dropAnim {
  0% {
    transform: translateY(-200px) rotate(0deg);
    opacity: 1;
  }
  30% {
    opacity: 1;
  }
  100% {
    transform: translateY(var(--fallY, 550px)) rotate(var(--rotate, 0deg));
    opacity: 1;
  }
}

/* Special drop & fade for disappearing icons */
.icon[data-disappear="true"].drop {
  animation-name: dropAndFade;
}
@keyframes dropAndFade {
  0% {
    transform: translateY(-200px) rotate(0deg);
    opacity: 1;
  }
  80% {
    opacity: 1;
  }
  100% {
    transform: translateY(var(--fallY, 1200px)) rotate(var(--rotate, 0deg));
    opacity: 0;
  }
}
.double-text {
  position: relative;
  display: inline-block;
  text-align: center; /* ensures inline content is centered */
}

.double-text .layer1,
.double-text .layer2 {
  display: block;
  position: absolute;
  top: 50%;        /* center vertically */
  left: 45%;       /* center horizontally */
  transform: translate(-50%, -50%); /* perfectly center */
  font-family: "Poppins", sans-serif;
  font-weight: 700;
  font-size: 9rem;
  line-height: 1;
  letter-spacing: -0.05em;
}

.double-text .layer1 {
  color: #828c80; 
  z-index: 1;
  transform: translate(-52%, -60%) translate(4px, 4px); /* center + offset for shadow */
}

.double-text .layer2 {
  color: #ffd700; 
  z-index: 2;

}
</style>
