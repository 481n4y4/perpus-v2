<template>
  <div
    class="min-h-screen bg-gray-900 text-white flex flex-col justify-center items-center py-12 sm:px-6 lg:px-8"
  >
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
      <h2 class="text-center text-3xl font-extrabold">Masuk ke Akun Anda</h2>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md w-full px-4">
      <div class="bg-gray-800 py-8 px-4 shadow-md rounded-lg sm:px-10">
        <form class="space-y-6" @submit.prevent="handleLogin">
          <div>
            <label for="email" class="block text-sm font-medium">
              Alamat Email
            </label>
            <div class="mt-1">
              <input
                id="email"
                v-model="form.email"
                type="email"
                autocomplete="email"
                required
                class="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500 transition-all text-gray-900" 
                placeholder="nama@contoh.com"
                :disabled="pending"
              />
            </div>
          </div>

          <div>
            <label for="password" class="block text-sm font-medium">
              Kata Sandi
            </label>
            <div class="mt-1">
              <input
                id="password"
                v-model="form.password"
                type="password"
                autocomplete="current-password"
                required
                class="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500 transition-all text-gray-900" 
                placeholder="••••••••"
                :disabled="pending"
              />
            </div>
          </div>

          <div v-if="errorMessage" class="text-red-400 text-sm font-medium bg-red-900/30 border border-red-500/50 p-3 rounded-md">
            {{ errorMessage }}
          </div>

          <div>
            <button
              type="submit"
              :disabled="pending"
              class="w-full flex justify-center py-2.5 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors disabled:opacity-50 cursor-pointer"
            >
              {{ pending ? "Memproses..." : "Masuk" }}
            </button>

            <p class="mt-5 text-center text-sm text-gray-300">
              Atau
              <NuxtLink
                to="/auth/register"
                class="font-medium text-blue-400 hover:text-blue-300 transition-colors ml-1"
              >
                buat akun baru
              </NuxtLink>
            </p>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
  layout: "auth",
});

const {
  public: { apiBase },
} = useRuntimeConfig();

const form = ref({
  email: "",
  password: "",
});

const errorMessage = ref("");
const pending = ref(false);

const token = useCookie("auth_token", {
  maxAge: 60 * 60 * 24 * 7, 
  path: '/',                
  sameSite: 'lax',          
})

const handleLogin = async () => {
  pending.value = true;
  errorMessage.value = "";
  try {
    const res = await $fetch(`${apiBase}/auth/login`, {
      method: "POST",
      body: form.value,
      credentials: "include" 
    });
    
    // Explicitly set the cookie on the frontend so it's always accessible and avoids domain mismatch
    if (res.token) {
      token.value = res.token;
    }
    
    await navigateTo("/");
  } catch (err) {
    console.error(err);
    errorMessage.value = err.data?.message || "Terjadi kesalahan sistem.";
  } finally {
    pending.value = false;
  }
};
</script>