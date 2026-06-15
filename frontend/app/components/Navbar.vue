<template>
  <nav
    class="fixed z-50 bg-gray-900 flex items-center w-screen justify-between p-6 lg:px-8"
  >
    <div class="flex lg:hidden">
      <button
        type="button"
        command="show-modal"
        commandfor="mobile-menu"
        class="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5"
      >
        <span class="sr-only">Open main menu</span>
        <svg
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="1.5"
          data-slot="icon"
          aria-hidden="true"
          class="size-6"
        >
          <path
            d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </button>
    </div>
    <div class="lg:flex lg:gap-x-12">
      <NuxtLink to="/" class="text-sm/6 font-semibold text-white"
        >{{ item1 }}</NuxtLink
      >
      <NuxtLink to="/about" class="text-sm/6 font-semibold text-white"
        >{{item2}}</NuxtLink
      >
      <NuxtLink to="/books" class="text-sm/6 font-semibold text-white"
        >{{ item3 }}</NuxtLink
      >
    </div>
    <div class="lg:flex lg:flex-1 lg:justify-end lg:gap-x-12">
      <NuxtLink
        v-if="token"
        to="/profile"
        class="text-sm/6 font-semibold text-white"
        >Profile</NuxtLink
      >
      <NuxtLink
        v-else
        to="/auth/signin"
        class="text-sm/6 font-semibold text-white"
        >Log in <span aria-hidden="true">&rarr;</span></NuxtLink
      >
    </div>
  </nav>
</template>

<script setup>
defineProps({
  item1: String,
  item2: String,
  item3: String
})
const token = useCookie("auth_token");
</script>
