// frontend/middleware/auth.global.js
export default defineNuxtRouteMiddleware((to) => {
  const { isAuthenticated } = useAuth();

  // Pages that require authentication
  const protectedRoutes = [
    "/profile",
    "/my-recipes",
    "/recipes/upload",
    "/shop",
  ];

  // Check if route is protected
  if (protectedRoutes.includes(to.path)) {
    console.log(
      "Protected route accessed:",
      to.path,
      "Authenticated:",
      isAuthenticated.value
    );

    if (!isAuthenticated.value) {
      console.log("Redirecting to login");
      return navigateTo("/login");
    }
  }

  // Fix: If user is authenticated but page is still loading, wait a bit
  if (protectedRoutes.includes(to.path) && isAuthenticated.value) {
    // This ensures the page loads properly when authenticated
    return;
  }
});
