<template>
  <div class="bg-gray-900 text-white">
    <div class="px-6 pt-14 lg:px-8">
      <h1 class="text-2xl font-bold mb-6 text-center">Detail Data Book</h1>

      <!-- FORM DUMMY -->
      <div class="bg-gray-800 p-6 rounded shadow space-y-4">
        <!-- Name -->
        <div>
          <label class="block font-semibold mb-1">Name</label>
          <div
            class="block w-full rounded-md bg-gray-700 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 sm:text-sm/6"
          >
            {{ book.data.name }}
          </div>
        </div>

        <!-- Author -->
        <div>
          <label class="block font-semibold mb-1">Author</label>
          <div
            class="block w-full rounded-md bg-gray-700 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 sm:text-sm/6"
          >
            {{book.data.author}}
          </div>
        </div>

        <!-- Publisher -->
        <div>
          <label class="block font-semibold mb-1">Publisher</label>
          <div
            class="block w-full rounded-md bg-gray-700 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 sm:text-sm/6"
          >
            {{book.data.publisher}}
          </div>
        </div>

        <!-- Publish Date -->
        <div>
          <label class="block font-semibold mb-1">Publish Date</label>
          <div
            class="block w-full rounded-md bg-gray-700 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 sm:text-sm/6"
          >
            {{ book.data.publish_date }}
          </div>
        </div>

        <!-- Price -->
        <div>
          <label class="block font-semibold mb-1">Price</label>
          <div
            class="block w-full rounded-md bg-gray-700 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 sm:text-sm/6"
          >
            {{ book.data.price }}
          </div>
        </div>

        <!-- Stock -->
        <div>
          <label class="block font-semibold mb-1">Stock</label>
          <div
            class="block w-full rounded-md bg-gray-700 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 sm:text-sm/6"
          >
            {{ book.data.stock }}
          </div>
        </div>

        <!-- Book Cover (Image) -->
        <!-- <div>
          <label class="block font-semibold mb-1">Book Cover</label>
          <div
            class="mt-2 flex justify-center rounded-lg border border-dashed border-white/25 px-6 py-10"
          >
            <div class="flex flex-col gap-4 items-center">
              <img
                src="https://via.placeholder.com/300x400/1e293b/ffffff?text=The+Great+Gatsby+Cover"
                alt="Book cover"
                class="rounded-lg shadow-lg max-h-64 object-cover"
              />
              <div class="flex justify-center">
                <a
                  href="#"
                  class="bg-blue-500 hover:bg-blue-600 text-white font-semibold px-4 py-2 rounded inline-block cursor-default"
                >
                  Open File
                </a>
              </div>
              <p class="text-gray-400 text-sm">book-cover-gatsby.jpg</p>
            </div>
          </div>
        </div> -->

        <!-- Back Button -->
        <div class="flex justify-end mt-4">
          <NuxtLink
            :to="`/books`"
            class="bg-gray-600 hover:bg-gray-500 text-white font-semibold px-6 py-2 rounded cursor-default"
          >
            Kembali
          </NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
definePageMeta({
    layout:"book"
})

const apiBase = useRuntimeConfig().public.apiBase

const route = useRoute();
const id = route.params.id;
const book = await $fetch(`${apiBase}/books/${id}`)
</script>