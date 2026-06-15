<template>
  <div class="bg-gray-900 text-white">
    <div class="px-6 pt-14 lg:px-8">
      <h1 class="text-2xl font-bold mb-6 text-center">Edit Data Book</h1>

      <!-- FORM DUMMY -->
      <form 
          @submit.prevent="submitBook"
          class="bg-gray-800 p-6 rounded shadow space-y-4"
        >
        <!-- Name -->
        <div>
          <label class="block font-semibold mb-1">Name</label>
          <input
            v-model="form.name"
            type="text"
            class="block w-full rounded-md bg-gray-700 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6"
            placeholder="The Great Gatsby"
          />
        </div>

        <!-- Author -->
        <div>
          <label class="block font-semibold mb-1">Author</label>
          <input
          v-model="form.author"
            type="text"
            class="block w-full rounded-md bg-gray-700 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6"
            placeholder="F. Scott Fitzgerald"
          />
        </div>

        <!-- Publisher -->
        <div>
          <label class="block font-semibold mb-1">Publisher</label>
          <input
          v-model="form.publisher"
            type="text"
            class="block w-full rounded-md bg-gray-700 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6"
            placeholder="Scribner"
            />
        </div>

        <!-- Publish Date -->
        <div>
          <label class="block font-semibold mb-1">Publish Date</label>
          <input
          v-model="form.publish_date"
            type="number"
            class="block w-full rounded-md bg-gray-700 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6"
            placeholder="1925"
          />
        </div>

        <!-- Price -->
        <div>
          <label class="block font-semibold mb-1">Price</label>
          <input
          v-model="form.price"
            type="number"
            class="block w-full rounded-md bg-gray-700 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6"
            placeholder="12.99"
          />
        </div>

        <!-- Stock -->
        <div>
          <label class="block font-semibold mb-1">Stock</label>
          <input
          v-model="form.stock"
            type="number"
            class="block w-full rounded-md bg-gray-700 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-400 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6"
            placeholder="24"
            />
        </div>

        <!-- Book Cover -->
        <!-- <div>
          <label class="block font-semibold mb-1">Book Cover</label>
          <div
            class="mt-2 flex justify-center rounded-lg border border-dashed border-white/25 px-6 py-10"
          >
            <div class="text-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="mx-auto size-12 text-gray-600"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M2.25 15.75l5.159-5.159a2.25 2.25 0 013.182 0l5.159 5.159m-1.5-1.5l1.409-1.409a2.25 2.25 0 013.182 0l2.909 2.909m-18 3.75h16.5a1.5 1.5 0 001.5-1.5V6a1.5 1.5 0 00-1.5-1.5H3.75A1.5 1.5 0 002.25 6v12a1.5 1.5 0 001.5 1.5zm10.5-11.25h.008v.008h-.008V8.25zm.375 0a.375.375 0 11-.75 0 .375.375 0 01.75 0z"
                />
              </svg>
              <div class="mt-4 text-sm/6 text-gray-400">
                <label
                  for="file-upload"
                  class="relative cursor-pointer rounded-md bg-transparent font-semibold text-indigo-400 hover:text-indigo-300"
                >
                  <div>
                    <span>book-cover-gatsby.jpg</span>
                    <input
                      id="file-upload"
                      name="file-upload"
                      type="file"
                      class="sr-only"
                      disabled
                    />
                  </div>
                </label>
                <p class="pl-1">or drag and drop</p>
              </div>
              <p class="text-xs/5 text-gray-400">PNG, JPG, GIF up to 10MB</p>
            </div>
          </div>
        </div> -->

        <!-- Submit Button -->
        <div class="flex justify-end">
          <button
            type="submit"
            class="bg-blue-500 hover:bg-blue-600 text-white font-semibold px-6 py-2 rounded cursor-default"
    
          >
            Simpan
          </button>
        </div>
      </form>
    </div>
  </div>
</template>


<script setup>
definePageMeta({
    layout:"book"
})

const route = useRoute()
const id = route.params.id;
const apiBase = useRuntimeConfig().public.apiBase;

const book = await $fetch(`${apiBase}/books/${id}`)

const form = ref({
    name: book.data.name,
    author: book.data.author,
    publisher: book.data.publisher,
    publish_date: book.data.publish_date,
    price: book.data.price,
    stock: book.data.stock,
})

const submitBook = async () => {
    try {
        await $fetch(`${apiBase}/books/${id}`, {
            method: "PUT",
            body: form.value
        })
        alert("book successfully updated")
        navigateTo("/books")
    } catch (error) {
        console.error(error)
        const backendMessage = error.response?._data?.error || error.message
        alert("Failed to update book: " + backendMessage)
    }
}
</script>