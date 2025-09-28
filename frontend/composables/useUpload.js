// frontend/composables/useUpload.js
export const useUpload = () => {
  const uploadImage = async (file) => {
    const authToken = useCookie("auth-token");

    console.log("Uploading image:", file.name, "Size:", file.size);

    if (!authToken.value) {
      throw new Error("Authentication required for upload");
    }

    const formData = new FormData();
    formData.append("file", file);

    try {
      const response = await $fetch("http://localhost:8081/actions/upload", {
        method: "POST",
        body: formData,
        headers: {
          Authorization: `Bearer ${authToken.value}`,
        },
      });

      console.log("Upload response:", response);

      if (!response.url) {
        throw new Error("No URL returned from upload");
      }

      // Validate that it's a Supabase URL
      if (!response.url.includes("supabase.co")) {
        console.warn("Uploaded image is not from Supabase:", response.url);
      }

      return response;
    } catch (error) {
      console.error("Upload error:", error);
      throw new Error(error.data?.error || "Upload failed");
    }
  };

  return { uploadImage };
};
