<!-- pages/profile.vue -->
<template>
  <div
    class="min-h-screen bg-gray-900 text-white flex items-center justify-center p-4"
  >
    <div class="rounded-lg shadow-md p-6 w-full max-w-md bg-gray-700">
      <!-- Judul -->
      <h1 class="text-2xl font-bold text-center mb-6">Profil Saya</h1>

      <!-- Informasi -->
      <div class="space-y-4">
        <div>
          <label class="text-sm text-gray-200 font-medium">Nama</label>
          <p class="text-lg text-white mt-1">{{ user?.name }}</p>
        </div>

        <div>
          <label class="text-sm text-gray-200 font-medium">Email</label>
          <p class="text-lg text-white mt-1">{{ user?.email }}</p>
        </div>
      </div>

      <!-- Tombol Logout -->
      <button
        @click="handleLogout"
        class="mt-6 w-full bg-red-500 hover:bg-red-600 text-white py-2 rounded transition"
      >
        Logout
      </button>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: "auth",
});
const apiBase = useRuntimeConfig().public.apiBase;
const token = useCookie("auth_token");

const { data: user, error } = await useFetch("/auth/me", {
  baseURL: apiBase,
  method: "POST",
  headers: {
    Authorization: `Bearer ${token.value}`,
  },
});
console.log(token.value);
const handleLogout = async () => {
  try {
    await $fetch(`${apiBase}/auth/logout`, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token.value}`,
      },
    });
    token.value = null;
    navigateTo("/");
  } catch (error) {
    console.error("Logout error:", error);
  }
};
</script>
