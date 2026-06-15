<template>
  <div class="bg-gray-900">
    <div class="relative isolate px-6 pt-14 lg:px-8">
      <div class="mx-auto max-w-2xl py-32 sm:py-48 lg:py-56">
        <div class="hidden sm:mb-8 sm:flex sm:justify-center">
          <div
            class="relative rounded-full px-3 py-1 text-sm/6 text-gray-400 ring-1 ring-white/10 hover:ring-white/20"
          >
            Hello world!
          </div>
        </div>
        <div class="text-center">
          <h1
            class="text-5xl font-semibold tracking-tight text-balance text-white sm:text-7xl"
          >
            Perpustakaan Digital untuk Semua
          </h1>
          <p
            class="mt-8 text-lg font-medium text-pretty text-gray-400 sm:text-xl/8"
          >
            Akses ribuan buku kapan saja dan di mana saja. Tingkatkan literasi
            dan pengetahuan Anda dengan sistem perpustakaan yang modern dan
            mudah digunakan.
          </p>
          <!-- <div class="mt-10 flex items-center justify-center gap-x-6">
            <NuxtLink
              v-if="token"
              to="/books"
              class="rounded-md bg-indigo-500 px-3.5 py-2.5 text-sm font-semibold text-white shadow-xs hover:bg-indigo-400 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500"
              >Get started</NuxtLink
            >
            <NuxtLink
              v-else
              to="/auth/signup"
              class="rounded-md bg-indigo-500 px-3.5 py-2.5 text-sm font-semibold text-white shadow-xs hover:bg-indigo-400 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500"
              >Sign Up</NuxtLink
            >
          </div> -->
        </div>
      </div>
    </div>
  </div>
</template>


<!-- <script setup>
const token = useCookie("auth_token");
</script> -->