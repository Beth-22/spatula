// composables/useAuth.js
import { gql } from "graphql-tag";

export const useAuth = () => {
  const user = useState("user", () => null);
  const authToken = useCookie("auth-token");
  const isAuthenticated = computed(() => !!authToken.value);

  // Signup function
  const signup = async (email, username, password) => {
    try {
      console.log("📝 Attempting signup for:", email, username);

      const response = await $fetch("http://localhost:8081/actions/signup", {
        method: "POST",
        body: JSON.stringify({
          email,
          username,
          password,
        }),
        headers: { "Content-Type": "application/json" },
      });

      console.log("✅ Signup response received:", response);

      if (!response?.token) {
        throw new Error(response?.error || "No token received");
      }

      // Set the auth token
      authToken.value = response.token;

      // Restore user session
      const restored = await restoreUserFromToken();

      if (!restored) {
        throw new Error("Failed to restore user session after signup");
      }

      console.log("🎉 Signup successful!");
      return response;
    } catch (error) {
      console.error("❌ Signup failed:", error);
      throw error;
    }
  };

  // Fetch user profile from GraphQL - REMOVE EMAIL FIELD
  const fetchUserProfile = async (userId) => {
    try {
      console.log("🔍 Fetching user profile for ID:", userId);

      const { $apollo } = useNuxtApp();

      const query = gql`
        query GetUserProfile($user_id: uuid!) {
          users_by_pk(id: $user_id) {
            id
            username
            created_at
          }
        }
      `;

      console.log("📡 Sending GraphQL query...");

      const result = await $apollo.query({
        query,
        variables: { user_id: userId },
        fetchPolicy: "network-only",
      });

      console.log("✅ GraphQL response received:", result);

      if (result.data?.users_by_pk) {
        console.log("✅ User profile fetched:", result.data.users_by_pk);
        return result.data.users_by_pk;
      } else {
        console.log("❌ No user data found in response");
        return null;
      }
    } catch (error) {
      console.error("❌ Error fetching user profile:", error);
      if (error.graphQLErrors) {
        error.graphQLErrors.forEach((err) => {
          console.error("❌ GraphQL Error:", err.message);
        });
      }
      return null;
    }
  };

  const restoreUserFromToken = async () => {
    if (!authToken.value) {
      console.log("❌ No auth token found");
      return false;
    }

    try {
      console.log("🔍 Restoring user from token...");
      const payload = authToken.value.split(".")[1];
      const decoded = JSON.parse(atob(payload));
      console.log("✅ Decoded JWT:", decoded);

      const hasuraClaims = decoded["https://hasura.io/jwt/claims"];
      const userId = hasuraClaims?.["x-hasura-user-id"] || decoded.sub;

      if (!userId) {
        console.error("❌ No user ID found in token");
        return false;
      }

      console.log("✅ User ID from token:", userId);

      // Fetch complete user profile from database
      const userProfile = await fetchUserProfile(userId);

      if (userProfile) {
        user.value = userProfile;
        console.log("🎉 User restored from database:", user.value);
        return true;
      } else {
        // Fallback: use basic info from token
        console.log("⚠️ Using fallback user data");
        user.value = {
          id: userId,
          username: "User",
        };
        console.log("✅ Basic user data set:", user.value);
        return true;
      }
    } catch (error) {
      console.error("❌ Failed to restore user from token:", error);
      return false;
    }
  };

  // Add refreshUser function
  const refreshUser = async () => {
    if (authToken.value) {
      return await restoreUserFromToken();
    }
    return false;
  };

  const login = async (email, password) => {
    try {
      console.log("🔐 Attempting login for:", email);

      const response = await $fetch("http://localhost:8081/actions/login", {
        method: "POST",
        body: JSON.stringify({ email, password }),
        headers: { "Content-Type": "application/json" },
      });

      console.log("✅ Login response received");

      if (!response?.token) {
        throw new Error("No token received");
      }

      authToken.value = response.token;
      const restored = await restoreUserFromToken();

      if (!restored) {
        throw new Error("Failed to restore user session");
      }

      console.log("🎉 Login successful!");
      return response;
    } catch (error) {
      console.error("❌ Login failed:", error);
      throw error;
    }
  };

  const logout = () => {
    console.log("🚪 Logging out");
    authToken.value = null;
    user.value = null;
    navigateTo("/login");
  };

  // Add updateProfile function for profile page
  const updateProfile = async (profileData) => {
    // For now, just update local state
    if (user.value) {
      user.value = { ...user.value, ...profileData };
    }
    return { success: true };
  };

  // Initialize on client side
  if (process.client) {
    console.log("🏁 Client-side auth initialization");
    if (authToken.value && !user.value) {
      console.log("🔄 Restoring user session...");
      restoreUserFromToken();
    }
  }

  // Debug function
  const debug = () => {
    console.log("=== AUTH DEBUG ===");
    console.log("Has token:", !!authToken.value);
    console.log("User state:", user.value);
    console.log("Is authenticated:", isAuthenticated.value);

    if (authToken.value) {
      try {
        const payload = authToken.value.split(".")[1];
        const decoded = JSON.parse(atob(payload));
        console.log("JWT user ID:", decoded.sub);
        console.log("Hasura claims:", decoded["https://hasura.io/jwt/claims"]);
      } catch (e) {
        console.error("Token decode error:", e);
      }
    }
  };

  return {
    user: readonly(user),
    isAuthenticated,
    signup, // Added signup function
    login,
    logout,
    debug,
    refreshUser,
    updateProfile,
  };
};
