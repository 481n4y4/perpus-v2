<template>
  <div
    class="min-h-screen text-white bg-gray-900 flex flex-col justify-center py-12 sm:px-6 lg:px-8"
  >
    <div class="sm:mx-auto sm:w-full sm:max-w-md">
      <h2 class="text-center text-3xl font-extrabold">Buat Akun Baru</h2>
    </div>

    <div class="mt-8 sm:mx-auto sm:w-full sm:max-w-md bg-gray-700 rounded-lg">
      <div class="text-white py-8 px-4 shadow-md rounded-lg sm:px-10">
        <!-- Alert Success -->
        <div
          v-if="successMessage"
          class="mb-6 p-4 rounded-md bg-green-50 border border-green-200"
        >
          <div class="flex">
            <div class="ml-3">
              <h3 class="text-sm font-medium text-green-800">
                {{ successMessage }}
              </h3>
            </div>
          </div>
        </div>

        <!-- Alert Error -->
        <div
          v-if="errorMessage"
          class="mb-6 p-4 rounded-md bg-red-50 border border-red-200"
        >
          <div class="flex">
            <div class="ml-3">
              <h3 class="text-sm font-medium text-red-800">
                {{ errorMessage }}
              </h3>
              <div v-if="validationErrors" class="mt-2 text-sm text-red-700">
                <ul class="list-disc pl-5 space-y-1">
                  <li v-for="(error, field) in validationErrors" :key="field">
                    {{ Array.isArray(error) ? error[0] : error }}
                  </li>
                </ul>
              </div>
            </div>
          </div>
        </div>

        <form @submit.prevent="handleSignUp" class="space-y-6 bg-gray-700">
            <div>
                <label for="email" class="block text-sm font-medium">Alamat Email</label>
                <div class="mt-1">
                <input
                    id="email"
                    v-model="form.email"
                    name="email"
                    type="email"
                    autocomplete="email"
                    required
                    :disabled="isLoading"
                    class="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500 transition-all disabled:bg-gray-100 disabled:cursor-not-allowed text-gray-900" placeholder="nama@contoh.com"
                />
                </div>
            </div>

            <div>
                <label for="name" class="block text-sm font-medium">Username</label>
                <div class="mt-1">
                <input
                    id="name"
                    v-model="form.name"
                    name="name"
                    type="text"
                    autocomplete="username"
                    required
                    :disabled="isLoading"
                    class="appearance-none block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500 transition-all disabled:bg-gray-100 disabled:cursor-not-allowed text-gray-900" placeholder="nama_pengguna"
                />
                </div>
                <p class="mt-1 text-xs text-gray-300">
                Username hanya boleh berisi huruf, angka, dan garis bawah.
                </p>
            </div>

            <div>
                <label for="password" class="block text-sm font-medium">Kata Sandi</label>
                <div class="mt-1">
                <input
                    id="password"
                    v-model="form.password"
                    name="password"
                    type="password"
                    autocomplete="new-password"
                    required
                    :disabled="isLoading"
                    class="block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm placeholder-gray-400 focus:outline-none focus:ring-blue-500 focus:border-blue-500 transition-all disabled:bg-gray-100 disabled:cursor-not-allowed text-gray-900" placeholder="••••••••"
                />
                </div>
                <p class="mt-1 text-xs text-gray-300">Minimal 8 karakter.</p>
            </div>

            <div>
                <button
                type="submit"
                :disabled="isLoading"
                class="w-full flex justify-center items-center py-2.5 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer"
                >
                <span v-if="isLoading"> Memproses... </span>
                <span v-else> Daftar </span>
                </button>
                <p class="mt-4 text-center text-sm">
                Sudah punya akun?
                <NuxtLink
                    to="/auth/login"
                    class="font-medium text-blue-400 hover:text-blue-300 transition-colors ml-1"
                >
                    masuk di sini
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

const isLoading = ref(false);
const successMessage = ref("");
const errorMessage = ref("");
const validationErrors = ref({});

const form = ref({
  email: "",
  name: "",
  password: "",
});

const handleSignUp = async () => {
  successMessage.value = "";
  errorMessage.value = "";
  validationErrors.value = {};
  if (!form.value.email || !form.value.name || !form.value.password) {
    return (errorMessage.value = "Semua field harus diisi");
  }
  if (form.value.password.length < 8) {
    return (errorMessage.value = "password minimal 8 karakter");
  }

  isLoading.value = true;
  try {
    const response = await $fetch(`${apiBase}/auth/register`, {
      method: "POST",
      body: form.value,
      headers: { Accept: "application/json" },
    });
    successMessage.value = `${response.message || "Pendaftaran Berhasil"}`;

    form.value = { email: "", name: "", password: "" };
    setTimeout(() => navigateTo("/auth/login"), 2000);
  } catch (error) {
    const data = error.data;

    if (error.statusCode === 422) {
      validationErrors.value = data.errors || {};
      errorMessage.value = "terdapat kesalahan dalam pengisian form";
    } else {
      errorMessage.value = data?.message || "terdapat kesahalan saat mendaftar";
    }
  } finally {
    isLoading.value = false;
  }
};
</script>
